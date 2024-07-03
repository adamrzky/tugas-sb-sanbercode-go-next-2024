package routes

import (
	"final-project-golang-individu/controllers"
	"final-project-golang-individu/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	user := r.Group("/user")
	user.Use(middleware.GetAuthMiddleware())
	{
		user.GET("/profile", controllers.GetProfile)
	}
}
