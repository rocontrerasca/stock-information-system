package repositories

import (
	"backend/models"
	"database/sql"
	"sort"
	"strconv"
	"strings"
	"time"
)

type StockRepository struct {
	DB *sql.DB
}

func NewStockRepository(db *sql.DB) *StockRepository {
	return &StockRepository{DB: db}
}

func (r *StockRepository) SaveStocks(stocks []models.Stock) error {
	for _, stock := range stocks {
		_, err := r.DB.Exec(`
			INSERT INTO stock_data (ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
			ON CONFLICT (ticker) DO NOTHING`,
			stock.Ticker, stock.TargetFrom, stock.TargetTo, stock.Company, stock.Action, stock.Brokerage, stock.RatingFrom, stock.RatingTo, stock.Time,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *StockRepository) GetAllStocks() ([]models.Stock, error) {
	rows, err := r.DB.Query("SELECT * FROM stock_data")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []models.Stock
	for rows.Next() {
		var stock models.Stock
		if err := rows.Scan(&stock.Ticker, &stock.TargetFrom, &stock.TargetTo, &stock.Company, &stock.Action, &stock.Brokerage, &stock.RatingFrom, &stock.RatingTo, &stock.Time); err != nil {
			return nil, err
		}
		stocks = append(stocks, stock)
	}
	return stocks, nil
}

func (r *StockRepository) GetRecommendations() ([]models.Stock, error) {
	rows, err := r.DB.Query(`
        SELECT ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time
        FROM stock_data
        ORDER BY time DESC
    `)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stocks []models.Stock
	for rows.Next() {
		var stock models.Stock
		var timeString string // Variable temporal para almacenar el tiempo como string
		if err := rows.Scan(&stock.Ticker, &stock.TargetFrom, &stock.TargetTo, &stock.Company, &stock.Action, &stock.Brokerage, &stock.RatingFrom, &stock.RatingTo, &timeString); err != nil {
			return nil, err
		}

		// Parsear el tiempo desde el string
		stock.Time, err = time.Parse(time.RFC3339, timeString)
		if err != nil {
			return nil, err
		}

		stocks = append(stocks, stock)
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
	highCredibilityBrokerages := []string{"The Goldman Sachs Group", "Morgan Stanley", "J.P. Morgan"}
	for _, b := range highCredibilityBrokerages {
		if b == brokerage {
			return true
		}
	}
	return false
}
