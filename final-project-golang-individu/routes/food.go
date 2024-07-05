package routes

import (
	"final-project-golang-individu/controllers"
	"final-project-golang-individu/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(r *gin.Engine) {
	authMiddleware, err := middleware.AuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	food := r.Group("/foods")
	food.GET("/", controllers.GetFoods)
	food.Use(authMiddleware.MiddlewareFunc()) // Apply JWT middleware to all routes in this group
	{
		food.POST("/", controllers.CreateFood)
	}
}
