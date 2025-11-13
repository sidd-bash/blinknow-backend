package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/sidd-bash/blinknow-backend/internal/handlers"
	"github.com/sidd-bash/blinknow-backend/internal/middleware"
	"github.com/sidd-bash/blinknow-backend/internal/services"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	twilioService := services.NewTwilioService()
	authHandler := handlers.NewAuthHandler(db, twilioService)
	userHandler := &handlers.UserHandler{DB: db}

	auth := r.Group("/auth")
	{
		auth.POST("/request-otp", authHandler.RequestOTP)
		auth.POST("/verify-otp", authHandler.VerifyOTP)
	}

	protected := r.Group("/user")
	protected.Use(middleware.AuthMiddleware(db))
	{
		protected.POST("/profile", userHandler.CompleteProfile)
		protected.GET("/profile", userHandler.GetProfile)
	}

	return r
}
