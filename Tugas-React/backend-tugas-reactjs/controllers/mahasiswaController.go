package controllers

import (
	"backend-tugas-reactjs/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllMahasiswa godoc
// @Summary Get all Mahasiswa
// @Description Get list of all Mahasiswa
// @Tags Mahasiswa
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Mahasiswa
// @Router /mahasiswa [get]
func GetAllMahasiswa(c *gin.Context) {
	var mahasiswas []models.Mahasiswa
	db := c.MustGet("db").(*gorm.DB)
	db.Find(&mahasiswas)
	c.JSON(http.StatusOK, gin.H{"data": mahasiswas})
}

// CreateMahasiswa godoc
// @Summary Create new Mahasiswa
// @Description Create a new Mahasiswa
// @Tags Mahasiswa
// @Accept  json
// @Produce  json
// @Param mahasiswa body models.Mahasiswa true "Mahasiswa"
// @Success 200 {object} models.Mahasiswa
// @Router /mahasiswa [post]
func CreateMahasiswa(c *gin.Context) {
	var input models.Mahasiswa
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

// GetMahasiswaByID godoc
// @Summary Get Mahasiswa by ID
// @Description Get a Mahasiswa by ID
// @Tags Mahasiswa
// @Accept  json
// @Produce  json
// @Param id path int true "Mahasiswa ID"
// @Success 200 {object} models.Mahasiswa
// @Router /mahasiswa/{id} [get]
func GetMahasiswaByID(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": mahasiswa})
}

// UpdateMahasiswa godoc
// @Summary Update a Mahasiswa
// @Description Update an existing Mahasiswa
// @Tags Mahasiswa
// @Accept  json
// @Produce  json
// @Param id path int true "Mahasiswa ID"
// @Param mahasiswa body models.Mahasiswa true "Mahasiswa"
// @Success 200 {object} models.Mahasiswa
// @Router /mahasiswa/{id} [put]
func UpdateMahasiswa(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&mahasiswa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&mahasiswa)
	c.JSON(http.StatusOK, gin.H{"data": mahasiswa})
}

// DeleteMahasiswa godoc
// @Summary Delete a Mahasiswa
// @Description Delete a Mahasiswa by ID
// @Tags Mahasiswa
// @Accept  json
// @Produce  json
// @Param id path int true "Mahasiswa ID"
// @Success 200 {object} map[string]bool
// @Router /mahasiswa/{id} [delete]
func DeleteMahasiswa(c *gin.Context) {
	var mahasiswa models.Mahasiswa
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&mahasiswa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&mahasiswa)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
