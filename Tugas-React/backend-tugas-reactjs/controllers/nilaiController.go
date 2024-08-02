package controllers

import (
	"backend-tugas-reactjs/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllNilai godoc
// @Summary Get all Nilai
// @Description Get list of all Nilai with details
// @Tags Nilai
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Nilai
// @Router /nilai [get]
func GetAllNilai(c *gin.Context) {
	var nilai []models.Nilai
	db := c.MustGet("db").(*gorm.DB)
	db.Preload("Mahasiswa").Preload("MataKuliah").Preload("Users").Find(&nilai)
	c.JSON(http.StatusOK, gin.H{"data": nilai})
}

// GetNilaiByID godoc
// @Summary Get Nilai by ID
// @Description Get a Nilai by ID with details
// @Tags Nilai
// @Accept  json
// @Produce  json
// @Param id path int true "Nilai ID"
// @Success 200 {object} models.Nilai
// @Router /nilai/{id} [get]
func GetNilaiByID(c *gin.Context) {
	var nilai models.Nilai
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Preload("Mahasiswa").Preload("MataKuliah").Preload("Users").Where("id = ?", c.Param("id")).First(&nilai).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": nilai})
}

// CreateNilai godoc
// @Summary Create new Nilai
// @Description Create a new Nilai
// @Tags Nilai
// @Accept  json
// @Produce  json
// @Param nilai body models.Nilai true "Nilai"
// @Success 200 {object} models.Nilai
// @Router /nilai [post]
// CreateNilai godoc
// @Summary Create new Nilai
// @Description Create a new Nilai
// @Tags Nilai
// @Accept  json
// @Produce  json
// @Param nilai body models.Nilai true "Nilai"
// @Success 200 {object} models.Nilai
// @Router /nilai [post]
func CreateNilai(c *gin.Context) {
	var input models.Nilai
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ambil userID dari konteks
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID tidak ditemukan dalam konteks"})
		return
	}

	// Pastikan userID adalah integer
	id, ok := userID.(uint)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID tidak valid"})
		return
	}
	input.UsersID = id

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

// UpdateNilai godoc
// @Summary Update a Nilai
// @Description Update an existing Nilai
// @Tags Nilai
// @Accept  json
// @Produce  json
// @Param id path int true "Nilai ID"
// @Param nilai body models.Nilai true "Nilai"
// @Success 200 {object} models.Nilai
// @Router /nilai/{id} [put]
func UpdateNilai(c *gin.Context) {
	var nilai models.Nilai
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&nilai).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&nilai); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&nilai)
	c.JSON(http.StatusOK, gin.H{"data": nilai})
}

// DeleteNilai godoc
// @Summary Delete a Nilai
// @Description Delete a Nilai by ID
// @Tags Nilai
// @Accept  json
// @Produce  json
// @Param id path int true "Nilai ID"
// @Success 200 {object} map[string]bool
// @Router /nilai/{id} [delete]
func DeleteNilai(c *gin.Context) {
	var nilai models.Nilai
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&nilai).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&nilai)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
