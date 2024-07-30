package controllers

import (
	"backend-tugas-reactjs/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllJadwalKuliah godoc
// @Summary Get all Jadwal Kuliah
// @Description Get list of all Jadwal Kuliah
// @Tags JadwalKuliah
// @Accept  json
// @Produce  json
// @Success 200 {array} models.JadwalKuliah
// @Router /jadwal-kuliah [get]
func GetAllJadwalKuliah(c *gin.Context) {
	var jadwals []models.JadwalKuliah
	db := c.MustGet("db").(*gorm.DB)
	db.Preload("Dosen").Preload("Mahasiswa").Find(&jadwals)
	c.JSON(http.StatusOK, gin.H{"data": jadwals})
}

// CreateJadwalKuliah godoc
// @Summary Create new Jadwal Kuliah
// @Description Create a new Jadwal Kuliah
// @Tags JadwalKuliah
// @Accept  json
// @Produce  json
// @Param jadwal body models.JadwalKuliah true "Jadwal Kuliah"
// @Success 200 {object} models.JadwalKuliah
// @Router /jadwal-kuliah [post]
func CreateJadwalKuliah(c *gin.Context) {
	var input models.JadwalKuliah
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&input)
	c.JSON(http.StatusOK, gin.H{"data": input})
}

// GetJadwalKuliahByID godoc
// @Summary Get Jadwal Kuliah by ID
// @Description Get a Jadwal Kuliah by ID
// @Tags JadwalKuliah
// @Accept  json
// @Produce  json
// @Param id path int true "Jadwal Kuliah ID"
// @Success 200 {object} models.JadwalKuliah
// @Router /jadwal-kuliah/{id} [get]
func GetJadwalKuliahByID(c *gin.Context) {
	var jadwal models.JadwalKuliah
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Preload("Dosen").Preload("Mahasiswa").Where("id = ?", c.Param("id")).First(&jadwal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": jadwal})
}

// UpdateJadwalKuliah godoc
// @Summary Update a Jadwal Kuliah
// @Description Update an existing Jadwal Kuliah
// @Tags JadwalKuliah
// @Accept  json
// @Produce  json
// @Param id path int true "Jadwal Kuliah ID"
// @Param jadwal body models.JadwalKuliah true "Jadwal Kuliah"
// @Success 200 {object} models.JadwalKuliah
// @Router /jadwal-kuliah/{id} [put]
func UpdateJadwalKuliah(c *gin.Context) {
	var jadwal models.JadwalKuliah
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&jadwal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	if err := c.ShouldBindJSON(&jadwal); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Save(&jadwal)
	c.JSON(http.StatusOK, gin.H{"data": jadwal})
}

// DeleteJadwalKuliah godoc
// @Summary Delete a Jadwal Kuliah
// @Description Delete a Jadwal Kuliah by ID
// @Tags JadwalKuliah
// @Accept  json
// @Produce  json
// @Param id path int true "Jadwal Kuliah ID"
// @Success 200 {object} map[string]bool
// @Router /jadwal-kuliah/{id} [delete]
func DeleteJadwalKuliah(c *gin.Context) {
	var jadwal models.JadwalKuliah
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&jadwal).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&jadwal)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
