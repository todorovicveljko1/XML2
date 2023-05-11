package auth

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"github.com/gin-gonic/gin"
)

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func ChangePasswordHandler(ctx *gin.Context, clients *client.Clients) {
	var changePasswordRequest ChangePasswordRequest

	err := ctx.BindJSON(&changePasswordRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}

	userId, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "No user key present in JWT",
		})
		return
	}

	_, err = clients.AuthClient.ChangePassword(ctx.Request.Context(), &pb.ChangePasswordRequest{
		Id:          userId.(string),
		OldPassword: changePasswordRequest.OldPassword,
		NewPassword: changePasswordRequest.NewPassword,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Password changed",
	})
}
