package saga

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"acc.accommodation.com/config"
	"acc.accommodation.com/src/model"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var USER_RESTORE = "user_restore"
var USER_DELETED = "user_deleted"
var ACCOMMODATION_DELETED = "accommodation_deleted"
var ACCOMMODATION_RESTORE = "accommodation_restore"

type SAGA struct {
	// NATS connection
	nc *nats.Conn

	accommodation_collection *mongo.Collection
}

func CreateSAGA(cfg *config.Config, accommodation_collection *mongo.Collection) (*SAGA, error) {
	nc, _ := nats.Connect(cfg.NatsAddress)
	saga := &SAGA{nc: nc, accommodation_collection: accommodation_collection}
	log.Println("SAGA AccommodationService: Connected to NATS")
	saga.RegisterConsumers()
	log.Println("SAGA AccommodationService: Registered consumers")

	return saga, nil
}

func (s *SAGA) Close() {
	s.nc.Close()
}

func (s *SAGA) RegisterConsumers() {
	s.nc.Subscribe(USER_DELETED, s.DeleteAccommodationConsumer)
	s.nc.Subscribe(ACCOMMODATION_RESTORE, s.RestoreAccommodationConsumer)
}

// Action functions
func (s *SAGA) DeleteAccommodationConsumer(m *nats.Msg) {
	// Get userId and role from message that is json
	var data map[string]string
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	json.Unmarshal(m.Data, &data)
	role := data["role"]
	userId, err := primitive.ObjectIDFromHex(data["userId"])
	if err != nil {
		log.Println("SAGA AccommodationService: User id is not valid")
		m.Term()
		return
	}
	if role == "H" {
		log.Println("SAGA AccommodationService: Deleting accommodation for host:: " + userId.Hex() + "...")
		// find all accommodations for host with userId
		var accommodationIds []string
		cursor, err := s.accommodation_collection.Find(ctx, primitive.M{"user_id": userId})
		if err != nil {
			log.Println("SAGA AccommodationService: Error while finding accommodations for host:: " + userId.Hex())
			m.Term()
			return
		}
		for cursor.Next(ctx) {
			var accommodation model.Accommodation
			cursor.Decode(&accommodation)
			accommodationIds = append(accommodationIds, accommodation.Id.Hex())
		}
		// Delete all accommodations for host with userId update deletedAt field
		_, err = s.accommodation_collection.UpdateMany(ctx, primitive.M{"user_id": userId}, primitive.M{"$set": primitive.M{"deleted_at": time.Now()}})
		if err != nil {
			log.Println("SAGA AccommodationService: Error while deleting accommodations for host:: " + userId.Hex())
			m.Term()
			// Public restore user
			s.PublishRestoreUser(userId.Hex())
			return
		}

		// Public message to NATS what accommodations are deleted
		s.PublishAccommodationsDeleted(accommodationIds)
	}

	m.Ack()

}

func (s *SAGA) RestoreAccommodationConsumer(m *nats.Msg) {
	// Get accommodationIds from message that is json
	var accommodationIds []string
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	json.Unmarshal(m.Data, &accommodationIds)
	log.Println("SAGA AccommodationService: Restoring accommodations...")
	// convert accommodationIds to primitive.ObjectID
	var accommodationObjectIds []primitive.ObjectID
	for _, id := range accommodationIds {
		accommodationObjectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Println("SAGA AccommodationService: Accommodation id is not valid")
			m.Term()
			return
		}
		accommodationObjectIds = append(accommodationObjectIds, accommodationObjectId)
	}
	// Restore all accommodations with accommodationIds update deletedAt field
	_, err := s.accommodation_collection.UpdateMany(ctx, primitive.M{"_id": primitive.M{"$in": accommodationObjectIds}}, primitive.M{"$unset": primitive.M{"deleted_at": ""}})
	if err != nil {
		log.Println("SAGA AccommodationService: Error while restoring accommodations")
		m.Term()
		return
	}

	// find hostId for accommodationIds
	var accommodation model.Accommodation

	err = s.accommodation_collection.FindOne(ctx, primitive.M{"_id": primitive.M{"$in": accommodationObjectIds}}).Decode(&accommodation)
	if err != nil {
		log.Println("SAGA AccommodationService: Error while finding hostId for accommodationIds")
		m.Term()
		return
	}
	s.PublishRestoreUser(accommodation.UserId.Hex())
	m.Ack()

}

// Publish functions
func (s *SAGA) PublishAccommodationsDeleted(accommodations []string) {
	// Publish message to NATS what accommodations are deleted as json
	json, _ := json.Marshal(accommodations)
	s.nc.Publish(ACCOMMODATION_DELETED, json)
}

func (s *SAGA) PublishRestoreUser(userID string) {
	s.nc.Publish(USER_RESTORE, []byte(userID))
}
