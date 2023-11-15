package saga

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"reservation.accommodation.com/config"
)

var ACCOMMODATION_DELETED = "accommodation_deleted"
var ACCOMMODATION_RESTORE = "accommodation_restore"

type SAGA struct {
	// NATS connection
	nc *nats.Conn

	reservation_collection *mongo.Collection
}

func CreateSAGA(cfg *config.Config, reservation_collection *mongo.Collection) (*SAGA, error) {
	nc, _ := nats.Connect(cfg.NatsAddress)
	saga := &SAGA{nc: nc, reservation_collection: reservation_collection}
	log.Println("SAGA ReservationService: Connected to NATS")
	saga.RegisterConsumers()
	log.Println("SAGA ReservationService: Registered consumers")

	return saga, nil
}

func (s *SAGA) Close() {
	s.nc.Close()
}

func (s *SAGA) RegisterConsumers() {
	s.nc.Subscribe(ACCOMMODATION_DELETED, s.DeleteReservationConsumer)
}

// Action functions
func (s *SAGA) DeleteReservationConsumer(m *nats.Msg) {
	// Get accommodationIds from message that is json

	var accommodationIds []string
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	json.Unmarshal(m.Data, &accommodationIds)
	// TEST if some error occurs
	// s.RestoreAccommodationProducer(accommodationIds)
	// m.Ack()
	// return
	// Convert accommodationIds to primitive.ObjectID
	var accommodationObjectIds []primitive.ObjectID
	for _, id := range accommodationIds {
		accommodationObjectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Println("SAGA ReservationService: Accommodation id is not valid")
			m.Term()
			return
		}
		accommodationObjectIds = append(accommodationObjectIds, accommodationObjectId)
	}
	// Delete reservations with accommodationIds by updating deleted_at field
	filter := bson.M{"accommodation_id": bson.M{"$in": accommodationObjectIds}}
	update := bson.M{"$set": bson.M{"deleted_at": time.Now()}}
	_, err := s.reservation_collection.UpdateMany(ctx, filter, update)
	if err != nil {
		log.Println("SAGA ReservationService: Failed to delete reservations")
		m.Term()
		s.RestoreAccommodationProducer(accommodationIds)
		return
	}

	log.Println("SAGA ReservationService: Deleted reservations")
	m.Ack()
}

func (s *SAGA) RestoreAccommodationProducer(accommodationIds []string) {
	// Send message to NATS
	data, _ := json.Marshal(accommodationIds)
	s.nc.Publish(ACCOMMODATION_RESTORE, data)
}
