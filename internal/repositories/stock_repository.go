package repositories

import (
	"context"
	"log"

	"github.com/cristianrubioa/stockwise/db"
	"github.com/cristianrubioa/stockwise/internal/models"
)

func SaveStock(stock models.Stock) error {
	query := `UPSERT INTO stocks (ticker, company, brokerage, action, rating_from, rating_to, target_from, target_to, time)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := db.Conn.Exec(context.Background(), query,
		stock.Ticker, stock.Company, stock.Brokerage, stock.Action,
		stock.RatingFrom, stock.RatingTo, stock.TargetFrom, stock.TargetTo, stock.Time,
	)

	if err != nil {
		log.Println("Error saving stock:", err)
		return err
	}

	return nil
}

