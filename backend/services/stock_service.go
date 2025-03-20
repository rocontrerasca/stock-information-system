package services

import (
	"backend/models"
	"backend/repositories"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"slices"

	"github.com/joho/godotenv"
)

var highCredibilityBrokerages = []string{"The Goldman Sachs Group", "Morgan Stanley", "J.P. Morgan"}

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
	stocks, err := s.StockRepo.GetAllStocks()
	if err != nil {
		return nil, err
	}
	// Algoritmo de recomendación mejorado
	type StockScore struct {
		Stock models.Stock
		Score float64
	}

	var scoredStocks []StockScore
	for _, stock := range stocks {
		score := 0.0

		// 1. Cambio en el precio objetivo
		targetFrom := parsePrice(stock.TargetFrom)
		targetTo := parsePrice(stock.TargetTo)
		if targetTo > targetFrom {
			score += 10 // Aumento en el precio objetivo
		} else if targetTo < targetFrom {
			score -= 5 // Disminución en el precio objetivo
		}

		// 2. Mejora en la calificación
		if stock.RatingTo == "Buy" && stock.RatingFrom != "Buy" {
			score += 20 // Mejora a "Buy"
		} else if stock.RatingTo == "Sell" && stock.RatingFrom != "Sell" {
			score -= 15 // Empeora a "Sell"
		}

		// 3. Tiempo de la recomendación (más reciente = mayor peso)
		timeWeight := time.Since(stock.Time).Hours()
		if timeWeight < 24 {
			score += 30 // Recomendación en las últimas 24 horas
		} else if timeWeight < 168 {
			score += 10 // Recomendación en la última semana
		}

		// 4. Brokerage de confianza
		if isHighCredibilityBrokerage(stock.Brokerage) {
			score += 15 // Brokerage de alta credibilidad
		}

		scoredStocks = append(scoredStocks, StockScore{Stock: stock, Score: score})
	}

	// Ordenar por puntuación (de mayor a menor)
	sort.Slice(scoredStocks, func(i, j int) bool {
		return scoredStocks[i].Score > scoredStocks[j].Score
	})

	// Devolver las 5 mejores recomendaciones
	var recommendations []models.Stock
	for i := 0; i < len(scoredStocks) && i < 5; i++ {
		recommendations = append(recommendations, scoredStocks[i].Stock)
	}

	return recommendations, nil
}

// Función para convertir el precio de string a float64
func parsePrice(price string) float64 {
	cleaned := strings.TrimPrefix(price, "$")
	value, _ := strconv.ParseFloat(cleaned, 64)
	return value
}

// Función para verificar si la correduría es de alta credibilidad
func isHighCredibilityBrokerage(brokerage string) bool {
	return slices.Contains(highCredibilityBrokerages, brokerage)
}
