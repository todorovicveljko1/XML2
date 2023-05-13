package accommodation

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

func AddAvailableIntervalHandler(ctx *gin.Context, clients *client.Clients) {
	var addAvailableIntervalRequest pb.AddAvailabilityRequest

	err := ctx.BindJSON(&addAvailableIntervalRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}

	/*
		TODO: Check if user is owner of accommodation
		userId, exists := ctx.Get("user")
		if !exists {
			ctx.AbortWithStatusJSON(401, gin.H{
				"error": "You are not authorized to perform this action",
			})
			return
		}
	*/

	_, err = clients.AccommodationClient.AddAccommodationAvailability(ctx, &addAvailableIntervalRequest)
	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Successfully added availability",
	})
}
