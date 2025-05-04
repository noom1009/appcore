package database

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var DB *sql.DB

func InitPostgres(dsn string) {
    var err error
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal("Cannot connect to Postgres:", err)
    }
    if err = DB.Ping(); err != nil {
        log.Fatal("Cannot ping Postgres:", err)
    }
}
func ConnectDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
		return nil, err
	}

	log.Println("✅ Connected to database successfully")
	return db, nil
}