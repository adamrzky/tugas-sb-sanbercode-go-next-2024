package models

import "time"

type Nilai struct {
	ID             int       `json:"id"`
	Indeks         string    `json:"indeks"`
	Skor           int       `json:"skor"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	MahasiswaID    int       `json:"mahasiswa_id"`
	MahasiswaNama  string    `json:"mahasiswa_nama,omitempty"`
	MataKuliahID   int       `json:"mata_kuliah_id"`
	MataKuliahNama string    `json:"mata_kuliah_nama,omitempty"`
}

type MahasiswaNilaiDetail struct {
	NamaMahasiswa  string
	NamaMataKuliah string
	Skor           int
	Indeks         string
}
type MahasiswaNilai struct {
	MahasiswaID int    `json:"mahasiswa_id"`
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	Skor        int    `json:"skor"`
	Indeks      string `json:"indeks"`
}
