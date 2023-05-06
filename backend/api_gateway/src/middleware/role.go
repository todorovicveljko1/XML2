package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleInRoles(role string, roles []string) bool {
	for _, v := range roles {
		if v == role {
			return true
		}
	}
	return false
}

func HasRole(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		value, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No role key present in JWT"})
			return
		}
		role := value.(string)
		if !RoleInRoles(role, roles) {

			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "You dont have premission"})
			return
		}
		c.Next()

	}
}
