package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Init() error {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Error loading .env file:", err)
	}

	dbURL := os.Getenv("DB_URL")
	var err error
	DB, err = sql.Open("pgx", dbURL)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	log.Println("Database connected successfully")
	return nil
}
