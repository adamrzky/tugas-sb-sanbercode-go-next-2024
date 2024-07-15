package controllers

import (
	"backend-tugas-reactjs/models"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllBooks retrieves all books
func GetAllBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var books []models.Book
	if result := db.Find(&books); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// GetBookByID retrieves a single book by ID
func GetBookByID(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var book models.Book
	if result := db.First(&book, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	c.JSON(http.StatusOK, book)
}

// CreateBook creates a new book
func CreateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input models.BookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation
	errors := validateBook(input)
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	// Convert total_page to thickness
	thickness := convertThickness(input.TotalPage)

	book := models.Book{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		ReleaseYear: input.ReleaseYear,
		Price:       input.Price,
		TotalPage:   input.TotalPage,
		Thickness:   thickness,
	}

	if result := db.Create(&book); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

// UpdateBook updates an existing book
func UpdateBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	var book models.Book
	if result := db.First(&book, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Bind input
	var input models.BookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validation
	errors := validateBook(input)
	if len(errors) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
		return
	}

	// Convert total_page to thickness
	thickness := convertThickness(input.TotalPage)

	updatedData := models.Book{
		Title:       input.Title,
		Description: input.Description,
		ImageURL:    input.ImageURL,
		ReleaseYear: input.ReleaseYear,
		Price:       input.Price,
		TotalPage:   input.TotalPage,
		Thickness:   thickness,
	}

	db.Model(&book).Updates(updatedData)
	c.JSON(http.StatusOK, book)
}

// DeleteBook removes a book
func DeleteBook(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	id := c.Param("id")
	if result := db.Delete(&models.Book{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}

// validateBook checks the fields for book creation and update
func validateBook(input models.BookInput) map[string]string {
	errors := make(map[string]string)
	if _, err := url.ParseRequestURI(input.ImageURL); err != nil {
		errors["image_url"] = "The image URL must be valid"
	}
	if input.ReleaseYear < 1980 || input.ReleaseYear > 2021 {
		errors["release_year"] = "The release year must be between 1980 and 2021"
	}
	return errors
}

// convertThickness determines the book thickness based on page count
func convertThickness(totalPage int) string {
	switch {
	case totalPage <= 100:
		return "tipis"
	case totalPage <= 200:
		return "sedang"
	default:
		return "tebal"
	}
}
