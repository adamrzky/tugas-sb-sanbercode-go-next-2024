package models

import (
	"time"
)

type Dosen struct {
	ID         uint `gorm:"primarykey"`
	Nama       string
	MataKuliah string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Jadwals    []JadwalKuliah `gorm:"foreignKey:DosenID"`
}

// TableName overrides the table name used by User to `dosen`
func (Dosen) TableName() string {
	return "dosen"
}
