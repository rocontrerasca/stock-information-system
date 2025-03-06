package repositories

import (
	"backend/models"
	"backend/repositories"
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	var DB *sql.DB

	// Obtener variables de entorno
	dbURL := "postgresql://root@localhost:26257/stocks?sslmode=disable"
	var err error
	DB, err = sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	fmt.Println("Connected to the database!")

	return DB
}

func TestInsertStock(t *testing.T) {
	db := setupTestDB()
	repo := *repositories.NewStockRepository(db)
	var stockTime, err = time.Parse(time.RFC3339, "2025-01-01T00:00:00Z")
	stocks := []models.Stock{
		{
			Ticker:     "TEST1",
			TargetFrom: "$10.00",
			TargetTo:   "$15.00",
			Company:    "Test Company 1",
			Action:     "Test Action 1",
			Brokerage:  "Test Brokerage 1",
			RatingFrom: "Test Rating From 1",
			RatingTo:   "Test Rating To 1",
			Time:       stockTime,
		},
		{
			Ticker:     "TEST2",
			TargetFrom: "$20.00",
			TargetTo:   "$25.00",
			Company:    "Test Company 2",
			Action:     "Test Action 2",
			Brokerage:  "Test Brokerage 2",
			RatingFrom: "Test Rating From 2",
			RatingTo:   "Test Rating To 2",
			Time:       stockTime,
		},
	}

	err = repo.SaveStocks(stocks)
	assert.NoError(t, err)
}

func TestGetAllStocks(t *testing.T) {
	db := setupTestDB()
	repo := *repositories.NewStockRepository(db)

	stocks, err := repo.GetAllStocks()
	assert.NoError(t, err)
	assert.NotEmpty(t, stocks)
}
