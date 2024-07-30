// Dosen.go
package models

import (
	"time"
)

type Dosen struct {
	ID           uint `gorm:"primarykey"`
	Nama         string
	MataKuliahID uint
	MataKuliah   MataKuliah `gorm:"foreignKey:MataKuliahID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Jadwals      []JadwalKuliah `gorm:"foreignKey:DosenID"`
}

func (Dosen) TableName() string {
	return "dosen"
}
