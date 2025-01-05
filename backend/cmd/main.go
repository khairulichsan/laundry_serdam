package main

import (
	"backend/config"
	"backend/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Initialize database connection
	config.ConnectDatabase()

	log.Println("Server started on port 8080")
	// Create Gin router
	r := gin.Default()

	// Load API routes
	routes.SetupRoutes(r)

	// Start server
	r.Run(":8080")
}
