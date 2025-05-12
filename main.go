package main

import (
	"cp-raf-be/database"
    "cp-raf-be/internal/models"
    "cp-raf-be/routes"
	"cp-raf-be/validators"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	validators.InitValidator()


	database.Connect()
	database.DB.AutoMigrate(&models.Page{}, &models.Banner{})
	
	router := gin.Default()
	routes.SetupRoutes(router)

	router.Run(":8080")
}
