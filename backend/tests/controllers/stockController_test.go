package controllers

import (
	"backend/controllers"
	"backend/repositories"
	"backend/services"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
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

func setupTestController() *controllers.StockController {
	db := setupTestDB()
	stockRepo := repositories.NewStockRepository(db)
	stockService := services.NewStockService(stockRepo)
	return controllers.NewStockController(stockService)
}

func TestGetStocks(t *testing.T) {
	controller := setupTestController()
	w := httptest.NewRecorder()

	// Configurar el router de Gin
	router := gin.Default()
	router.GET("/stocks", controller.GetAllStocks)

	// Crear una solicitud HTTP de prueba
	req, _ := http.NewRequest("GET", "/stocks", nil)

	// Ejecutar la solicitud
	router.ServeHTTP(w, req)

	// Verificar el código de estado y la respuesta
	assert.Equal(t, http.StatusOK, w.Code)
}
