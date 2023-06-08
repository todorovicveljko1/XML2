package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"rating.accommodation.com/pb"
)

type AccommodationRating struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	AccommodationId primitive.ObjectID `bson:"accommodation_id"`
	UserId          primitive.ObjectID `bson:"user_id"`
	Rating          int32              `bson:"rating"`
}

// convertToPbAccommodationRating converts AccommodationRating to pb.AccommodationRating
func (a *AccommodationRating) ConvertToPbAccommodationRating() *pb.AccommodationRating {

	return &pb.AccommodationRating{
		Id:              a.ID.Hex(),
		AccommodationId: a.AccommodationId.Hex(),
		UserId:          a.UserId.Hex(),
		Rating:          a.Rating,
	}
}
