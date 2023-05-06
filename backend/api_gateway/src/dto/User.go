package dto

import "api.accommodation.com/pb"

type User struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	PlaceOfLiving string `json:"place_of_living"`
	Role          string `json:"role"`
}

// from pb.User to User
func UserFromProto(u *pb.User) *User {
	return &User{
		Id:            u.Id,
		Username:      u.Username,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Email:         u.Email,
		PlaceOfLiving: u.PlaceOfLiving,
		Role:          u.Role,
	}
}

// from User to pb.User
func (u *User) ToProto() *pb.User {
	return &pb.User{
		Id:            u.Id,
		Username:      u.Username,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Email:         u.Email,
		PlaceOfLiving: u.PlaceOfLiving,
		Role:          u.Role,
	}
}
