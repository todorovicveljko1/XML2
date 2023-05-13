package accommodation

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

func GetAccommodationHandler(ctx *gin.Context, clients *client.Clients) {

	id := ctx.Param("id")

	acc, err := clients.AccommodationClient.GetAccommodation(ctx, &pb.GetAccommodationRequest{
		Id: id,
	})

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, acc)

}
