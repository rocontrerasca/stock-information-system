package services

import (
	"backend/models"
	"backend/repositories"
	"backend/services"
	"database/sql"
	"fmt"
	"log"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/assert"
)

type MockStockRepository struct {
	stocks []models.Stock
}

func (m *MockStockRepository) InsertStock(stock models.Stock) error {
	m.stocks = append(m.stocks, stock)
	return nil
}

func (m *MockStockRepository) GetAllStocks() ([]models.Stock, error) {
	return m.stocks, nil
}

func (m *MockStockRepository) GetRecommendedStocks() ([]models.Stock, error) {
	return m.stocks, nil
}

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

func TestFetchAndStoreStocks(t *testing.T) {
	db := setupTestDB()
	stockRepo := repositories.NewStockRepository(db)
	stockService := services.NewStockService(stockRepo)

	stocks, err := stockService.GetAllStocks()
	assert.NoError(t, err)
	assert.NotEmpty(t, stocks)
}
