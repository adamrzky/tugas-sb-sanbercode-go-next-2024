package routes

import (
	"final-project-golang-individu/controllers"
	"final-project-golang-individu/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func RestaurantRoutes(r *gin.Engine) {
	authMiddleware, err := middleware.AuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	restaurant := r.Group("/restaurants")
	restaurant.Use(authMiddleware.MiddlewareFunc()) // Apply JWT middleware to all routes in this group
	{
		restaurant.POST("/", controllers.CreateRestaurant)
		restaurant.GET("/", controllers.GetRestaurants)
	}
}
