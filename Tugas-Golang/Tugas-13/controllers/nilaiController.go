package controllers

import (
	"Tugas-13/config"
	"Tugas-13/models"
	"Tugas-13/queries"
	"Tugas-13/utils"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func GetNilaiByMahasiswaID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Mengambil ID mahasiswa dari parameter URL
	id := ps.ByName("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Membuat koneksi database
	db, err := config.ConnectToMySQL()
	if err != nil {
		http.Error(w, "Failed to connect to database: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close() // Pastikan untuk menutup koneksi database

	// Menggunakan context dengan timeout (opsional, tergantung kebutuhan)
	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	// Pemanggilan fungsi query yang mengambil data nilai berdasarkan ID mahasiswa
	results, err := queries.GetNilaiByMahasiswaID(ctx, db, id)
	if err != nil {
		http.Error(w, "Failed to retrieve data: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Mengirimkan hasil sebagai respons JSON
	utils.ResponseJSON(w, results, http.StatusOK)
}

// GetAllNilai mengembalikan semua nilai
func GetAllNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	nilai, err := queries.GetAllNilai(ctx)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	utils.ResponseJSON(w, nilai, http.StatusOK)
}

func CreateNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var nilai models.Nilai
	err := json.NewDecoder(r.Body).Decode(&nilai)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	err = queries.InsertNilai(ctx, nilai)
	if err != nil {
		http.Error(w, "Failed to create nilai: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ResponseJSON(w, nilai, http.StatusCreated)
}

// UpdateNilai handles the updating of an existing nilai
func UpdateNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var nilai models.Nilai
	err := json.NewDecoder(r.Body).Decode(&nilai)
	if err != nil {
		http.Error(w, "Invalid input "+err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	err = queries.UpdateNilai(ctx, id, nilai)
	if err != nil {
		http.Error(w, "Failed to update nilai "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ResponseJSON(w, nilai, http.StatusOK)
}

// DeleteNilai handles the deletion of a nilai
func DeleteNilai(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	err := queries.DeleteNilai(ctx, id)
	if err != nil {
		http.Error(w, "Failed to delete nilai "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ResponseJSON(w, map[string]string{"result": "success"}, http.StatusOK)
}
