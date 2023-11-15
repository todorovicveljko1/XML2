package rating

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

func MyRatingsHandler(ctx *gin.Context, clients *client.Clients) {

	userId, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(401, gin.H{"message": "Unauthorized"})
		return
	}
	// get ratings
	ratings, err := clients.RatingClient.GetMyRatings(ctx, &pb.RatingIdRequest{Id: userId.(string)})
	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, ratings)
}
