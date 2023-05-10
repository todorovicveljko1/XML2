package model

import (
	"acc.accommodation.com/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Amenity int64

const (
	WIFI             Amenity = 0
	KITCHEN          Amenity = 1
	AIR_CONDITIONING Amenity = 2
	PARKING          Amenity = 3
)

type Accommodation struct {
	Id           primitive.ObjectID `bson:"_id" json:"id"`
	Amenity      []Amenity          `bson:"amenity" json:"amenity"`
	DefaultPrice float64            `bson:"default_price" json:"default_price"`
	Location     string             `bson:"location" json:"location"`
	MaxGuests    int                `bson:"max_guests" json:"max_guests"`
	MinGuests    int                `bson:"min_guests" json:"min_guests"`
	Name         string             `bson:"name" json:"name"`
	PhotoURL     []string           `bson:"photo_url" json:"photo_url"`
	UserId       primitive.ObjectID `bson:"user_id" json:"user_id"`
}

type Price struct {
	StartDate     primitive.Timestamp `bson:"start_date" json:"start_date"`
	EndDate       primitive.Timestamp `bson:"end_date" json:"end_date"`
	PricePerNight float64             `bson:"price_per_night" json:"price_per_night"`
}

type UnavailableInterval struct {
	StartDate primitive.Timestamp `bson:"start_date" json:"start_date"`
	EndDate   primitive.Timestamp `bson:"end_date" json:"end_date"`
}

// Convert to proto
func (a *Accommodation) ToProto() *pb.Accommodation {
	temp := make([]pb.Amenity, len(a.Amenity))
	for i, v := range a.Amenity {
		temp[i] = pb.Amenity(v)
	}

	return &pb.Accommodation{
		Id:           a.Id.Hex(),
		Amenity:      temp,
		DefaultPrice: a.DefaultPrice,
		Location:     a.Location,
		MaxGuests:    int32(a.MaxGuests),
		MinGuests:    int32(a.MinGuests),
		Name:         a.Name,
		PhotoUrl:     a.PhotoURL,
		UserId:       a.UserId.Hex(),
	}
}
