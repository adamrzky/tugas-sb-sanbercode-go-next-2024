package routes

import (
	"final-project-golang-individu/controllers"
	"final-project-golang-individu/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	authMiddleware, err := middleware.AuthMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	auth := r.Group("/auth")
	{
		auth.POST("/register", controllers.Register)
		auth.POST("/login", authMiddleware.LoginHandler)
		auth.POST("/change-password", authMiddleware.MiddlewareFunc(), controllers.ChangePassword)
	}
}
