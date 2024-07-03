package models

import (
	"time"
)

type Review struct {
	ID           uint `gorm:"primaryKey"`
	UserID       uint `gorm:"not null"`
	RestaurantID uint `gorm:"not null"`
	Rating       int  `gorm:"not null"`
	Comment      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Comments     []Comment
}

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	ReviewID  uint   `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	Content   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
