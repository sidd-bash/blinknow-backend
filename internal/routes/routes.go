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

	// 🔹 Services and handlers
	twilioService := services.NewTwilioService()
	authHandler := handlers.NewAuthHandler(db, twilioService)
	userHandler := &handlers.UserHandler{DB: db}
	productHandler := &handlers.ProductHandler{DB: db}

	// ✅ Public routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to Blinknow API 🚀"})
	})

	// ✅ Auth routes
	auth := r.Group("/auth")
	{
		auth.POST("/request-otp", authHandler.RequestOTP)
		auth.POST("/verify-otp", authHandler.VerifyOTP)
	}

	// ✅ Protected user routes
	user := r.Group("/user")
	user.Use(middleware.AuthMiddleware(db))
	{
		user.POST("/profile", userHandler.CompleteProfile)
		user.GET("/profile", userHandler.GetProfile)
	}

	// ✅ Protected product & category routes
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(db))
	{
		api.GET("/categories", productHandler.GetCategories)
		api.GET("/products", productHandler.GetProducts)
		api.GET("/products/:id", productHandler.GetProductByID)
	}

	return r
}
