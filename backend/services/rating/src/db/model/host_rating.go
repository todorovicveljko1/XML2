package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rating.accommodation.com/pb"
)

type HostRating struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	HostId primitive.ObjectID `bson:"host_id"`
	UserId primitive.ObjectID `bson:"user_id"`
	Rating int32              `bson:"rating"`
}

// convertToPbHostRating converts HostRating to pb.HostRating
func (h *HostRating) ConvertToPbHostRating() *pb.HostRating {

	return &pb.HostRating{
		Id:     h.ID.Hex(),
		HostId: h.HostId.Hex(),
		UserId: h.UserId.Hex(),
		Rating: h.Rating,
	}
}
