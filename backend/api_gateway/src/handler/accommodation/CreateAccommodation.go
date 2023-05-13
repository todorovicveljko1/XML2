package accommodation

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/dto"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

type CreateAccommodationRequest struct {
	Amenity         []string `json:"amenity"`
	DefaultPrice    float64  `json:"default_price"`
	Location        string   `json:"location"`
	MaxGuests       int32    `json:"max_guests"`
	MinGuests       int32    `json:"min_guests"`
	Name            string   `json:"name"`
	PhotoURL        []string `json:"photo_url"`
	IsPricePerNight bool     `json:"is_price_per_night"`
	IsManual        bool     `json:"is_manual"`
}

// Convert to proto
func (r *CreateAccommodationRequest) ToProto(userId string) *pb.CreateAccommodationRequest {

	return &pb.CreateAccommodationRequest{
		Amenity:         r.Amenity,
		DefaultPrice:    r.DefaultPrice,
		Location:        r.Location,
		MaxGuests:       r.MaxGuests,
		MinGuests:       r.MinGuests,
		Name:            r.Name,
		PhotoUrl:        r.PhotoURL,
		IsPricePerNight: r.IsPricePerNight,
		IsManual:        r.IsManual,
		UserId:          userId,
	}
}

func CreateAccommodationHandler(ctx *gin.Context, clients *client.Clients) {

	var createAccommodationRequest CreateAccommodationRequest

	err := ctx.BindJSON(&createAccommodationRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}

	userId, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}

	_, err = clients.AccommodationClient.CreateAccommodation(ctx, createAccommodationRequest.ToProto(userId.(string)))

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, dto.ResMessage{Message: "Accommodation created"})

}
