package models

import (
	"time"
)

type (
	Footballer struct {
		ID          uint      `gorm:"primary_key" json:"id"`
		Name        string    `json:"name"`
		Nationality string    `json:"nationality"`
		Age         int       `json:"age"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
