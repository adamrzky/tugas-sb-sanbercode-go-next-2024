package controllers

import (
	"Tugas-13/models"
	"Tugas-13/queries"
	"Tugas-13/utils"
	"context"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// GetAllMahasiswa mengembalikan semua mahasiswa
func GetAllMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()                              // Mendapatkan context dari http.Request
	mahasiswas, err := queries.GetAllMahasiswa(ctx) // Menambahkan ctx sebagai argumen
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	utils.ResponseJSON(w, mahasiswas, http.StatusOK)
}

func GetMahasiswaByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	// Menggunakan context.Background(), bisa diganti dengan context yang lebih spesifik jika diperlukan
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mahasiswa, err := queries.GetMahasiswaByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ResponseJSON(w, mahasiswa, http.StatusOK)
}

// CreateMahasiswa menangani pembuatan mahasiswa
func CreateMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var mahasiswa models.Mahasiswa
	err := json.NewDecoder(r.Body).Decode(&mahasiswa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assuming that the InsertMahasiswa function returns an error if something goes wrong
	err = queries.InsertMahasiswa(mahasiswa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ResponseJSON(w, mahasiswa, http.StatusCreated)
}

// UpdateMahasiswa menangani pembaruan data mahasiswa
func UpdateMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var mahasiswa models.Mahasiswa
	err := json.NewDecoder(r.Body).Decode(&mahasiswa)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	updatedRows, err := queries.UpdateMahasiswa(id, mahasiswa)
	if err != nil {
		http.Error(w, "Failed to update mahasiswa: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if updatedRows == 0 {
		http.Error(w, "No mahasiswa found with provided ID", http.StatusNotFound)
		return
	}

	utils.ResponseJSON(w, map[string]string{"result": "success"}, http.StatusOK)
}

// DeleteMahasiswa menghapus mahasiswa berdasarkan ID
func DeleteMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	deletedRows, err := queries.DeleteMahasiswa(id)
	if err != nil {
		http.Error(w, "Failed to delete mahasiswa: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if deletedRows == 0 {
		http.Error(w, "No mahasiswa found with provided ID", http.StatusNotFound)
		return
	}

	utils.ResponseJSON(w, map[string]string{"result": "success"}, http.StatusOK)
}
