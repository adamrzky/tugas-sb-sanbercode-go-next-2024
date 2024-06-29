package queries

import (
	"Tugas-13/config"
	"Tugas-13/models"
	"context"
	"log"
	"time"
)

// const layoutDateTime = "2006-01-02 15:04:05"
// GetAllNilai fetches all records from the nilai table
func GetAllNilai(ctx context.Context) ([]models.Nilai, error) {
	var nilais []models.Nilai
	db, err := config.ConnectToMySQL()
	if err != nil {
		log.Fatalf("Cannot connect to MySQL: %v", err)
		return nil, err
	}
	defer db.Close()

	queryText := "SELECT id, skor, indeks, mata_kuliah_id, mahasiswa_id, created_at, updated_at FROM nilai ORDER BY created_at DESC"
	rows, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Printf("Error executing query GetAllNilai: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var n models.Nilai
		var createdAt, updatedAt string
		if err := rows.Scan(&n.ID, &n.Skor, &n.Indeks, &n.MataKuliahID, &n.MahasiswaID, &createdAt, &updatedAt); err != nil {
			log.Printf("Error scanning row: %v", err)
			return nil, err
		}

		// Parse the datetime fields
		n.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Printf("Error parsing created_at: %v", err)
			return nil, err
		}

		n.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Printf("Error parsing updated_at: %v", err)
			return nil, err
		}

		nilais = append(nilais, n)
	}
	return nilais, nil
}

// GetNilaiByID fetches a single nilai record by ID
func GetNilaiByID(ctx context.Context, id string) (models.Nilai, error) {
	db, err := config.ConnectToMySQL()
	if err != nil {
		log.Println("Error connecting to MySQL:", err)
		return models.Nilai{}, err
	}
	defer db.Close()

	var n models.Nilai
	query := "SELECT id, indeks, skor, created_at, updated_at FROM nilai WHERE id = ?"
	if err := db.QueryRowContext(ctx, query, id).Scan(&n.ID, &n.Indeks, &n.Skor, &n.CreatedAt, &n.UpdatedAt); err != nil {
		return models.Nilai{}, err
	}
	return n, nil
}

// InsertNilai adds a new nilai to the database
func InsertNilai(ctx context.Context, n models.Nilai) error {
	db, err := config.ConnectToMySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO nilai (indeks, skor, created_at, updated_at) VALUES (?, ?, NOW(), NOW())"
	_, err = db.ExecContext(ctx, query, n.Indeks, n.Skor)
	return err
}

// UpdateNilai updates an existing nilai record
func UpdateNilai(ctx context.Context, id string, n models.Nilai) error {
	db, err := config.ConnectToMySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "UPDATE nilai SET indeks = ?, skor = ?, updated_at = NOW() WHERE id = ?"
	_, err = db.ExecContext(ctx, query, n.Indeks, n.Skor, id)
	return err
}

// DeleteNilai removes a nilai record from the database
func DeleteNilai(ctx context.Context, id string) error {
	db, err := config.ConnectToMySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "DELETE FROM nilai WHERE id = ?"
	_, err = db.ExecContext(ctx, query, id)
	return err
}
