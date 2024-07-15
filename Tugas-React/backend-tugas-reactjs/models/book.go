package models

import (
	"time"
)

// Book represents the book model for database.
type Book struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Title       string    `gorm:"size:255" json:"title"`
	Description string    `gorm:"size:1000" json:"description"`
	ImageURL    string    `gorm:"size:1000" json:"image_url"`
	ReleaseYear int       `json:"release_year"`
	Price       string    `gorm:"size:100" json:"price"`
	TotalPage   int       `json:"total_page"`
	Thickness   string    `gorm:"size:100" json:"thickness"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// BookInput represents the data structure for input validation.
type BookInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageURL    string `json:"image_url" binding:"required,url"`
	ReleaseYear int    `json:"release_year" binding:"required,gte=1980,lte=2021"`
	Price       string `json:"price" binding:"required"`
	TotalPage   int    `json:"total_page" binding:"required"`
}
