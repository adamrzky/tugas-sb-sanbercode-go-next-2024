package controllers

import (
	"backend-tugas-reactjs/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func parseTime(input string) (time.Time, error) {
	// Tentukan tanggal default untuk digunakan bersama waktu
	defaultDate := "2000-01-01"
	layout := "2006-01-02T15:04"
	completeInput := fmt.Sprintf("%sT%s", defaultDate, input)

	parsedTime, err := time.Parse(layout, completeInput)
	if err != nil {
		return time.Time{}, fmt.Errorf("parsing time error: %w", err)
	}
	return parsedTime, nil
}

// GetAllJadwalKuliah godoc
// @Summary Get all Jadwal Kuliah
// @Description Get list of all Jadwal Kuliah
// @Tags JadwalKuliah
// @Accept  json
// @Produce  json
// @Success 200 {array} models.JadwalKuliah
// @Router /jadwal-kuliah [get]
// Function to get all schedule including Dosen and their MataKuliah
func GetAllJadwalKuliah(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var schedules []models.JadwalKuliah

	// Preload Dosen dan Mahasiswa beserta MataKuliah yang terkait
	if err := db.Preload("Dosen.MataKuliah").Preload("Mahasiswa").Find(&schedules).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": schedules})
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
	var input struct {
		models.JadwalKuliah
		JamMulai   string `json:"jamMulai"`
		JamSelesai string `json:"jamSelesai"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Parse JamMulai and JamSelesai
	startTime, err := parseTime(input.JamMulai)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start time format"})
		return
	}

	endTime, err := parseTime(input.JamSelesai)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end time format"})
		return
	}

	// Set parsed times in input.JadwalKuliah
	input.JadwalKuliah.JamMulai = startTime
	input.JadwalKuliah.JamSelesai = endTime

	db := c.MustGet("db").(*gorm.DB)
	result := db.Create(&input.JadwalKuliah)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input.JadwalKuliah})
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
