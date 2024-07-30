package models

import (
	"time"
)

type JadwalKuliah struct {
	ID          uint `gorm:"primarykey"`
	Hari        string
	JamMulai    time.Time
	JamSelesai  time.Time
	DosenID     uint
	MahasiswaID uint
	Dosen       Dosen
	Mahasiswa   Mahasiswa
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// TableName overrides the table name used by User to `jadwal_kuliah`
func (JadwalKuliah) TableName() string {
	return "jadwal_kuliah"
}
