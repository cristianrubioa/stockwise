package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/cristianrubioa/stockwise/db"
	"github.com/cristianrubioa/stockwise/internal/api"
	"github.com/gin-gonic/gin"
)

func main() {
	// Define a CLI flag for the server port (default: 8088)
	port := flag.String("port", "8088", "API server port")
	flag.Parse() // Parse command-line arguments

	// Initialize the database connection
	db.InitDB()

	// Create a new Gin router
	router := gin.Default()

	// Define API routes
	router.GET("/stocks", api.GetStocks)
	router.GET("/stocks/:ticker", api.GetStockByTicker)

	// Start the server on the specified port
	fmt.Println("🚀 Server running on http://localhost:" + *port)
	err := router.Run(":" + *port)
	if err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}
}
