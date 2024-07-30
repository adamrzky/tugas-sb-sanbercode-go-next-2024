package controllers

import (
	"backend-tugas-reactjs/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllDosen godoc
// @Summary Get all Dosen
// @Description Get list of all Dosen
// @Tags Dosen
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Dosen
// @Router /dosen [get]
func GetAllDosen(c *gin.Context) {
	var dosens []models.Dosen
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&dosens)
	c.JSON(http.StatusOK, gin.H{"data": dosens})
}

// CreateDosen godoc
// @Summary Create new Dosen
// @Description Create a new Dosen
// @Tags Dosen
// @Accept  json
// @Produce  json
// @Param dosen body models.Dosen true "Dosen"
// @Success 200 {object} models.Dosen
// @Router /dosen [post]
func CreateDosen(c *gin.Context) {
	var input models.Dosen
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

// GetDosenByID godoc
// @Summary Get Dosen by ID
// @Description Get a Dosen by ID
// @Tags Dosen
// @Accept  json
// @Produce  json
// @Param id path int true "Dosen ID"
// @Success 200 {object} models.Dosen
// @Router /dosen/{id} [get]
func GetDosenByID(c *gin.Context) {
	var dosen models.Dosen
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&dosen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": dosen})
}

// UpdateDosen godoc
// @Summary Update a Dosen
// @Description Update an existing Dosen
// @Tags Dosen
// @Accept  json
// @Produce  json
// @Param id path int true "Dosen ID"
// @Param dosen body models.Dosen true "Dosen"
// @Success 200 {object} models.Dosen
// @Router /dosen/{id} [put]
func UpdateDosen(c *gin.Context) {
	var dosen models.Dosen
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&dosen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&dosen); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&dosen)
	c.JSON(http.StatusOK, gin.H{"data": dosen})
}

// DeleteDosen godoc
// @Summary Delete a Dosen
// @Description Delete a Dosen by ID
// @Tags Dosen
// @Accept  json
// @Produce  json
// @Param id path int true "Dosen ID"
// @Success 200 {object} map[string]bool
// @Router /dosen/{id} [delete]
func DeleteDosen(c *gin.Context) {
	var dosen models.Dosen
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&dosen).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&dosen)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
