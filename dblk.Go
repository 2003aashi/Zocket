package db

import (
    "database/sql"
    "log"

    _ "github.com/jackc/pgx/v4/stdlib" // PostgreSQL driver
)

var DB *sql.DB

func InitDB() {
    var err error
    DB, err = sql.Open("pgx", "postgres://postgres:secret@localhost:5432/postgres")
    if err != nil {
        log.Fatalf("Failed to connect to the database: %v", err)
    }

    if err := DB.Ping(); err != nil {
        log.Fatalf("Database is unreachable: %v", err)
    }

    log.Println("Database connection established.")
}
