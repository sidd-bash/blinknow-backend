package services

import (
	"fmt"

	"github.com/sidd-bash/blinknow-backend/internal/models"
	"gorm.io/gorm"
)

func SeedData(db *gorm.DB) {
	var count int64
	db.Model(&models.Category{}).Count(&count)
	if count > 0 {
		fmt.Println("✅ Database already seeded")
		return
	}

	fmt.Println("🌱 Seeding initial categories and products...")

	categories := []models.Category{
		{Name: "Groceries"},
		{Name: "Beverages"},
		{Name: "Snacks"},
	}

	for _, cat := range categories {
		db.Create(&cat)
	}

	products := []models.Product{
		{Name: "Coca Cola", Price: 40, ImageURL: "https://example.com/coke.jpg", CategoryID: 2},
		{Name: "Potato Chips", Price: 25, ImageURL: "https://example.com/chips.jpg", CategoryID: 3},
		{Name: "Rice 5kg", Price: 299, ImageURL: "https://example.com/rice.jpg", CategoryID: 1},
	}

	for _, p := range products {
		db.Create(&p)
	}

	fmt.Println("✅ Seed completed!")
}
