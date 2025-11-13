package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sidd-bash/blinknow-backend/internal/models"
	"gorm.io/gorm"
)

type UserHandler struct {
	DB *gorm.DB
}

// POST /user/profile
func (h *UserHandler) CompleteProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var body struct {
		Name    string `json:"name"`
		Email   string `json:"email"`
		Address string `json:"address"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data"})
		return
	}

	u := user.(models.User)
	h.DB.Model(&u).Updates(models.User{
		Name:              body.Name,
		Email:             body.Email,
		Address:           body.Address,
		IsProfileComplete: true,
	})

	c.JSON(http.StatusOK, gin.H{"message": "Profile completed successfully"})
}

// GET /user/profile
func (h *UserHandler) GetProfile(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	c.JSON(http.StatusOK, user)
}
