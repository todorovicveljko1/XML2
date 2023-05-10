package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"reservation.accommodation.com/pb"
)

// STATUS:: PENNDING, APPROVED, REJECTED, CANCELED, AUTO_REJECTED
type Reservation struct {
	Id              primitive.ObjectID `bson:"_id" json:"id"`
	UserId          primitive.ObjectID `bson:"user_id" json:"user_id"`
	AccommodationId primitive.ObjectID `bson:"accommodation_id" json:"accommodation_id"`
	StartDate       primitive.DateTime `bson:"start_date" json:"start_date"`
	EndDate         primitive.DateTime `bson:"end_date" json:"end_date"`
	Status          string             `bson:"status" json:"status"`
	Price           float64            `bson:"price" json:"price"`
}

// ConvertToPbReservation converts Reservation to pb.Reservation
func (r *Reservation) ConvertToPbReservation() *pb.Reservation {
	return &pb.Reservation{
		Id:              r.Id.Hex(),
		UserId:          r.UserId.Hex(),
		AccommodationId: r.AccommodationId.Hex(),
		StartDate:       timestamppb.New(r.StartDate.Time()),
		EndDate:         timestamppb.New(r.StartDate.Time()),
		Status:          r.Status,
		Price:           r.Price,
	}
}
