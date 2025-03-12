package api

import (
	"fmt"
	"log"
	"os"

	"github.com/cristianrubioa/stockwise/internal/models"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func FetchStockData(nextPage string) (*models.ApiResponse, error) {
	apiURL := os.Getenv("API_URL")
	apiKey := os.Getenv("API_KEY")

	client := resty.New()
	request := client.R().
		SetHeader("Authorization", "Bearer "+apiKey).
		SetHeader("Content-Type", "application/json")

	if nextPage != "" {
		request.SetQueryParam("next_page", nextPage)
	}

	var response models.ApiResponse
	resp, err := request.SetResult(&response).Get(apiURL)
	if err != nil {
		log.Println("Error making HTTP request:", err)
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode(), resp.String())
	}

	return &response, nil
}
