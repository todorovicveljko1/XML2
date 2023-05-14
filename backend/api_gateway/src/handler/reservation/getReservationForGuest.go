package reservation

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"github.com/gin-gonic/gin"
)

func GetReservationsForGuestHandler(ctx *gin.Context, clients *client.Clients) {

	// get user id from context
	userId, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	// Get reservations
	reservations, err := clients.ReservationClient.GetReservationsForGuest(ctx, &pb.IdRequest{
		Id: userId.(string),
	})

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": "Error while getting reservations"})
		return
	}

	ctx.JSON(200, reservations)

}
