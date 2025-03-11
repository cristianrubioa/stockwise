package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cristianrubioa/stockwise/db"
	"github.com/cristianrubioa/stockwise/internal/api"
	"github.com/cristianrubioa/stockwise/internal/repositories"
)

func main() {
	fmt.Println("🚀 Starting database initialization...")

	// Connect to CockroachDB
	db.InitDB()

	// Ensure the `stockwise` database exists
	createDatabaseIfNotExists("stockwise")

	// Check if the `stocks` table exists
	if !tableExists("stocks") {
		createStocksTable()
	} else {
		fmt.Println("✅ The 'stocks' table already exists.")
	}

	// Check if data already exists before inserting
	if dataExists("stocks") {
		fmt.Println("⚠️ Data is already loaded. No duplicates will be inserted.")
		return
	}

	// Fetch and store data from the API
	fmt.Println("📡 Fetching data from the API...")
	response, err := api.FetchStockData("")
	if err != nil {
		log.Fatal("❌ Error fetching data from the API:", err)
	}

	fmt.Println("💾 Saving data to the database...")
	for _, stock := range response.Items {
		err := repositories.SaveStock(stock)
		if err != nil {
			log.Println("❌ Error saving stock:", err)
		} else {
			fmt.Println("✅ Stock saved:", stock.Ticker, "-", stock.Company)
		}
	}

	fmt.Println("🎉 Initialization complete.")
}

// Ensures the database exists
func createDatabaseIfNotExists(databaseName string) {
	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s;", databaseName)
	_, err := db.Conn.Exec(context.Background(), query)
	if err != nil {
		log.Fatal("❌ Error creating database:", err)
	}
	fmt.Println("✅ Database ensured:", databaseName)
}

// Checks if a table exists in CockroachDB
func tableExists(tableName string) bool {
	query := `SELECT to_regclass($1)`
	var exists string
	err := db.Conn.QueryRow(context.Background(), query, tableName).Scan(&exists)
	return err == nil && exists != ""
}

// Creates the `stocks` table
func createStocksTable() {
	query := `CREATE TABLE stocks (
		id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
		ticker STRING NOT NULL,
		company STRING NOT NULL,
		brokerage STRING NOT NULL,
		action STRING NOT NULL,
		rating_from STRING NOT NULL,
		rating_to STRING NOT NULL,
		target_from STRING NOT NULL,
		target_to STRING NOT NULL,
		time TIMESTAMP NOT NULL,
		UNIQUE (ticker, time)
	)`
	_, err := db.Conn.Exec(context.Background(), query)
	if err != nil {
		log.Fatal("❌ Error creating the 'stocks' table:", err)
	}
	fmt.Println("✅ The 'stocks' table was successfully created.")
}

// Checks if data already exists in the `stocks` table
func dataExists(tableName string) bool {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)
	var count int
	err := db.Conn.QueryRow(context.Background(), query).Scan(&count)
	return err == nil && count > 0
}
