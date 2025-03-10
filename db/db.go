package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var Conn *pgx.Conn // Global variable for the database connection

func InitDB() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}

	// Get the database URL from environment variables
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL is not defined in .env")
	}

	// Connect to CockroachDB
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		log.Fatal("❌ Failed to connect to CockroachDB:", err)
	}

	fmt.Println("✅ Successfully connected to CockroachDB")
	Conn = conn
}
