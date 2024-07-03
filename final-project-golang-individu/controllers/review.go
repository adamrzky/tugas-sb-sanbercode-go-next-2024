package controllers

import (
	"final-project-golang-individu/config"
	"final-project-golang-individu/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateReview handles the creation of a new review
func CreateReview(c *gin.Context) {
	var review models.Review
	if err := c.ShouldBindJSON(&review); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validasi rating
	if review.Rating < 1 || review.Rating > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Rating must be between 1 and 5"})
		return
	}

	// Simpan review ke database
	if err := config.DB.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, review)
}

// GetReviews handles fetching all reviews
func GetReviews(c *gin.Context) {
	var reviews []models.Review
	if err := config.DB.Preload("Comments").Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
