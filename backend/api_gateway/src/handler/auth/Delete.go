package auth

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(ctx *gin.Context, clients *client.Clients) {

	userId, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "No user key present in JWT",
		})
		return
	}

	// TODO: Check if user does not have any reservations before deleting
	_, err := clients.AuthClient.DeleteUser(ctx.Request.Context(), &pb.GetUserRequest{Id: userId.(string)})
	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "User deleted",
	})
}
