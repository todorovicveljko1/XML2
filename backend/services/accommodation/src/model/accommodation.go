package model

import (
	"time"

	"acc.accommodation.com/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Accommodation struct {
	Id              primitive.ObjectID `bson:"_id" json:"id"`
	Amenity         []string           `bson:"amenity" json:"amenity"`
	DefaultPrice    float64            `bson:"default_price" json:"default_price"`
	Location        string             `bson:"location" json:"location"`
	MaxGuests       int                `bson:"max_guests" json:"max_guests"`
	MinGuests       int                `bson:"min_guests" json:"min_guests"`
	Name            string             `bson:"name" json:"name"`
	PhotoURL        []string           `bson:"photo_url" json:"photo_url"`
	UserId          primitive.ObjectID `bson:"user_id" json:"user_id"`
	IsPricePerNight bool               `bson:"is_price_per_night" json:"is_price_per_night"`
	IsManual        bool               `bson:"is_manual" json:"is_manual"`
	DeletedAt       *time.Time         `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`

	Price float64 `bson:"-" json:"price"`
}

// Convert to proto
func (a *Accommodation) ToProto() *pb.Accommodation {
	return &pb.Accommodation{
		Id:              a.Id.Hex(),
		Amenity:         a.Amenity,
		DefaultPrice:    a.DefaultPrice,
		Location:        a.Location,
		MaxGuests:       int32(a.MaxGuests),
		MinGuests:       int32(a.MinGuests),
		Name:            a.Name,
		PhotoUrl:        a.PhotoURL,
		UserId:          a.UserId.Hex(),
		IsPricePerNight: a.IsPricePerNight,
		IsManual:        a.IsManual,

		Price: &a.Price,
	}
}
