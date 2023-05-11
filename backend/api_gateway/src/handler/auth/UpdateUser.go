package auth

import (
	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"api.accommodation.com/src/helper"
	"github.com/gin-gonic/gin"
)

type UpdateUserRequest struct {
	Username      string `json:"username"`
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	Email         string `json:"email"`
	PlaceOfLiving string `json:"place_of_living"`
}

// Convert to pb user
func (u *UpdateUserRequest) ConvertToPbUser(id string, role string) *pb.User {
	return &pb.User{
		Id:            id,
		Username:      u.Username,
		FirstName:     u.FirstName,
		LastName:      u.LastName,
		Email:         u.Email,
		PlaceOfLiving: u.PlaceOfLiving,
		Role:          role,
	}
}

func UpdateUserHandler(ctx *gin.Context, clients *client.Clients) {
	var updateUserRequest UpdateUserRequest

	err := ctx.BindJSON(&updateUserRequest)

	if err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "Invalid request",
		})
		return
	}

	// Get user id from context
	userId, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "No user key present in JWT",
		})
		return
	}
	// Ger user role from context
	userRole, exists := ctx.Get("role")
	if !exists {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": "No role key present in JWT",
		})
		return
	}

	res, err := clients.AuthClient.UpdateUser(ctx.Request.Context(), updateUserRequest.ConvertToPbUser(userId.(string), userRole.(string)))

	if err != nil {
		helper.PrettyGRPCError(ctx, err)
		return
	}

	ctx.JSON(200, res)

}
