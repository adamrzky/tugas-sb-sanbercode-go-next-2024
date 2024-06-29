package config

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (*sql.DB, error) {
	dsn := "root:@tcp(localhost:3306)/quiz-1?parseTime=True"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
