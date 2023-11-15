package accommodation

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

type Accommodation struct {
	Id              string   `json:"id,omitempty"`
	Name            string   `json:"name,omitempty"`
	Location        string   `json:"location,omitempty"`
	Amenity         []string `json:"amenity,omitempty"`
	PhotoUrl        []string `json:"photo_url,omitempty"`
	MaxGuests       int32    `json:"max_guests,omitempty"`
	MinGuests       int32    `json:"min_guests,omitempty"`
	DefaultPrice    float64  `json:"default_price,omitempty"`
	UserId          string   `json:"user_id,omitempty"`
	IsPricePerNight bool     `json:"is_price_per_night,omitempty"`
	IsManual        bool     `json:"is_manual,omitempty"`
	Price           float64  `json:"price,omitempty"`
	IsSuperHost     bool     `json:"is_super_host,omitempty"`
}

type GetAccommodationResponse struct {
	Accommodation      Accommodation           `json:"accommodation,omitempty"`
	AvailableIntervals []*pb.AvailableInterval `json:"available_intervals,omitempty"`
	PriceIntervals     []*pb.PriceInterval     `json:"price_intervals,omitempty"`
}

// pb.Accommodation -> Accommodation
func (a *Accommodation) FromProto(acc *pb.Accommodation) {
	a.Id = acc.Id
	a.Name = acc.Name
	a.Location = acc.Location
	a.Amenity = acc.Amenity
	a.PhotoUrl = acc.PhotoUrl
	a.MaxGuests = acc.MaxGuests
	a.MinGuests = acc.MinGuests
	a.DefaultPrice = acc.DefaultPrice
	a.UserId = acc.UserId
	a.IsPricePerNight = acc.IsPricePerNight
	a.IsManual = acc.IsManual
}

func GetAccommodationHandler(ctx *gin.Context, clients *client.Clients) {

	id := ctx.Param("id")

	acc, err := clients.AccommodationClient.GetAccommodation(ctx, &pb.GetAccommodationRequest{
		Id: id,
	})
	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	// Check if owner is super host
	isSuperHost, err := helper.IsSuperHost(clients, acc.Accommodation.UserId)
	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	// Create response
	response := GetAccommodationResponse{
		Accommodation:      Accommodation{},
		AvailableIntervals: acc.AvailableIntervals,
		PriceIntervals:     acc.PriceIntervals,
	}

	// Convert to Accommodation
	response.Accommodation.FromProto(acc.Accommodation)
	response.Accommodation.IsSuperHost = isSuperHost

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, response)

}
