package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rating.accommodation.com/pb"
)

type Rating struct {
	ID                  primitive.ObjectID `bson:"_id,omitempty"`
	ReservationId       primitive.ObjectID `bson:"reservation_id"`
	HostId              primitive.ObjectID `bson:"host_id"`
	AccommodationId     primitive.ObjectID `bson:"accommodation_id"`
	UserId              primitive.ObjectID `bson:"user_id"`
	HostRating          int32              `bson:"host_rating"`
	AccommodationRating int32              `bson:"accommodation_rating"`
}

func (a *Rating) ConvertToPbAccommodationRating() *pb.Rating {

	return &pb.Rating{
		Id:                  a.ID.Hex(),
		ReservationId:       a.ReservationId.Hex(),
		HostId:              a.HostId.Hex(),
		AccommodationId:     a.AccommodationId.Hex(),
		UserId:              a.UserId.Hex(),
		HostRating:          a.HostRating,
		AccommodationRating: a.AccommodationRating,
	}
}
