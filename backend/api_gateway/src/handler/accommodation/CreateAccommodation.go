package accommodation

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

type CreateAccommodationRequest struct {
	Amenity      []int32  `json:"amenity"`
	DefaultPrice float64  `json:"default_price"`
	Location     string   `json:"location"`
	MaxGuests    int32    `json:"max_guests"`
	MinGuests    int32    `json:"min_guests"`
	Name         string   `json:"name"`
	PhotoURL     []string `json:"photo_url"`
}

// Convert to proto
func (r *CreateAccommodationRequest) ToProto() *pb.AddAccommodationRequest {
	temp := make([]pb.Amenity, len(r.Amenity))
	for i, v := range r.Amenity {
		temp[i] = pb.Amenity(v)
	}

	return &pb.AddAccommodationRequest{
		Amenity:      temp,
		DefaultPrice: r.DefaultPrice,
		Location:     r.Location,
		MaxGuests:    r.MaxGuests,
		MinGuests:    r.MinGuests,
		Name:         r.Name,
		PhotoUrl:     r.PhotoURL,
	}
}

type CreateAccommodationResponse struct {
	Success bool `json:"success"`
}

// Convert from proto
func CreateAccommodationResponseFromProto(r *pb.AddAccommodationResponse) *CreateAccommodationResponse {
	return &CreateAccommodationResponse{
		Success: r.Success,
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

	md := metadata.New(map[string]string{"user": userId.(string)})
	new_ctx := metadata.NewOutgoingContext(ctx.Request.Context(), md)
	res, err := clients.AccommodationClient.AddAccommodation(new_ctx, createAccommodationRequest.ToProto())

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, CreateAccommodationResponseFromProto(res))

}
