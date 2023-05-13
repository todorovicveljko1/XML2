package accommodation

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

func SearchAccommodationsHandler(ctx *gin.Context, clients *client.Clients) {

	var searchAccommodationsRequest pb.SearchRequest

	err := ctx.BindJSON(&searchAccommodationsRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}
	searchAccommodationsRequest.UserId = ""
	if searchAccommodationsRequest.ShowMy {
		userId, exists := ctx.Get("user")
		if !exists {
			ctx.AbortWithStatusJSON(400, gin.H{
				"error": "Invalid request",
			})
			return
		}
		searchAccommodationsRequest.UserId = userId.(string)
	}

	acc, err := clients.AccommodationClient.SearchAccommodations(ctx, &searchAccommodationsRequest)

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, acc)

}
