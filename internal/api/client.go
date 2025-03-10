package api

import (
	"fmt"
	"log"

	"github.com/go-resty/resty/v2"
	"github.com/cristianrubioa/stockwise/internal/models"
)

const API_URL = "https://8j5baasof2.execute-api.us-west-2.amazonaws.com/production/swechallenge/list"
const API_KEY = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdHRlbXB0cyI6NDcsImVtYWlsIjoiY3Jpc3RpYW5ydWJpb2FAZ21haWwuY29tIiwiZXhwIjoxNzQxMDYwNDYyLCJpZCI6IjAiLCJwYXNzd29yZCI6IicgT1IgJzEnPScxIn0.wvNrQzjc6SJ5rxQL8QdSvdDJXEQvSyPaQkaCcrUx0Cw"

func FetchStockData(nextPage string) (*models.ApiResponse, error) {
	client := resty.New()
	request := client.R().
		SetHeader("Authorization", "Bearer "+API_KEY).
		SetHeader("Content-Type", "application/json")

	if nextPage != "" {
		request.SetQueryParam("next_page", nextPage)
	}

	var response models.ApiResponse
	resp, err := request.SetResult(&response).Get(API_URL)
	if err != nil {
		log.Println("Error making HTTP request:", err)
		return nil, err
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("API error: %d - %s", resp.StatusCode(), resp.String())
	}

	return &response, nil
}
