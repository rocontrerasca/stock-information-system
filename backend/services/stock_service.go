package services

import (
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ApiResponse struct {
	Items    []models.Stock `json:"items"`
	NextPage string         `json:"next_page"`
}

type StockService struct {
	StockRepo *repositories.StockRepository
}

func NewStockService(stockRepo *repositories.StockRepository) *StockService {
	return &StockService{StockRepo: stockRepo}
}

func (s *StockService) FetchStocksFromAPI() ([]models.Stock, error) {
	// Cargar variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Obtener variables de entorno
	apiURL := os.Getenv("API_URL")
	authToken := os.Getenv("API_TOKEN")

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+authToken)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var apiResponse ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	return apiResponse.Items, nil
}

func (s *StockService) SaveStocks(stocks []models.Stock) error {
	return s.StockRepo.SaveStocks(stocks)
}

func (s *StockService) GetAllStocks() ([]models.Stock, error) {
	return s.StockRepo.GetAllStocks()
}

func (s *StockService) GetRecommendations() ([]models.Stock, error) {
	return s.StockRepo.GetRecommendations()
}
