package model

import (
	"time"

	"acc.accommodation.com/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AvailableInterval struct {
	Id              primitive.ObjectID `bson:"_id" json:"id"`
	AccommodationId primitive.ObjectID `bson:"accommodation_id" json:"accommodation_id"`
	StartDate       time.Time          `bson:"start_date" json:"start_date"`
	EndDate         time.Time          `bson:"end_date" json:"end_date"`
	IsAvailable     bool               `bson:"is_available" json:"is_available"`
}

// Convert to proto

func (a *AvailableInterval) ToProto() *pb.AvailableInterval {
	return &pb.AvailableInterval{
		// Id:              a.Id.Hex(),
		// AccommodationId: a.AccommodationId.Hex(),
		StartDate:   a.StartDate.String(),
		EndDate:     a.EndDate.String(),
		IsAvailable: a.IsAvailable,
	}
}
