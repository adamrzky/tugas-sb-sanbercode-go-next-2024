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
func GetAllMataKuliah(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()                               // Mendapatkan context dari http.Request
	mahasiswas, err := queries.GetAllMataKuliah(ctx) // Menambahkan ctx sebagai argumen
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	utils.ResponseJSON(w, mahasiswas, http.StatusOK)
}

func GetMataKuliahByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	ctx, cancel := context.WithCancel(r.Context()) // Gunakan context dari request
	defer cancel()

	mataKuliah, err := queries.GetMataKuliahByID(ctx, id) // Tambahkan ctx sebagai argumen
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ResponseJSON(w, mataKuliah, http.StatusOK)
}

func CreateMataKuliah(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var mataKuliah models.MataKuliah
	err := json.NewDecoder(r.Body).Decode(&mataKuliah)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(r.Context()) // Gunakan context dari request
	defer cancel()

	err = queries.InsertMataKuliah(ctx, mataKuliah) // Tambahkan ctx sebagai argumen
	if err != nil {
		http.Error(w, "Failed to create mata kuliah "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ResponseJSON(w, mataKuliah, http.StatusCreated)
}

// UpdateMataKuliah handles the updating of an existing mata kuliah
func UpdateMataKuliah(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	var mataKuliah models.MataKuliah
	err := json.NewDecoder(r.Body).Decode(&mataKuliah)
	if err != nil {
		http.Error(w, "Invalid input "+err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	err = queries.UpdateMataKuliah(ctx, id, mataKuliah)
	if err != nil {
		http.Error(w, "Failed to update mata kuliah "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ResponseJSON(w, mataKuliah, http.StatusOK)
}

// DeleteMataKuliah handles the deletion of a mata kuliah
func DeleteMataKuliah(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	err := queries.DeleteMataKuliah(ctx, id)
	if err != nil {
		http.Error(w, "Failed to delete mata kuliah "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ResponseJSON(w, map[string]string{"result": "success"}, http.StatusOK)
}
