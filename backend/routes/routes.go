// routes/routes.go
package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/register", controllers.RegisterUser)
		v1.POST("/login", controllers.LoginUser)
		v1.GET("/users", controllers.GetUsers)
		v1.GET("/orders", controllers.GetOrders)
		v1.GET("/services", controllers.GetServices)
		v1.POST("/orders", controllers.CreateOrder)
	}
}
