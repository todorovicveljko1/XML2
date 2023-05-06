package auth

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

func MeAuthHandler(ctx *gin.Context, clients *client.Clients) {
	userId, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "No user key present in JWT",
		})
		return
	}
	user, err := clients.AuthClient.GetUser(ctx.Request.Context(), &pb.GetUserRequest{Id: userId.(string)})
	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}
	ctx.JSON(200, user)
}
