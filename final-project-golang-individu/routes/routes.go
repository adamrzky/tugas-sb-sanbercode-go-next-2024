package routes

import (
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	AuthRoutes(r)
	UserRoutes(r)
	RestaurantRoutes(r)
	FoodRoutes(r)
	ReviewRoutes(r)
}
