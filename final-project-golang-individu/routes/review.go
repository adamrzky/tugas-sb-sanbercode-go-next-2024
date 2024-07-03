package routes

import (
	"final-project-golang-individu/controllers"
	"final-project-golang-individu/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func ReviewRoutes(r *gin.Engine) {
	authMiddleware, err := middleware.AuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	review := r.Group("/reviews")
	review.Use(authMiddleware.MiddlewareFunc()) // Apply JWT middleware to all routes in this group
	{
		review.POST("/", controllers.CreateReview)
		review.GET("/", controllers.GetReviews)
	}
}
