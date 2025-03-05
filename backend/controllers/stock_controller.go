package controllers

import (
	"net/http"

	"backend/services"

	"github.com/gin-gonic/gin"
)

type StockController struct {
	StockService *services.StockService
}

func NewStockController(stockService *services.StockService) *StockController {
	return &StockController{StockService: stockService}
}

func (c *StockController) GetAllStocks(ctx *gin.Context) {
	stocks, err := c.StockService.GetAllStocks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, stocks)
}

func (c *StockController) GetRecommendations(ctx *gin.Context) {
	recommendations, err := c.StockService.GetRecommendations()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, recommendations)
}
