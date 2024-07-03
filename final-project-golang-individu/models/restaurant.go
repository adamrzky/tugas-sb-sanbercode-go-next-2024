package models

import (
	"time"
)

type Restaurant struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Address   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Foods     []Food
	Reviews   []Review
}
