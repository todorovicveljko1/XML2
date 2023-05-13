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
	// ----------------- PUBLIC ROUTES -----------------
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello world")
	})

	r.POST("/auth/login", func(ctx *gin.Context) {
		auth.LoginHandler(ctx, clients)
	})

	r.POST("/auth/register", func(ctx *gin.Context) {
		auth.RegisterHandler(ctx, clients)
	})

	r.GET("/accommodation", func(ctx *gin.Context) {
		accommodation.SearchAccommodationsHandler(ctx, clients)
	})

	r.GET("/accommodation/{id}", func(ctx *gin.Context) {
		accommodation.GetAccommodationHandler(ctx, clients)
	})

	// ----------------- AUTH ROUTES -----------------
	authGroup := r.Group("/")
	authGroup.Use(middleware.Authorized(clients))
	authGroup.GET("/auth", func(ctx *gin.Context) {
		auth.MeAuthHandler(ctx, clients)
	})

	authGroup.PUT("/auth", func(ctx *gin.Context) {
		auth.UpdateUserHandler(ctx, clients)
	})
	authGroup.PUT("/auth/change-password", func(ctx *gin.Context) {
		auth.ChangePasswordHandler(ctx, clients)
	})

	authGroup.DELETE("/auth", func(ctx *gin.Context) {
		auth.DeleteUserHandler(ctx, clients)
	})

	authGroup.GET("/accommodation/{id}/reservation", func(ctx *gin.Context) {
		//accommodation.GetReservationHandler(ctx, clients)
		ctx.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
	})
	authGroup.GET("/accommodation/{id}/reservation/{reservation_id}", func(ctx *gin.Context) {
		//accommodation.GetReservationHandler(ctx, clients)
		ctx.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
	})
	authGroup.PUT("/accommodation/{id}/reservation/{reservation_id}", func(ctx *gin.Context) {
		//accommodation.UpdateReservationHandler(ctx, clients)
		ctx.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
	})

	// ----------------- HOST ROUTES -----------------
	hostGroup := authGroup.Group("/")
	hostGroup.Use(middleware.HasRole([]string{"H"}))

	hostGroup.POST("/accommodation", func(ctx *gin.Context) {
		accommodation.CreateAccommodationHandler(ctx, clients)
	})

	hostGroup.PUT("/accommodation/{id}", func(ctx *gin.Context) {
		//accommodation.UpdateAccommodationHandler(ctx, clients)
		ctx.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
	})

	hostGroup.PUT("/accommodation/{id}/unavailability", func(ctx *gin.Context) {
		//accommodation.UpdateAccommodationUnvailabilityHandler(ctx, clients)
		ctx.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
	})
	hostGroup.PUT("/accommodation/{id}/price", func(ctx *gin.Context) {
		//accommodation.UpdateAccommodationPriceHandler(ctx, clients)
		ctx.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
	})

	// ----------------- GUEST ROUTES -----------------
	guestGroup := authGroup.Group("/")
	guestGroup.Use(middleware.HasRole([]string{"G"}))

	guestGroup.POST("/accommodation/{id}/reservation", func(ctx *gin.Context) {
		//accommodation.CreateReservationHandler(ctx, clients)
		ctx.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
	})

}
