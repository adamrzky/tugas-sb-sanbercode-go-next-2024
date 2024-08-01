package models

import (
	"time"
)

type Mahasiswa struct {
	ID   uint `gorm:"primarykey"`
	Nama string
	// Email     string `gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Jadwals   []JadwalKuliah `gorm:"foreignKey:MahasiswaID"`
}

// TableName overrides the table name used by User to `mahasiswa`
func (Mahasiswa) TableName() string {
	return "mahasiswa"
}
