package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProfile(c *gin.Context) {
	// Handle getting user profile
	c.JSON(http.StatusOK, gin.H{"message": "Get Profile"})
}
