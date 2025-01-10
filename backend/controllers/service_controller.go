// controllers/service_controller.go
package controllers

import (
	"backend/config"
	"backend/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetServices retrieves all services from the database
func GetServices(c *gin.Context) {
	var services []models.Service
	// Retrieve all services from the database
	if err := config.DB.Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to retrieve services"})
		return
	}

	// Return the services as a JSON response
	c.JSON(http.StatusOK, gin.H{"services": services})
}

// InputService allows the user to input a new service to the database
func InputService(c *gin.Context) {
	var service models.Service
	// Bind the incoming JSON to the service model
	if err := c.BindJSON(&service); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	// Save the service to the database
	if err := config.DB.Create(&service).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create service"})
		return
	}

	// Return success message
	c.JSON(http.StatusOK, gin.H{"message": "Service created successfully", "service": service})
}

// DeleteService deletes a service from the database by its ID
func DeleteService(c *gin.Context) {
	// Get the service ID from the URL parameters
	id := c.Param("id")
	log.Println("Received request to delete service with ID:", id)

	// Find the service by ID
	var service models.Service
	if err := config.DB.First(&service, id).Error; err != nil {
		log.Println("Service not found:", err)
		c.JSON(http.StatusNotFound, gin.H{"message": "Service not found"})
		return
	}

	// Delete the service from the database
	if err := config.DB.Delete(&service).Error; err != nil {
		log.Println("Error deleting service:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete service"})
		return
	}

	// Return success message
	log.Println("Service deleted successfully:", id)
	c.JSON(http.StatusOK, gin.H{"message": "Service deleted successfully"})
}
