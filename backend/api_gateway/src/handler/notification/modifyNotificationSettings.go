package notification

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

func ModifyNotificationSettingsHandler(ctx *gin.Context, clients *client.Clients) {
	// get user id from context
	userId, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(401, gin.H{
			"message": "user id not found",
		})
		return
	}

	// get request body
	var req pb.ChangeNotifcationSettingsRequest = pb.ChangeNotifcationSettingsRequest{}
	if err := ctx.ShouldBindJSON(&req.Settings); err != nil {
		ctx.JSON(400, gin.H{
			"message": "invalid request body",
		})
		return
	}
	req.UserId = userId.(string)

	// modify notification settings
	_, err := clients.NotificationClient.ChangeNotifcationSettings(ctx, &req)

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "notification settings modified",
	})

}
