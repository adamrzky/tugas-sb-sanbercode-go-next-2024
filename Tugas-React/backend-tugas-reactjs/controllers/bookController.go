package controllers

import (
	"net/http"
	"time"

	"backend-tugas-reactjs/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type bookInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageURL    string `json:"image_url" binding:"required,url"`
	ReleaseYear int    `json:"release_year" binding:"required,gte=1980,lte=2021"`
	Price       string `json:"price" binding:"required"`
	TotalPage   int    `json:"total_page" binding:"required"`
}

// GetAllBooks godoc
// @Summary Get all books
// @Description Retrieve a list of books
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetAllBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var books []models.Book
	db.Find(&books)
	c.JSON(http.StatusOK, gin.H{"data": books})
}

// CreateBook godoc
// @Summary Create a new book
// @Description Add a new book to the library
// @Tags books
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert JWT token here" default(Bearer <add_token_here>)
// @Param book body bookInput true "Create book"
// @Success 200 {object} models.Book
// @Router /books [post]
func CreateBook(c *gin.Context) {
	var input bookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		ReleaseYear: input.ReleaseYear,
		Price:       input.Price,
		TotalPage:   input.TotalPage,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&book)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// GetBookByID godoc
// @Summary Get a book by ID
// @Description Get detailed information about a book
// @Tags books
// @Accept  json
// @Produce  json
// @Param id path int true "Book ID"
// @Success 200 {object} models.Book
// @Router /books/{id} [get]
func GetBookByID(c *gin.Context) {
	var book models.Book
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// UpdateBook godoc
// @Summary Update a book
// @Description Update book details
// @Tags books
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert JWT token here" default(Bearer <add_token_here>)
// @Param id path int true "Book ID"
// @Param book body bookInput true "Update book"
// @Success 200 {object} models.Book
// @Router /books/{id} [patch]
func UpdateBook(c *gin.Context) {
	var book models.Book
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input bookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateData := models.Book{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		ReleaseYear: input.ReleaseYear,
		Price:       input.Price,
		TotalPage:   input.TotalPage,
		UpdatedAt:   time.Now(),
	}

	db.Model(&book).Updates(updateData)
	c.JSON(http.StatusOK, gin.H{"data": book})
}

// DeleteBook godoc
// @Summary Delete a book
// @Description Remove a book from the library
// @Tags books
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Insert JWT token here" default(Bearer <add_token_here>)
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]bool "book deleted"
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	var book models.Book
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
