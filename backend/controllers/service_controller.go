// controllers/service_controller.go
package controllers

import (
	"backend/config"
	"backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServices(c *gin.Context) {
	var services []models.Service
	if err := config.DB.Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to get services"})
		return
	}

	c.JSON(http.StatusOK, services)
}
