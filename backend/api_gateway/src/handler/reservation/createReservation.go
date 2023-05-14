package reservation

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"github.com/gin-gonic/gin"
)

type CreateReservationRequest struct {
	StartDate      string  `json:"start_date" binding:"required"`
	EndDate        string  `json:"end_date" binding:"required"`
	NumberOfGuests int32   `json:"number_of_guests" binding:"required"`
	Price          float64 `json:"price" binding:"required"`
}

func CreateReservationHandler(ctx *gin.Context, clients *client.Clients) {

	// get accommodation id from path
	accommodationId := ctx.Param("id")

	// get user id from context
	userId, exists := ctx.Get("user_id")
	if !exists {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	// get request body
	var request CreateReservationRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"message": "Bad request"})
		return
	}

	// Get accommodation
	accommodation, err := clients.AccommodationClient.GetAccommodation(ctx, &pb.GetAccommodationRequest{Id: accommodationId})
	if err != nil {
		ctx.AbortWithStatusJSON(404, gin.H{"message": "Accommodation not found"})
		return
	}

	// Create reservation
	reservation, err := clients.ReservationClient.CreateReservation(ctx, &pb.CreateReservationRequest{
		UserId:          userId.(string),
		AccommodationId: accommodationId,
		StartDate:       request.StartDate,
		EndDate:         request.EndDate,
		NumberOfGuests:  request.NumberOfGuests,
		Price:           request.Price,
	})
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{"message": "Error while creating reservation"})
		return
	}

	if !accommodation.Accommodation.IsManual {
		// Approve reservation
		_, err = clients.ReservationClient.ApproveReservation(ctx, &pb.GetReservationRequest{
			ReservationId: reservation.Id,
		})
		if err != nil {
			ctx.AbortWithStatusJSON(500, gin.H{"message": "Error while approving reservation"})
			return
		}
	}

	ctx.JSON(200, reservation)

}
