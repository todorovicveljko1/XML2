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
	Location       string `form:"location,omitempty"`
	NumGuests      int32  `form:"num_guests,omitempty"`
	StartDate      string `form:"start_date,omitempty"`
	EndDate        string `form:"end_date,omitempty"`
	Amenity        string `form:"amenity,omitempty"`
	ShowMy         bool   `form:"show_my,omitempty"`
	ShowSuperHosts bool   `form:"show_super_hosts,omitempty"`
	PriceMin       int32  `form:"price_min,omitempty"`
	PriceMax       int32  `form:"price_max,omitempty"`
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

	// filter out min and max price
	if searchAccommodationsRequest.PriceMax != 0 {
		var filteredAcc []*pb.Accommodation
		for _, a := range acc.Accommodations {
			if *a.Price >= float64(searchAccommodationsRequest.PriceMin) && *a.Price <= float64(searchAccommodationsRequest.PriceMax) {
				filteredAcc = append(filteredAcc, a)
			}
		}
		acc.Accommodations = filteredAcc
	}

	if len(acc.Accommodations) == 0 {
		ctx.JSON(200, acc)
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

		// filter for super hosts
		if searchAccommodationsRequest.ShowSuperHosts && len(filteredAcc) > 0 {
			// get only unique host ids from filtered accommodations same as adding to map and then getting keys
			var hostIds map[string]bool = make(map[string]bool)
			for _, a := range filteredAcc {
				hostIds[a.UserId] = true
			}
			var hostIdsSlice []string
			for k := range hostIds {
				hostIdsSlice = append(hostIdsSlice, k)
			}
			// Filter out super hosts
			temp, err := helper.FilterOutSuperHostIds(clients, hostIdsSlice)
			if err != nil {
				helper.PrettyGRPCError(ctx, err)
				return
			}

			var filteredAcc2 []*pb.Accommodation
			for _, a := range filteredAcc {
				for _, id := range temp {
					if a.UserId == id {
						filteredAcc2 = append(filteredAcc2, a)
						break
					}
				}
			}
			filteredAcc = filteredAcc2

		}

		acc.Accommodations = filteredAcc
	}

	ctx.JSON(200, acc)

}
