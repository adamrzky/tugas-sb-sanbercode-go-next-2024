package controllers

import (
	"Tugas-13/models"
	"Tugas-13/queries"
	"Tugas-13/utils"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// func GetNilaiByMahasiswaID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	id := ps.ByName("id")
// 	if id == "" {
// 		http.Error(w, "ID is required", http.StatusBadRequest)
// 		return
// 	}

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	results, err := queries.GetNilaiMahasiswa(ctx, id)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	utils.ResponseJSON(w, results, http.StatusOK)
// }

func GetAllMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()
	mahasiswas, err := queries.GetAllMahasiswa(ctx)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	utils.ResponseJSON(w, mahasiswas, http.StatusOK)
}

func GetMahasiswaByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	mahasiswa, err := queries.GetMahasiswaByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.ResponseJSON(w, mahasiswa, http.StatusOK)
}

func CreateMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var mahasiswa models.Mahasiswa
	err := json.NewDecoder(r.Body).Decode(&mahasiswa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = queries.InsertMahasiswa(mahasiswa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ResponseJSON(w, mahasiswa, http.StatusCreated)
}

func UpdateMahasiswa(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	log.Println("Update requested for ID:", id)

	var mahasiswa models.Mahasiswa
	err := json.NewDecoder(r.Body).Decode(&mahasiswa)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = queries.UpdateMahasiswa(r.Context(), id, mahasiswa)
	if err != nil {
		http.Error(w, "Failed to update mahasiswa: "+err.Error(), http.StatusInternalServerError)
		return
	}

	utils.ResponseJSON(w, map[string]string{"result": "success"}, http.StatusOK)
}

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
