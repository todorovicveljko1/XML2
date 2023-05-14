package auth

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

func DeleteUserHandler(ctx *gin.Context, clients *client.Clients) {

	userId, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "No user key present in JWT",
		})
		return
	}
	// Get users role
	role, exists := ctx.Get("role")
	if !exists {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "No user role present in JWT",
		})
		return
	}
	if role == "G" {
		// Check if he has any reservations in future
		res, err := clients.ReservationClient.HasGuestActiveReservationInFuture(ctx, &pb.IdRequest{Id: userId.(string)})
		if err != nil {
			helper.PrettyGRPCError(ctx, err)
			return
		}

		if res.Value {
			ctx.AbortWithStatusJSON(400, gin.H{
				"error": "User has active reservations",
			})
			return
		}

	} else if role == "H" {
		// Get all accommodations for this host
		res, err := clients.AccommodationClient.GetAccommodationsForHost(ctx, &pb.GetAccommodationRequest{Id: userId.(string)})
		if err != nil {
			helper.PrettyGRPCError(ctx, err)
			return
		}
		// Check if any of them has active reservations
		var ids []string
		for _, acc := range res.Accommodations {
			ids = append(ids, acc.Id)
		}
		reservations, err := clients.ReservationClient.HasHostActiveReservationInFuture(ctx, &pb.IdList{Ids: ids})
		if err != nil {
			helper.PrettyGRPCError(ctx, err)
			return
		}
		if reservations.Value {
			ctx.AbortWithStatusJSON(400, gin.H{
				"error": "User has active reservations",
			})
			return
		}
		// If not delete all accommodations
		for _, acc := range res.Accommodations {
			_, err := clients.AccommodationClient.DeleteAccommodation(ctx, &pb.GetAccommodationRequestWithUser{Id: acc.Id, UserId: userId.(string)})
			if err != nil {
				helper.PrettyGRPCError(ctx, err)
				return
			}
		}
	}
	// TODO: Check if user does not have any reservations before deleting
	_, err := clients.AuthClient.DeleteUser(ctx.Request.Context(), &pb.GetUserRequest{Id: userId.(string)})
	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, gin.H{
		"message": "User deleted",
	})
}
