package auth

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Token string `json:"token"`
}

// convert LoginRequest to pb.LoginRequest
func (r *LoginRequest) ToProto() *pb.LoginRequest {
	return &pb.LoginRequest{
		Username: r.Username,
		Password: r.Password,
	}
}

// convert pb.LoginResponse to LoginResponse
func LoginResponseFromProto(r *pb.LoginResponse) *LoginResponse {
	return &LoginResponse{
		Token: r.Token,
	}
}

func LoginHandler(ctx *gin.Context, clients *client.Clients) {
	var loginRequest LoginRequest

	err := ctx.BindJSON(&loginRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}

	res, err := clients.AuthClient.Login(ctx.Request.Context(), loginRequest.ToProto())

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, LoginResponseFromProto(res))
}
