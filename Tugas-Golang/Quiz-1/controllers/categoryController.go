package controllers

import (
	"Quiz-1/config"
	"Quiz-1/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

var db *sql.DB

func init() {
	var err error
	db, err = config.GetDB()
	if err != nil {
		panic(err)
	}
}

// GetAllCategories fetches all categories from the database
func GetAllCategories(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rows, err := db.Query("SELECT id, name, created_at, updated_at FROM categories")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	categories := []models.Category{}
	for rows.Next() {
		var c models.Category
		if err := rows.Scan(&c.ID, &c.Name, &c.CreatedAt, &c.UpdatedAt); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		categories = append(categories, c)
	}

	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var cat models.Category
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	stmt, err := db.Prepare("INSERT INTO categories (name) VALUES (?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(cat.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, err := result.LastInsertId() // Dapatkan ID sebagai int64
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	cat.ID = int(id) // Konversi dari int64 ke int
	cat.CreatedAt = time.Now()
	cat.UpdatedAt = time.Now()

	json.NewEncoder(w).Encode(cat)
}

// UpdateCategory handles PUT requests to update a category
func UpdateCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid category ID", http.StatusBadRequest)
		return
	}
	var cat models.Category
	if err := json.NewDecoder(r.Body).Decode(&cat); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Example: Update category in a database
	cat.ID = id
	cat.UpdatedAt = time.Now()
	json.NewEncoder(w).Encode(cat)
}

// DeleteCategory handles DELETE requests to remove a category
func DeleteCategory(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))
	// Example: Delete category from a database
	fmt.Fprintf(w, "Category with ID %d has been deleted", id)
}
