package accommodation

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

func AddAvailableIntervalHandler(ctx *gin.Context, clients *client.Clients) {
	var addAvailableIntervalRequest pb.AddAvailabilityRequest

	err := ctx.BindJSON(&addAvailableIntervalRequest.Availability)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}
	addAvailableIntervalRequest.Id = ctx.Param("id")

	// check for active reservation in this interval
	res, err := clients.ReservationClient.HasActiveReservationInInterval(ctx, &pb.IntervalRequest{
		AccommodationId: addAvailableIntervalRequest.Id,
		StartDate:       addAvailableIntervalRequest.Availability.StartDate,
		EndDate:         addAvailableIntervalRequest.Availability.EndDate,
	})

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	if res.Value {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "There is an active reservation in this interval",
		})
		return
	}

	_, err = clients.AccommodationClient.AddAccommodationAvailability(ctx, &addAvailableIntervalRequest)
	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Successfully added availability",
	})
}
