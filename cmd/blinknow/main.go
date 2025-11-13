package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Simple health check route
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Coinly API 🚀",
		})
	})

	// Example user route
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"pong": "coinly"})
	})

	// Start server on port 8080
	r.Run(":8080")
}
