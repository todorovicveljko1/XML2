package router

import (
	"net/http"

	"api.accommodation.com/src/client"
	"api.accommodation.com/src/handler/accommodation"
	"api.accommodation.com/src/handler/auth"
	"api.accommodation.com/src/middleware"
	"github.com/gin-gonic/gin"
)

func ApiRouter(r *gin.RouterGroup, clients *client.Clients) {
	// Health check
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello world")
	})

	r.POST("/auth/login", func(ctx *gin.Context) {
		auth.LoginHandler(ctx, clients)
	})

	r.POST("/auth/register", func(ctx *gin.Context) {
		auth.RegisterHandler(ctx, clients)
	})

	authGroup := r.Group("/")
	authGroup.Use(middleware.Authorized(clients))
	authGroup.GET("/auth", func(ctx *gin.Context) {
		auth.MeAuthHandler(ctx, clients)
	})

	hostGroup := authGroup.Group("/")
	hostGroup.Use(middleware.HasRole([]string{"H"}))

	hostGroup.POST("/accommodation", func(ctx *gin.Context) {
		accommodation.CreateAccommodationHandler(ctx, clients)
	})

}
