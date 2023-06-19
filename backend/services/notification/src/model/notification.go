package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"notification.accommodation.com/pb"
)

// mongodb model
type Notification struct {
	Id         primitive.ObjectID `bson:"_id,omitempty"`
	Type       string             `bson:"type,omitempty"`
	ResourceId string             `bson:"resource_id,omitempty"`
	Body       string             `bson:"body,omitempty"`
	UserId     primitive.ObjectID `bson:"user_id,omitempty"`
	IsRead     bool               `bson:"is_read,omitempty"`
	CreatedAt  time.Time          `bson:"created_at,omitempty"`
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
