package notification

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

func GetNotificationHandler(ctx *gin.Context, clients *client.Clients) {
	// get user id from context
	userId, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(400, gin.H{
			"message": "user id not found",
		})
		return
	}

	res, err := clients.NotificationClient.GetNotifications(ctx, &pb.GetNotificationRequest{
		UserId: userId.(string),
		All:    false,
	})

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, res.Notifications)

}
