package models

import (
	"time"
)

type MataKuliah struct {
	ID        uint `gorm:"primarykey"`
	Nama      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Dosens    []Dosen `gorm:"many2many:dosen_mata_kuliah"`
}

// TableName overrides the table name used by User to `mata_kuliah`
func (MataKuliah) TableName() string {
	return "mata_kuliah"
}
