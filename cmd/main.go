package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/cristianrubioa/stockwise/db"
	"github.com/cristianrubioa/stockwise/internal/api"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	// Define a CLI flag for the server port (default: 8088)
	port := flag.String("port", "8088", "API server port")
	flag.Parse() // Parse command-line arguments

	// Initialize the database connection
	db.InitDB()

	// Create a new Gin router
	router := gin.Default()

	// ✅ Enable CORS (Allow frontend to access API)
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	// Define API routes
	router.GET("/stocks", api.GetStocks)
	router.GET("/stocks/:ticker", api.GetStockByTicker)
	router.GET("/recommendation", api.GetRecommendation)

	// Start the server on the specified port
	fmt.Println("🚀 Server running on http://localhost:" + *port)
	err := router.Run(":" + *port)
	if err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
