package accommodation

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

type SearchAccommodationsRequest struct {
	Location  *string  `json:"location,omitempty"`
	NumGuests *int32   `json:"num_guests,omitempty"`
	StartDate *string  `json:"start_date,omitempty"`
	EndDate   *string  `json:"end_date,omitempty"`
	Amenity   []string `json:"amenity,omitempty"`
	ShowMy    bool     `json:"show_my,omitempty"`
}

// Conver to proto
func (s *SearchAccommodationsRequest) ToProto(userId string) *pb.SearchRequest {
	return &pb.SearchRequest{
		Location:  s.Location,
		NumGuests: s.NumGuests,
		StartDate: s.StartDate,
		EndDate:   s.EndDate,
		Amenity:   s.Amenity,
		UserId:    userId,
		ShowMy:    s.ShowMy,
	}
}

func SearchAccommodationsHandler(ctx *gin.Context, clients *client.Clients) {

	var searchAccommodationsRequest SearchAccommodationsRequest

	err := ctx.BindQuery(&searchAccommodationsRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}

	searchAccommodationsRequestProto := searchAccommodationsRequest.ToProto("")
	if searchAccommodationsRequest.ShowMy {
		userId, exists := ctx.Get("user")
		if !exists {
			ctx.AbortWithStatusJSON(400, gin.H{
				"error": "Invalid request",
			})
			return
		}
		searchAccommodationsRequestProto.UserId = userId.(string)
	}

	acc, err := clients.AccommodationClient.SearchAccommodations(ctx, searchAccommodationsRequestProto)

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, acc)

}
