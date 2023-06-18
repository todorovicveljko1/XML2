package rating

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

type CreateReservationRequest struct {
	HostRating          int32 `json:"host_rating" binding:"required"`
	AccommodationRating int32 `json:"accommodation_rating" binding:"required"`
}

func ModifyRatingHandler(ctx *gin.Context, clients *client.Clients) {

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

	// get request body
	var request CreateReservationRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"message": "Bad request"})
		return
	}

	// rate host
	_, err = clients.RatingClient.Rate(ctx, &pb.RateRequest{
		ReservationId:       reservationId,
		HostId:              accommodation.Accommodation.UserId,
		AccommodationId:     accommodation.Accommodation.Id,
		UserId:              userId.(string),
		HostRating:          request.HostRating,
		AccommodationRating: request.AccommodationRating,
	})

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{"message": "Rating successful added"})

}
