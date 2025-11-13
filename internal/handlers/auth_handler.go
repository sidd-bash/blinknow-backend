package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sidd-bash/blinknow-backend/internal/models"
	"github.com/sidd-bash/blinknow-backend/internal/services"
	"gorm.io/gorm"
)







type AuthHandler struct {
	DB            *gorm.DB
	TwilioService *services.TwilioService
}

func NewAuthHandler(db *gorm.DB, twilio *services.TwilioService) *AuthHandler {
	return &AuthHandler{DB: db, TwilioService: twilio}
}

// POST /auth/request-otp
func (h *AuthHandler) RequestOTP(c *gin.Context) {
	var body struct {
		Phone string `json:"phone"`
	}
	if err := c.ShouldBindJSON(&body); err != nil || body.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number required"})
		return
	}

	err := h.TwilioService.SendOTP(body.Phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

// POST /auth/verify-otp
func (h *AuthHandler) VerifyOTP(c *gin.Context) {
	var body struct {
		Phone string `json:"phone"`
		OTP   string `json:"otp"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	valid, err := h.TwilioService.VerifyOTP(body.Phone, body.OTP)
	if err != nil || !valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	var user models.User
	result := h.DB.Where("phone = ?", body.Phone).First(&user)
	isNew := false
	if result.Error == gorm.ErrRecordNotFound {
		user = models.User{Phone: body.Phone}
		h.DB.Create(&user)
		isNew = true
	}

	token, _ := services.GenerateJWT(user.ID, user.Phone)

	c.JSON(http.StatusOK, gin.H{
		"token":               token,
		"is_new_user":         isNew,
		"is_profile_complete": user.IsProfileComplete,
	})
}

