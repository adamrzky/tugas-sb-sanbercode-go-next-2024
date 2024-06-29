package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Import driver MySQL
)

func ConnectToMySQL() (*sql.DB, error) {
	dsn := "root:@tcp(localhost:3306)/db_university"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
