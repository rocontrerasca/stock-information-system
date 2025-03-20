package repositories

import (
	"backend/models"
	"database/sql"
	"fmt"
	"strings"
)

type StockRepository struct {
	DB *sql.DB
}

func NewStockRepository(db *sql.DB) *StockRepository {
	return &StockRepository{DB: db}
}

func (r *StockRepository) SaveStocks(stocks []models.Stock) error {
	if len(stocks) == 0 {
		return nil
	}

	query := `INSERT INTO stock_data (ticker, target_from, target_to, company, action, brokerage, rating_from, rating_to, time) 
	          VALUES `
	values := []interface{}{}
	placeholders := []string{}

	for i, stock := range stocks {
		start := i*9 + 1
		ph := fmt.Sprintf("($%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d, $%d)",
			start, start+1, start+2, start+3, start+4, start+5, start+6, start+7, start+8)
		placeholders = append(placeholders, ph)
		values = append(values, stock.Ticker, stock.TargetFrom, stock.TargetTo, stock.Company, stock.Action, stock.Brokerage, stock.RatingFrom, stock.RatingTo, stock.Time)
	}

	query += strings.Join(placeholders, ",")
	query += " ON CONFLICT DO NOTHING"

	_, err := r.DB.Exec(query, values...)
	return err
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
