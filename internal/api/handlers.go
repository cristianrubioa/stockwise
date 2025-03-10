package api

import (
	"context"
	"net/http"

	"github.com/cristianrubioa/stockwise/db"
	"github.com/cristianrubioa/stockwise/internal/models"
	"github.com/gin-gonic/gin"
)

// GetStocks retrieves all stocks from the database
func GetStocks(c *gin.Context) {
	rows, err := db.Conn.Query(context.Background(), "SELECT ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time FROM stocks")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stocks"})
		return
	}
	defer rows.Close()

	var stocks []models.Stock
	for rows.Next() {
		var stock models.Stock
		err := rows.Scan(&stock.Ticker, &stock.Company, &stock.Brokerage, &stock.Action, &stock.RatingFrom, &stock.RatingTo, &stock.TargetFrom, &stock.TargetTo, &stock.Time)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning data"})
			return
		}
		stocks = append(stocks, stock)
	}

	c.JSON(http.StatusOK, stocks)
}

// GetStockByTicker retrieves a single stock by ticker
func GetStockByTicker(c *gin.Context) {
	ticker := c.Param("ticker")

	var stock models.Stock
	err := db.Conn.QueryRow(context.Background(), "SELECT ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time FROM stocks WHERE ticker=$1", ticker).
		Scan(&stock.Ticker, &stock.Company, &stock.Brokerage, &stock.Action, &stock.RatingFrom, &stock.RatingTo, &stock.TargetFrom, &stock.TargetTo, &stock.Time)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
		return
	}

	c.JSON(http.StatusOK, stock)
}
