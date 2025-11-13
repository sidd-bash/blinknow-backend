package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sidd-bash/blinknow-backend/internal/models"
	"gorm.io/gorm"
)

type ProductHandler struct {
	DB *gorm.DB
}

// 🔹 GET /api/categories
func (h *ProductHandler) GetCategories(c *gin.Context) {
	var categories []models.Category
	if err := h.DB.Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"categories": categories})
}

// 🔹 GET /api/products
func (h *ProductHandler) GetProducts(c *gin.Context) {
	var products []models.Product
	if err := h.DB.Preload("Category").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}

// 🔹 GET /api/products/:id
func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := h.DB.Preload("Category").First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"product": product})
}
