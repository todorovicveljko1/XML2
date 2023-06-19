package router

import (
	"net/http"

	"api.accommodation.com/src/client"
	"api.accommodation.com/src/handler/accommodation"
	"api.accommodation.com/src/handler/auth"
	"api.accommodation.com/src/handler/rating"
	"api.accommodation.com/src/handler/reservation"
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

	r.GET("/accommodation/:id", func(ctx *gin.Context) {
		accommodation.GetAccommodationHandler(ctx, clients)
	})

	r.GET("/accommodation/:id/rating", func(ctx *gin.Context) {
		rating.GetRatingHandler(ctx, clients)
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

	authGroup.GET("/accommodation/:id/reservation/:reservation_id", func(ctx *gin.Context) {
		//accommodation.GetReservationHandler(ctx, clients)
		ctx.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
	})
	authGroup.PUT("/accommodation/:id/reservation/:reservation_id", func(ctx *gin.Context) {
		reservation.HandleReservationStatusChange(ctx, clients)
	})
	// ----------------- HOST ROUTES -----------------
	hostGroup := authGroup.Group("/")
	hostGroup.Use(middleware.HasRole([]string{"H"}))

	hostGroup.POST("/accommodation", func(ctx *gin.Context) {
		accommodation.CreateAccommodationHandler(ctx, clients)
	})

	hostGroup.PUT("/accommodation/:id", func(ctx *gin.Context) {
		accommodation.UpdateAccommodationHandler(ctx, clients)
	})

	authGroup.GET("/accommodation/:id/reservation", func(ctx *gin.Context) {
		reservation.GetReservationsForAccommodationHandler(ctx, clients)
	})

	hostGroup.PUT("/accommodation/:id/availability", func(ctx *gin.Context) {
		accommodation.AddAvailableIntervalHandler(ctx, clients)
	})
	hostGroup.PUT("/accommodation/:id/price", func(ctx *gin.Context) {
		accommodation.AddPriceIntervalHandler(ctx, clients)
	})

	// ----------------- GUEST ROUTES -----------------
	guestGroup := authGroup.Group("/")
	guestGroup.Use(middleware.HasRole([]string{"G"}))

	guestGroup.POST("/accommodation/:id/reservation", func(ctx *gin.Context) {
		reservation.CreateReservationHandler(ctx, clients)
		//ctx.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{"message": "Not implemented"})
	})

	guestGroup.GET("/reservation", func(ctx *gin.Context) {
		reservation.GetReservationsForGuestHandler(ctx, clients)
	})

	guestGroup.PUT("/reservation/:id/rating", func(ctx *gin.Context) {
		rating.ModifyRatingHandler(ctx, clients)
	})

	guestGroup.GET("/rating", func(ctx *gin.Context) {
		rating.MyRatingsHandler(ctx, clients)
	})

	guestGroup.DELETE("/reservation/:id/rating", func(ctx *gin.Context) {
		rating.RemoveRatingHandler(ctx, clients)
	})

}
