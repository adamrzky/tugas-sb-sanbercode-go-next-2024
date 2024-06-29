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

// CreateNilai handles the creation of a new nilai
func CreateNilai(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var nilai models.Nilai
	err := json.NewDecoder(r.Body).Decode(&nilai)
	if err != nil {
		http.Error(w, "Invalid input "+err.Error(), http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithCancel(r.Context())
	defer cancel()

	err = queries.InsertNilai(ctx, nilai)
	if err != nil {
		http.Error(w, "Failed to create nilai "+err.Error(), http.StatusInternalServerError)
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
