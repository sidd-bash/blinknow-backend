package main

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
    "github.com/sidd-bash/blinknow-backend/internal/config"
    "github.com/sidd-bash/blinknow-backend/internal/models"
    "github.com/sidd-bash/blinknow-backend/internal/routes"
    "github.com/sidd-bash/blinknow-backend/internal/services"

)

func main() {
    // 🔹 Load environment variables once globally
    if err := godotenv.Load(); err != nil {
        fmt.Println("⚠️ No .env file found — using system environment variables")
    } else {
        fmt.Println("✅ Loaded .env successfully")
    }

    config.Init()

    config.DB.AutoMigrate(
        &models.User{},
        &models.Category{},
        &models.Product{},
    )
    services.SeedData(config.DB)


    r := routes.SetupRouter(config.DB)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Println("🚀 Starting blinknow backend on port", port)
    r.Run(":" + port)
}
