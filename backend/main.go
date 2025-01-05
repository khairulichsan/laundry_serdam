package main

import (
	"backend/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set up the Gin router
	r := gin.Default()
	// Enable CORS
	r.Use(cors.Default()) // This will allow all origins, you can customize it based on your needs
	// Register routes
	r.POST("/api/v1/login", controllers.LoginUser)
	r.POST("/api/v1/register", controllers.RegisterUser)
	// Other routes here...
	// Start the server
	r.Run(":8080")
}
