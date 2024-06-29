package queries

import (
	"Tugas-13/config"
	"Tugas-13/models"
	"context"
	"database/sql"
	"log"
	"time"
)

// Format standar datetime MySQL
const layoutDateTime = "2006-01-02 15:04:05"

// GetMahasiswaByID fetches a single mahasiswa by ID from the database
func GetMahasiswaByID(ctx context.Context, id string) (models.Mahasiswa, error) {
	var mahasiswa models.Mahasiswa
	db, err := config.ConnectToMySQL()
	if err != nil {
		log.Println("Cannot connect to MySQL:", err)
		return mahasiswa, err
	}
	defer db.Close()

	queryText := "SELECT id, nama, created_at, updated_at FROM mahasiswa WHERE id = ?"
	row := db.QueryRowContext(ctx, queryText, id)

	var createdAt, updatedAt string
	if err := row.Scan(&mahasiswa.ID, &mahasiswa.Nama, &createdAt, &updatedAt); err != nil {
		if err == sql.ErrNoRows {
			return mahasiswa, nil // No rows error is not actually an error in this context
		}
		log.Println("Error scanning row:", err)
		return mahasiswa, err
	}

	// Parse the datetime fields
	mahasiswa.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
	if err != nil {
		log.Println("Error parsing created_at:", err)
		return mahasiswa, err
	}

	mahasiswa.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
	if err != nil {
		log.Println("Error parsing updated_at:", err)
		return mahasiswa, err
	}

	return mahasiswa, nil
}

// GetAllMahasiswa fetches all mahasiswa records from the database
func GetAllMahasiswa(ctx context.Context) ([]models.Mahasiswa, error) {
	var mahasiswas []models.Mahasiswa
	db, err := config.ConnectToMySQL()
	if err != nil {
		log.Fatal("Cannot connect to MySQL", err)
	}
	defer db.Close()

	queryText := "SELECT id, nama, created_at, updated_at FROM mahasiswa ORDER BY created_at DESC"
	rows, err := db.QueryContext(ctx, queryText)
	if err != nil {
		log.Println("Error executing query GetAllMahasiswa:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var m models.Mahasiswa
		var createdAt, updatedAt string // tempatkan tanggal sebagai string
		if err := rows.Scan(&m.ID, &m.Nama, &createdAt, &updatedAt); err != nil {
			log.Println("Error scanning row:", err)
			return nil, err
		}

		// Konversi string ke time.Time
		m.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Println("Error parsing created_at:", err)
			return nil, err
		}

		m.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Println("Error parsing updated_at:", err)
			return nil, err
		}

		mahasiswas = append(mahasiswas, m)
	}
	return mahasiswas, nil
}

func InsertMahasiswa(mahasiswa models.Mahasiswa) error {
	db, err := config.ConnectToMySQL() // Adjust with your actual function to get DB connection
	if err != nil {
		return err
	}
	defer db.Close()

	query := `INSERT INTO mahasiswa (nama, created_at, updated_at) VALUES (?, NOW(), NOW())`
	_, err = db.Exec(query, mahasiswa.Nama)
	if err != nil {
		return err
	}

	return nil
}
func UpdateMahasiswa(id string, m models.Mahasiswa) (int64, error) {
	db, err := config.ConnectToMySQL()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	query := "UPDATE mahasiswa SET nama = ?, updated_at = NOW() WHERE id = ?"
	result, err := db.Exec(query, m.Nama, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func DeleteMahasiswa(id string) (int64, error) {
	db, err := config.ConnectToMySQL()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	query := "DELETE FROM mahasiswa WHERE id = ?"
	result, err := db.Exec(query, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}