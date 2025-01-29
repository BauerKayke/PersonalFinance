package config

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var DB *sql.DB

func Init(connStr string) *sql.DB {
	DB, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %q", err)
	}
	return DB

}
