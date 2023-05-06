package middleware

import (
	"net/http"
	"strings"

	"api.accommodation.com/pb"
	"api.accommodation.com/src/client"
	"github.com/gin-gonic/gin"
)

func Authorized(clients *client.Clients) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the authorization header
		authHeader := c.Request.Header.Get("Authorization")

		// Check that the authorization header is not empty
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is missing",
			})
			return
		}

		// Check that the authorization header starts with "Bearer"
		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || authHeaderParts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header format must be Bearer <token>",
			})
			return
		}

		// Extract the token from the authorization header
		tokenString := authHeaderParts[1]

		// Parse the token
		user, err := clients.AuthClient.AuthUser(c.Request.Context(), &pb.AuthUserRequest{Token: tokenString})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}
		// Set the token claims in the context for other handlers to use
		// TODO: Better structure
		c.Set("user", user.Id)
		c.Set("role", user.Role)
		// Call the next handler
		c.Next()
	}
}
