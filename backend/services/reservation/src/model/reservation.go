package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"reservation.accommodation.com/pb"
)

// STATUS:: PENDING, APPROVED, REJECTED, CANCELED, AUTO_REJECTED
type Reservation struct {
	Id              primitive.ObjectID `bson:"_id" json:"id"`
	UserId          primitive.ObjectID `bson:"user_id" json:"user_id"`
	HostId          primitive.ObjectID `bson:"host_id" json:"host_id"`
	AccommodationId primitive.ObjectID `bson:"accommodation_id" json:"accommodation_id"`
	StartDate       time.Time          `bson:"start_date" json:"start_date"`
	EndDate         time.Time          `bson:"end_date" json:"end_date"`
	Status          string             `bson:"status" json:"status"`
	Price           float64            `bson:"price" json:"price"`
	NumberOfGuests  int32              `bson:"number_of_guests" json:"number_of_guests"`
	DeletedAt       *time.Time         `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

// ConvertToPbReservation converts Reservation to pb.Reservation
func (r *Reservation) ConvertToPbReservation() *pb.Reservation {
	return &pb.Reservation{
		Id:              r.Id.Hex(),
		UserId:          r.UserId.Hex(),
		HostId:          r.HostId.Hex(),
		AccommodationId: r.AccommodationId.Hex(),
		StartDate:       r.StartDate.String(),
		EndDate:         r.EndDate.String(),
		Status:          r.Status,
		Price:           r.Price,
		NumberOfGuests:  r.NumberOfGuests,
	}
}
