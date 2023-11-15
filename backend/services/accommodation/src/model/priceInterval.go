package model

import (
	"time"

	"acc.accommodation.com/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PriceInterval struct {
	Id              primitive.ObjectID `bson:"_id" json:"id"`
	AccommodationId primitive.ObjectID `bson:"accommodation_id" json:"accommodation_id"`
	StartDate       time.Time          `bson:"start_date" json:"start_date"`
	EndDate         time.Time          `bson:"end_date" json:"end_date"`
	Price           float64            `bson:"price" json:"price"`
}

// Convert to protož

func (p *PriceInterval) ToProto() *pb.PriceInterval {
	return &pb.PriceInterval{
		// Id:              p.Id.Hex(),
		// AccommodationId: p.AccommodationId.Hex(),
		StartDate: p.StartDate.String(),
		EndDate:   p.EndDate.String(),
		Price:     p.Price,
	}
}
