package saga

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"auth.accommodation.com/config"
	"auth.accommodation.com/src/model"
	"github.com/nats-io/nats.go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var USER_DELETED = "user_deleted"
var USER_DELETE = "user_delete"
var USER_RESTORE = "user_restore"

type SAGA struct {
	// NATS connection
	nc *nats.Conn

	user_collection *mongo.Collection
}

func CreateSAGA(cfg config.Config, user_collection *mongo.Collection) (*SAGA, error) {
	nc, _ := nats.Connect(cfg.NatsAddress)

	saga := &SAGA{nc: nc, user_collection: user_collection}
	log.Println("SAGA AuthService: Connected to NATS")
	saga.RegisterConsumers()
	log.Println("SAGA AuthService: Registered consumers")

	return saga, nil
}

func (s *SAGA) Close() {
	s.nc.Close()
}

func (s *SAGA) RegisterConsumers() {
	s.nc.Subscribe(USER_DELETE, s.DeleteUserConsumer)
	s.nc.Subscribe(USER_RESTORE, s.RestoreUserConsumer)
}

// Action functions
func (s *SAGA) DeleteUserConsumer(m *nats.Msg) {
	var user model.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Get user id from message
	userID, err := primitive.ObjectIDFromHex(string(m.Data))
	if err != nil {
		log.Println("SAGA AuthService: User id is not valid")
		m.Term()
		return
	}
	log.Println("SAGA AuthService: Deleting user:: " + userID.Hex() + "...")
	// Get user from database
	err = s.user_collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		log.Println("SAGA AuthService: User not found")
		m.Term()
		return
	}
	// update user in database
	_, err = s.user_collection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{
		"$set": bson.M{
			"deleted_at": time.Now(),
		},
	})
	if err != nil {
		log.Println("SAGA AuthService: Failed to delete user")
		m.Term()
		return
	}
	// publish user deleted event
	s.DeletedUserProducer(user.Id.Hex(), string(user.Role))
	log.Println("SAGA AuthService: User deleted")
	m.Ack()
}
func (s *SAGA) RestoreUserConsumer(m *nats.Msg) {
	var user model.User
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Get user id from message
	userID, err := primitive.ObjectIDFromHex(string(m.Data))
	if err != nil {
		log.Println("SAGA AuthService: User id is not valid")
		m.Term()
		return
	}
	log.Println("SAGA AuthService: Restoring user:: " + userID.Hex() + "...")
	// Get user from database
	err = s.user_collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err != nil {
		log.Println("SAGA AuthService: User not found")
		m.Term()
		return
	}
	// set user deleted_at to nil
	user.DeletedAt = nil
	// update user in database
	s.user_collection.UpdateOne(ctx, bson.M{"_id": userID}, bson.M{
		"$unset": bson.M{
			"deleted_at": "",
		},
	})
	m.Ack()

}

// Next function
func (s *SAGA) DeletedUserProducer(userId string, role string) {
	// publish user deleted event that has json as payload with user id and role
	json, _ := json.Marshal(map[string]string{"userId": userId, "role": role})
	s.nc.Publish(USER_DELETED, json)

}
