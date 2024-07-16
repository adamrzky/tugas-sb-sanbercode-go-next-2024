package controllers

import (
	"net/http"
	"time"

	"backend-tugas-reactjs/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type footballerInput struct {
	Name        string `json:"name"`
	Nationality string `json:"nationality"`
	Age         int    `json:"age"`
}

// GetAllFootballer godoc
// @Summary Get all Footballer.
// @nationality Get a list of Footballer.
// @Tags Footballer
// @Produce json
// @Success 200 {object} []models.Footballer
// @Router /footballers [get]
func GetAllFootballer(c *gin.Context) {
	// get db from gin context
	db := c.MustGet("db").(*gorm.DB)
	var footballers []models.Footballer
	db.Find(&footballers)

	c.JSON(http.StatusOK, gin.H{"data": footballers})
}

// CreateFootballer godoc
// @Summary Create New Footballer.
// @nationality Creating a new Footballer.
// @Tags Footballer
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param Body body footballerInput true "the body to create a new Footballer"
// @Security BearerToken
// @Produce json
// @Success 200 {object} models.Footballer
// @Router /footballers [post]
func CreateFootballer(c *gin.Context) {
	// Validate input
	var input footballerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create footballer
	footballer := models.Footballer{Name: input.Name, Nationality: input.Nationality, Age: input.Age}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&footballer)

	c.JSON(http.StatusOK, gin.H{"data": footballer})
}

// GetFootballerById godoc
// @Summary Get Footballer.
// @nationality Get an Footballer by id.
// @Tags Footballer
// @Produce json
// @Param id path string true "Footballer id"
// @Success 200 {object} models.Footballer
// @Router /footballers/{id} [get]
func GetFootballerById(c *gin.Context) { // Get model if exist
	var footballer models.Footballer

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&footballer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": footballer})
}

// UpdateFootballer godoc
// @Summary Update Footballer.
// @nationality Update Footballer by id.
// @Tags Footballer
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "Footballer id"
// @Param Body body footballerInput true "the body to update footballer"
// @Security BearerToken
// @Success 200 {object} models.Footballer
// @Router /footballers/{id} [patch]
func UpdateFootballer(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var footballer models.Footballer
	if err := db.Where("id = ?", c.Param("id")).First(&footballer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input footballerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var updatedInput models.Footballer
	updatedInput.Name = input.Name
	updatedInput.Nationality = input.Nationality
	updatedInput.Age = input.Age
	updatedInput.UpdatedAt = time.Now()

	db.Model(&footballer).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": footballer})
}

// DeleteFootballer godoc
// @Summary Delete one Footballer.
// @nationality Delete a Footballer by id.
// @Tags Footballer
// @Produce json
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Param id path string true "Footballer id"
// @Security BearerToken
// @Success 200 {object} map[string]boolean
// @Router /footballers/{id} [delete]
func DeleteFootballer(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var footballer models.Footballer
	if err := db.Where("id = ?", c.Param("id")).First(&footballer).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&footballer)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
