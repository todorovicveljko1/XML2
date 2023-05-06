package auth

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/dto"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Username      string `json:"username"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	PlaceOfLiving string `json:"place_of_living"`
	Role          string `json:"role"`
}

// RegisterRequest to pb.RegisterRequest
func (r *RegisterRequest) ToProto() *pb.RegisterRequest {
	return &pb.RegisterRequest{
		Username:      r.Username,
		FirstName:     r.FirstName,
		LastName:      r.LastName,
		Email:         r.Email,
		Password:      r.Password,
		PlaceOfLiving: r.PlaceOfLiving,
		Role:          r.Role,
	}
}

func RegisterHandler(ctx *gin.Context, clients *client.Clients) {
	var registerRequest RegisterRequest

	err := ctx.BindJSON(&registerRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}

	res, err := clients.AuthClient.Register(ctx.Request.Context(), registerRequest.ToProto())

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, dto.UserFromProto(res))
}
