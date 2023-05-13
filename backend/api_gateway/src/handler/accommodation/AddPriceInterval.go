package accommodation

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

func AddPriceIntervalHandler(ctx *gin.Context, clients *client.Clients) {

	var addPriceIntervalRequest pb.AddPriceRequest

	err := ctx.BindJSON(&addPriceIntervalRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid request",
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
