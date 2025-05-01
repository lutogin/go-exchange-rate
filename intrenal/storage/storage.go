package storage

import (
	"database/sql"
	"fmt"
	"log"
)
import _ "github.com/mattn/go-sqlite3"

type Storage struct {
	Db *sql.DB
}

func New() *Storage {
	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", "cache.storage")
	if err != nil {
		log.Fatalf("Failed to connect to SQLite: %v", err)
	}

	// Create a table if it doesn't exist
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS cache (
		currency TEXT PRIMARY KEY,
		quotes TEXT,
		expiration DATETIME
	);`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	fmt.Println("SQLite initialized successfully!")

	return &Storage{
		Db: db,
	}
}
