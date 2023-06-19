package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"notification.accommodation.com/pb"
)

// mongodb model
type Notification struct {
	Id         primitive.ObjectID `bson:"_id"`
	Type       string             `bson:"type"`
	ResourceId string             `bson:"resource_id"`
	Body       string             `bson:"body"`
	UserId     primitive.ObjectID `bson:"user_id"`
	IsRead     bool               `bson:"is_read"`
	CreatedAt  time.Time          `bson:"created_at"`
}

// model -> proto
func (n *Notification) ToProto() *pb.Notification {
	return &pb.Notification{
		Id:         n.Id.Hex(),
		Type:       n.Type,
		ResourceId: n.ResourceId,
		Body:       n.Body,
		UserId:     n.UserId.Hex(),
		IsRead:     n.IsRead,
		CreatedAt:  n.CreatedAt.String(),
	}
}
