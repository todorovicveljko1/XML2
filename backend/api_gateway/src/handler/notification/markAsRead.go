package notification

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"github.com/gin-gonic/gin"
)

func MarkAsReadHandler(ctx *gin.Context, clients *client.Clients) {
	// get user id from context
	userId, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(401, gin.H{
			"message": "user id not found",
		})
		return
	}

	// get notification id from param
	notificationId := ctx.Param("id")

	// mark notification as read
	_, err := clients.NotificationClient.MarkNotificationAsRead(ctx, &pb.MarkNotificationAsReadRequest{
		Id:     notificationId,
		UserId: userId.(string),
	})

	if err != nil {
		ctx.JSON(400, gin.H{
			"message": "can't mark notification as read",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "notification marked as read",
	})

}
