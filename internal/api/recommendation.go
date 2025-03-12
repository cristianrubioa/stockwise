package api

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/cristianrubioa/stockwise/db"
	"github.com/gin-gonic/gin"
)

type StockRecommendation struct {
	Ticker     string  `json:"ticker"`
	Company    string  `json:"company"`
	Brokerage  string  `json:"brokerage"`
	Action     string  `json:"action"`
	RatingFrom string  `json:"rating_from"`
	RatingTo   string  `json:"rating_to"`
	TargetFrom float64 `json:"target_from"`
	TargetTo   float64 `json:"target_to"`
	Score      float64 `json:"score"`
}

// GetRecommendation returns the best stock based on the data
func GetRecommendation(c *gin.Context) {
	query := `SELECT ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to FROM stocks`
	rows, err := db.Conn.Query(context.Background(), query)
	if err != nil {
		log.Println("❌ Error fetching stocks:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch stocks"})
		return
	}
	defer rows.Close()

	var bestStock StockRecommendation
	var bestScore float64 = -1

	for rows.Next() {
		var stock StockRecommendation
		var targetFromStr, targetToStr string

		// Scan `target_from` and `target_to` as `string` and then convert them
		err := rows.Scan(&stock.Ticker, &stock.Company, &stock.Brokerage, &stock.Action, &stock.RatingFrom, &stock.RatingTo, &targetFromStr, &targetToStr)
		if err != nil {
			log.Println("❌ Error scanning stock:", err)
			continue
		}

		// Convert `target_from` and `target_to` from string to float64
		stock.TargetFrom, err = strconv.ParseFloat(targetFromStr[1:], 64) // Remove `$`
		if err != nil {
			log.Println("❌ Error converting target_from:", err)
			continue
		}
		stock.TargetTo, err = strconv.ParseFloat(targetToStr[1:], 64)
		if err != nil {
			log.Println("❌ Error converting target_to:", err)
			continue
		}

		// Calculate score based on rating changes and target price
		score := 0.0

		// If the target price increased, it's positive
		if stock.TargetTo > stock.TargetFrom {
			score += (stock.TargetTo - stock.TargetFrom) / stock.TargetFrom * 100 // % growth
		}

		// If the rating improved, it's positive
		if stock.RatingFrom != stock.RatingTo {
			score += 10
		}

		// If the broker upgraded the stock, it's positive
		if stock.Action == "upgraded by" {
			score += 15
		}

		// If this is the best stock so far, store it
		if score > bestScore {
			bestStock = stock
			bestStock.Score = score
			bestScore = score
		}
	}

	// If there is no data, return an error
	if bestScore == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No stocks found"})
		return
	}

	c.JSON(http.StatusOK, bestStock)
}
