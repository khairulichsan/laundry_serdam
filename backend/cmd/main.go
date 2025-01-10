package main

import (
	"backend/config"
	"backend/routes"
	"log"

	"github.com/gin-contrib/cors"
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
	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.1", "10.0.0.1"})
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	// Load API routes
	routes.SetupRoutes(r)

	// Start server
	r.Run(":8080")
}
