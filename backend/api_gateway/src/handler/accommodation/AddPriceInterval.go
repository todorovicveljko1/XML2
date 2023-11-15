package accommodation

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

func AddPriceIntervalHandler(ctx *gin.Context, clients *client.Clients) {

	var addPriceIntervalRequest pb.AddPriceRequest

	err := ctx.BindJSON(&addPriceIntervalRequest.Price)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}

	addPriceIntervalRequest.Id = ctx.Param("id")

	// check for active reservation in this interval
	res, err := clients.ReservationClient.HasActiveReservationInInterval(ctx, &pb.IntervalRequest{
		AccommodationId: addPriceIntervalRequest.Id,
		StartDate:       addPriceIntervalRequest.Price.StartDate,
		EndDate:         addPriceIntervalRequest.Price.EndDate,
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

	_, err = clients.AccommodationClient.AddAccommodationPrice(ctx, &addPriceIntervalRequest)

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return

	}
	ctx.JSON(200, gin.H{
		"message": "Successfully added price interval",
	})

}
