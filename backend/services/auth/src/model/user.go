package model

import (
	"time"

	"auth.accommodation.com/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Role string

const (
	HOST  Role = "H"
	GUEST Role = "G"
)

type User struct {
	Id            primitive.ObjectID `bson:"_id" json:"id"`
	Username      string             `bson:"username" json:"username"`
	FirstName     string             `bson:"first_name" json:"first_name"`
	LastName      string             `bson:"last_name" json:"last_name"`
	Email         string             `bson:"email" json:"email"`
	Password      string             `bson:"password" json:"-"`
	PlaceOfLiving string             `bson:"place_of_living" json:"place_of_living"`
	Role          Role               `bson:"role" json:"role"`
	DeletedAt     *time.Time         `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
}

// ConvertToPbUser converts User to pb.User]
func (u *User) ConvertToPbUser() *pb.User {
	return &pb.User{
		Id:            u.Id.Hex(),
		Username:      u.Username,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Email:         u.Email,
		PlaceOfLiving: u.PlaceOfLiving,
		Role:          string(u.Role),
	}
}
