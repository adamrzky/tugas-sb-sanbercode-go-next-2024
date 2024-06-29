package models

import "time"

type Nilai struct {
	ID           int       `json:"id"`
	Indeks       string    `json:"indeks"`
	Skor         int       `json:"skor"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	MataKuliahID int       `json:"mata_kuliah_id"`
	MahasiswaID  int       `json:"mahasiswa_id"`
}
