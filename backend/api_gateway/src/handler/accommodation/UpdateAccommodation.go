package accommodation

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

func UpdateAccommodationHandler(ctx *gin.Context, clients *client.Clients) {

	var UpdateAccommodationRequest pb.Accommodation

	err := ctx.BindJSON(&UpdateAccommodationRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}

	userId, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(401, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	accId := ctx.Param("id")

	UpdateAccommodationRequest.UserId = userId.(string)
	UpdateAccommodationRequest.Id = accId

	_, err = clients.AccommodationClient.UpdateAccommodation(ctx, &UpdateAccommodationRequest)

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Accommodation updated",
	})

}
