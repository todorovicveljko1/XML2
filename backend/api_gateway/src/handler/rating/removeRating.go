package rating

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

func RemoveRatingHandler(ctx *gin.Context, clients *client.Clients) {

	reservationId := ctx.Param("id")
	// get user id from context
	userId, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	// get reservation by id
	reservation, err := clients.ReservationClient.GetReservation(ctx, &pb.GetReservationRequest{ReservationId: reservationId})

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	// get accommodation by id
	accommodation, err := clients.AccommodationClient.GetAccommodation(ctx, &pb.GetAccommodationRequest{Id: reservation.AccommodationId})
	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	if reservation.UserId != userId.(string) {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	// remove rating
	_, err = clients.RatingClient.RemoveRating(ctx, &pb.RemoveRatingRequest{
		Id:     accommodation.Accommodation.UserId,
		UserId: userId.(string),
	})
	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"message": "Rating removed"})
}
