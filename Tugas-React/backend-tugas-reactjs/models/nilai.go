package models

import (
	"time"
)

type Nilai struct {
	ID           uint       `gorm:"primarykey"`
	Indeks       string     `json:"indeks"`
	Skor         int        `json:"skor"`
	MahasiswaID  uint       `json:"mahasiswa_id"`
	MataKuliahID uint       `json:"mata_kuliah_id"`
	UsersID      uint       `json:"users_id"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	Mahasiswa    Mahasiswa  `json:"mahasiswa" gorm:"foreignKey:MahasiswaID"`
	MataKuliah   MataKuliah `json:"mata_kuliah" gorm:"foreignKey:MataKuliahID"`
	Users        User       `json:"users" gorm:"foreignKey:UsersID"`
}

// TableName overrides the table name used by User to `nilai`
func (Nilai) TableName() string {
	return "nilai"
}
