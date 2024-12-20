package models

import (
	"time"
)

type Food struct {
	ID           uint   `gorm:"primaryKey"`
	RestaurantID uint   `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Name         string `gorm:"not null"`
	Description  string
	Price        float64 `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Restaurant   Restaurant `gorm:"foreignKey:RestaurantID"`
}
