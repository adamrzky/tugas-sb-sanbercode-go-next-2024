package models

import "time"

// Article represents the article data structure
type Article struct {
	ID            int       `json:"id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	ImageURL      string    `json:"image_url"`
	ArticleLength string    `json:"article_length"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	CategoryID    int       `json:"category_id"`
}
