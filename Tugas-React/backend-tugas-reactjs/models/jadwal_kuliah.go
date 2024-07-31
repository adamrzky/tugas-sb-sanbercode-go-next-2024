package models

import (
	"time"
)

type JadwalKuliah struct {
	ID          uint `gorm:"primarykey"`
	DosenID     uint
	MahasiswaID uint
	Hari        string
	JamMulai    string // Simpan sebagai 'HH:mm'
	JamSelesai  string // Simpan sebagai 'HH:mm'
	Dosen       Dosen
	Mahasiswa   Mahasiswa
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// TableName overrides the table name used by User to `jadwal_kuliah`
func (JadwalKuliah) TableName() string {
	return "jadwal_kuliah"
}
