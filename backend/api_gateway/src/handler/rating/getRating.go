package rating

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

func GetRatingHandler(ctx *gin.Context, clients *client.Clients) {
	// Get accommodation from id
	id := ctx.Param("id")
	acc, err := clients.AccommodationClient.GetAccommodation(ctx, &pb.GetAccommodationRequest{
		Id: id,
	})

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	// Get accommodation rating
	acc_rating, err := clients.RatingClient.AccommodationRating(ctx, &pb.RatingIdRequest{
		Id: acc.Accommodation.Id,
	})

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	// Get host rating
	host_rating, err := clients.RatingClient.HostRating(ctx, &pb.RatingIdRequest{
		Id: acc.Accommodation.UserId,
	})

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	// compose rating
	rating := make(map[string]float64)
	rating["accommodation"] = acc_rating.Rating
	rating["host"] = host_rating.Rating
	ctx.JSON(200, rating)
}
