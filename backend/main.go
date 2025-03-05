package main

import (
	"backend/controllers"
	"backend/database"
	"backend/repositories"
	"backend/services"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func enableCORS() gin.HandlerFunc {
	// Cargar variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Obtener variables de entorno
	origins := os.Getenv("CORS_ORIGIN")
	methods := os.Getenv("CORS_METHODS")
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", origins)
		c.Writer.Header().Set("Access-Control-Allow-Methods", methods)
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	// Inicializar la base de datos
	database.Init()

	// Inicializar el repositorio, servicio y controlador
	stockRepo := repositories.NewStockRepository(database.DB)
	stockService := services.NewStockService(stockRepo)
	stockController := controllers.NewStockController(stockService)

	// Obtener datos de la API externa al iniciar la aplicación
	stocks, err := stockService.FetchStocksFromAPI()
	if err != nil {
		log.Fatalf("Error fetching stocks from API: %v", err)
	}

	// Guardar los datos en la base de datos
	if err := stockService.SaveStocks(stocks); err != nil {
		log.Fatalf("Error saving stocks to database: %v", err)
	}
	log.Println("Stocks fetched from API and saved to database successfully!")

	// Configurar el enrutador
	router := gin.Default()
	router.Use(enableCORS())
	router.GET("/stocks", stockController.GetAllStocks)
	router.GET("/recommendations", stockController.GetRecommendations)

	// Iniciar el servidor
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
