package controllers

import (
	"final-project-golang-individu/config"
	"final-project-golang-individu/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFood(c *gin.Context) {
	var food models.Food
	if err := c.ShouldBindJSON(&food); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&food).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, food)
}

func GetFoods(c *gin.Context) {
	var foods []models.Food
	if err := config.DB.Find(&foods).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, foods)
}
