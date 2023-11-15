package reservation

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

type ReservationStatusChangeRequest struct {
	Status string `json:"status"`
}

func HandleReservationStatusChange(ctx *gin.Context, clients *client.Clients) {

	var request ReservationStatusChangeRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"message": "Bad request"})
		return
	}
	// get accommodation id from path
	accommodationId := ctx.Param("id")
	// get reservation id from path
	reservationId := ctx.Param("reservation_id")
	// get user id from context and role
	userId, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}
	role, exists := ctx.Get("role")
	if !exists {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	// Get reservation
	reservation, err := clients.ReservationClient.GetReservation(ctx, &pb.GetReservationRequest{
		ReservationId: reservationId,
	})
	if err != nil {

		helper.PrettyGRPCError(ctx, err)
		return
	}

	// Get accommodation
	accommodation, err := clients.AccommodationClient.GetAccommodation(ctx, &pb.GetAccommodationRequest{
		Id: accommodationId,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(404, gin.H{"message": "Accommodation not found"})
		return
	}

	if role.(string) == "G" {
		if reservation.UserId != userId.(string) {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}
		if request.Status == "CANCELLED" {
			// Cancel reservation
			_, err := clients.ReservationClient.CancelReservation(ctx, &pb.GetReservationRequest{
				ReservationId: reservationId,
			})
			if err != nil {
				helper.PrettyGRPCError(ctx, err)
				return
			}
			// Send notification to host
			helper.SendNotification(ctx, clients, &helper.Notification{
				Type:       helper.ReservationCanceled,
				ResourceId: reservationId,
				Body:       "Reservation cancelled for accommodation" + accommodation.Accommodation.Name + " in period" + reservation.StartDate + " - " + reservation.EndDate + ".",
				UserId:     reservation.HostId,
			})

			ctx.JSON(200, gin.H{"message": "Reservation cancelled"})
			return
		}
	}

	if role.(string) == "H" {

		// Get accommodation
		accommodation, err := clients.AccommodationClient.GetAccommodation(ctx, &pb.GetAccommodationRequest{
			Id: accommodationId,
		})

		if err != nil {
			ctx.AbortWithStatusJSON(404, gin.H{"message": "Accommodation not found"})
			return
		}

		if accommodation.Accommodation.UserId != userId.(string) {
			ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
			return
		}

		if request.Status == "APPROVED" {
			// Approve reservation
			_, err := clients.ReservationClient.ApproveReservation(ctx, &pb.GetReservationRequest{
				ReservationId: reservationId,
			})
			if err != nil {

				helper.PrettyGRPCError(ctx, err)
				return
			}

			// Send notification to guest
			helper.SendNotification(ctx, clients, &helper.Notification{
				Type:       helper.ReservationAccepted,
				ResourceId: reservationId,
				Body:       "Reservation accepted for accommodation" + accommodation.Accommodation.Name + " in period" + reservation.StartDate + " - " + reservation.EndDate + ".",
				UserId:     reservation.UserId,
			})

			ctx.JSON(200, gin.H{"message": "Reservation approved"})
			return
		}

		if request.Status == "REJECTED" {
			// Reject reservation
			_, err := clients.ReservationClient.RejectReservation(ctx, &pb.GetReservationRequest{
				ReservationId: reservationId,
			})
			if err != nil {

				helper.PrettyGRPCError(ctx, err)
				return
			}

			// Send notification to guest
			helper.SendNotification(ctx, clients, &helper.Notification{
				Type:       helper.ReservationRejected,
				ResourceId: reservationId,
				Body:       "Reservation rejected for accommodation" + accommodation.Accommodation.Name + " in period" + reservation.StartDate + " - " + reservation.EndDate + ".",
				UserId:     reservation.UserId,
			})

			ctx.JSON(200, gin.H{"message": "Reservation rejected"})
			return
		}

	}

	ctx.AbortWithStatusJSON(400, gin.H{"message": "Bad request"})

}
