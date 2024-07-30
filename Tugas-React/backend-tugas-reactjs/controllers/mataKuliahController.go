package controllers

import (
	"backend-tugas-reactjs/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllMataKuliah godoc
// @Summary Get all Mata Kuliah
// @Description Get list of all Mata Kuliah
// @Tags MataKuliah
// @Accept  json
// @Produce  json
// @Success 200 {array} models.MataKuliah
// @Router /mata-kuliah [get]
func GetAllMataKuliah(c *gin.Context) {
	var mataKuliahs []models.MataKuliah
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&mataKuliahs)
	c.JSON(http.StatusOK, gin.H{"data": mataKuliahs})
}

// CreateMataKuliah godoc
// @Summary Create new Mata Kuliah
// @Description Create a new Mata Kuliah
// @Tags MataKuliah
// @Accept  json
// @Produce  json
// @Param mataKuliah body models.MataKuliah true "Mata Kuliah"
// @Success 200 {object} models.MataKuliah
// @Router /mata-kuliah [post]
func CreateMataKuliah(c *gin.Context) {
	var input models.MataKuliah
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

// GetMataKuliahByID godoc
// @Summary Get Mata Kuliah by ID
// @Description Get a Mata Kuliah by ID
// @Tags MataKuliah
// @Accept  json
// @Produce  json
// @Param id path int true "Mata Kuliah ID"
// @Success 200 {object} models.MataKuliah
// @Router /mata-kuliah/{id} [get]
func GetMataKuliahByID(c *gin.Context) {
	var mataKuliah models.MataKuliah
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&mataKuliah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": mataKuliah})
}

// UpdateMataKuliah godoc
// @Summary Update a Mata Kuliah
// @Description Update an existing Mata Kuliah
// @Tags MataKuliah
// @Accept  json
// @Produce  json
// @Param id path int true "Mata Kuliah ID"
// @Param mataKuliah body models.MataKuliah true "Mata Kuliah"
// @Success 200 {object} models.MataKuliah
// @Router /mata-kuliah/{id} [put]
func UpdateMataKuliah(c *gin.Context) {
	var mataKuliah models.MataKuliah
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&mataKuliah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&mataKuliah); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&mataKuliah)
	c.JSON(http.StatusOK, gin.H{"data": mataKuliah})
}

// DeleteMataKuliah godoc
// @Summary Delete a Mata Kuliah
// @Description Delete a Mata Kuliah by ID
// @Tags MataKuliah
// @Accept  json
// @Produce  json
// @Param id path int true "Mata Kuliah ID"
// @Success 200 {object} map[string]bool
// @Router /mata-kuliah/{id} [delete]
func DeleteMataKuliah(c *gin.Context) {
	var mataKuliah models.MataKuliah
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&mataKuliah).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&mataKuliah)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
