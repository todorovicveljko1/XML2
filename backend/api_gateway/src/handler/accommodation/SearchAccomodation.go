package accommodation

import (
	"strings"

	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"api.accommodation.com/src/middleware"
	"github.com/gin-gonic/gin"
)

type SearchAccommodationsRequest struct {
	Location  string `form:"location,omitempty"`
	NumGuests int32  `form:"num_guests,omitempty"`
	StartDate string `form:"start_date,omitempty"`
	EndDate   string `form:"end_date,omitempty"`
	Amenity   string `form:"amenity,omitempty"`
	ShowMy    bool   `form:"show_my,omitempty"`
}

// Conver to proto
func (s *SearchAccommodationsRequest) ToProto(userId string) *pb.SearchRequest {
	return &pb.SearchRequest{
		Location:  &s.Location,
		NumGuests: &s.NumGuests,
		StartDate: &s.StartDate,
		EndDate:   &s.EndDate,
		Amenity:   strings.Split(s.Amenity, ","),
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
		userId, err := middleware.GetUserFromHeader(ctx, clients)
		if err != nil {
			ctx.AbortWithStatusJSON(400, gin.H{
				"error": "Invalid request",
			})
			return
		}
		searchAccommodationsRequestProto.UserId = *userId
	}
	acc, err := clients.AccommodationClient.SearchAccommodations(ctx, searchAccommodationsRequestProto)

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	if !searchAccommodationsRequest.ShowMy {
		var accIds []string
		for _, a := range acc.Accommodations {
			accIds = append(accIds, a.Id)
		}
		// Filter out taken accommodations
		list, err := clients.ReservationClient.FilterOutTakenAccommodations(ctx, &pb.FilterTakenAccommodationsRequest{
			AccommodationIds: accIds,
			StartDate:        searchAccommodationsRequest.StartDate,
			EndDate:          searchAccommodationsRequest.EndDate,
		})

		if err != nil {
			helper.PrettyGRPCError(ctx, err)
			return
		}
		// Filter out accommodations that are not in the list
		var filteredAcc []*pb.Accommodation
		for _, a := range acc.Accommodations {
			for _, id := range list.Ids {
				if a.Id == id {
					filteredAcc = append(filteredAcc, a)
					break
				}
			}
		}
		acc.Accommodations = filteredAcc
	}

	ctx.JSON(200, acc)

}
