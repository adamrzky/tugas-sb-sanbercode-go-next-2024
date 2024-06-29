package queries

import (
	"Tugas-13/config"
	"Tugas-13/models"
	"context"
	"database/sql"
	"log"
	"time"
)

const layoutDateTime = time.RFC3339

func GetNilaiByMahasiswaID(ctx context.Context, db *sql.DB, mahasiswaID string) ([]models.Nilai, error) {
	var results []models.Nilai
	query := `SELECT n.id, n.indeks, n.skor, n.created_at, n.updated_at, m.id AS mahasiswa_id, m.nama AS mahasiswa_nama, 
                     mk.id AS mata_kuliah_id, mk.nama AS mata_kuliah_nama
              FROM nilai n
              JOIN mahasiswa m ON m.id = n.mahasiswa_id
              JOIN mata_kuliah mk ON mk.id = n.mata_kuliah_id
              WHERE m.id = ?`
	rows, err := db.QueryContext(ctx, query, mahasiswaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var n models.Nilai
		var createdAt, updatedAt string
		if err := rows.Scan(&n.ID, &n.Indeks, &n.Skor, &createdAt, &updatedAt, &n.MahasiswaID, &n.MahasiswaNama, &n.MataKuliahID, &n.MataKuliahNama); err != nil {
			return nil, err
		}

		n.CreatedAt, err = time.Parse(layoutDateTime, createdAt)
		if err != nil {
			log.Println("Error parsing created_at:", err)
			return nil, err
		}

		n.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Println("Error parsing updated_at:", err)
			return nil, err
		}

		results = append(results, n)
	}
	return results, nil
}

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

func InsertNilai(ctx context.Context, n models.Nilai) error {
	db, err := config.ConnectToMySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	query := "INSERT INTO nilai (indeks, skor, mata_kuliah_id, mahasiswa_id, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())"
	_, err = db.ExecContext(ctx, query, n.Indeks, n.Skor, n.MataKuliahID, n.MahasiswaID)
	return err
}

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
