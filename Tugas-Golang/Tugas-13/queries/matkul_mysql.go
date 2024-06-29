package queries

import (
	"Tugas-13/config"
	"Tugas-13/models"
	"context"
	"log"
	"time"
)

// const layout = "2006-01-02 15:04:05"

func GetAllMataKuliah(ctx context.Context) ([]models.MataKuliah, error) {
	var mataKuliahs []models.MataKuliah
	db, err := config.ConnectToMySQL()
	if err != nil {
		log.Fatal("Cannot connect to MySQL", err)
	}
	defer db.Close()

	queryText := "SELECT id, nama, created_at, updated_at FROM mata_kuliah ORDER BY created_at DESC"
	rows, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Println("Error executing query GetAllMataKuliah:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var mk models.MataKuliah
		var createdAt, updatedAt string
		if err := rows.Scan(&mk.ID, &mk.Nama, &createdAt, &updatedAt); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		mk.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Println("Error parsing created_at:", err)
			return nil, err
		}

		mk.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Println("Error parsing updated_at:", err)
			return nil, err
		}

		mataKuliahs = append(mataKuliahs, mk)
	}
	return mataKuliahs, nil
}

// GetMataKuliahByID fetches a single mata kuliah record by ID
func GetMataKuliahByID(ctx context.Context, id string) (models.MataKuliah, error) {
	db, err := config.ConnectToMySQL()
	if err != nil {
		log.Println("Error connecting to MySQL:", err)
		return models.MataKuliah{}, err
	}
	defer db.Close()

	var mk models.MataKuliah
	query := "SELECT id, nama, created_at, updated_at FROM mata_kuliah WHERE id = ?"
	if err := db.QueryRowContext(ctx, query, id).Scan(&mk.ID, &mk.Nama, &mk.CreatedAt, &mk.UpdatedAt); err != nil {
		return models.MataKuliah{}, err
	}
	return mk, nil
}

// InsertMataKuliah adds a new mata kuliah to the database
func InsertMataKuliah(ctx context.Context, mk models.MataKuliah) error {
	db, err := config.ConnectToMySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO mata_kuliah (nama, created_at, updated_at) VALUES (?, NOW(), NOW())"
	_, err = db.ExecContext(ctx, query, mk.Nama)
	return err
}

// UpdateMataKuliah updates an existing mata kuliah record
func UpdateMataKuliah(ctx context.Context, id string, mk models.MataKuliah) error {
	db, err := config.ConnectToMySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "UPDATE mata_kuliah SET nama = ?, updated_at = NOW() WHERE id = ?"
	_, err = db.ExecContext(ctx, query, mk.Nama, id)
	return err
}

// DeleteMataKuliah removes a mata kuliah record from the database
func DeleteMataKuliah(ctx context.Context, id string) error {
	db, err := config.ConnectToMySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "DELETE FROM mata_kuliah WHERE id = ?"
	_, err = db.ExecContext(ctx, query, id)
	return err
}
