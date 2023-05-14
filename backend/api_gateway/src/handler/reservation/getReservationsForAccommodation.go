package reservation

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"github.com/gin-gonic/gin"
)

func GetReservationsForAccommodationHandler(ctx *gin.Context, clients *client.Clients) {

	// get accommodation id from path
	accommodationId := ctx.Param("id")

	// Get reservations
	reservations, err := clients.ReservationClient.GetReservationsForAccommodation(ctx, &pb.IdRequest{
		Id: accommodationId,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": "Error while getting reservations"})
		return
	}

	ctx.JSON(200, reservations)
}
