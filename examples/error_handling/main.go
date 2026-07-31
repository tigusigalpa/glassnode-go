package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	glassnode "github.com/tigusigalpa/glassnode-go"
)

func main() {
	client := glassnode.NewClient("YOUR_API_KEY",
		glassnode.WithRetry(3, time.Second),
	)

	ctx := context.Background()

	_, err := client.Indicators.SOPR(ctx, &glassnode.MetricQuery{Asset: "BTC"})
	if err != nil {
		// Check for specific error types
		switch {
		case errors.Is(err, glassnode.ErrUnauthorized):
			fmt.Println("Invalid API key — check your GLASSNODE_API_KEY")

		case errors.Is(err, glassnode.ErrRateLimited):
			var apiErr *glassnode.APIError
			if errors.As(err, &apiErr) {
				fmt.Printf("Rate limited after %d retries. Reset in %s seconds.\n",
					apiErr.Retried, apiErr.RateLimitReset)
			}

		case errors.Is(err, glassnode.ErrBadRequest):
			var apiErr *glassnode.APIError
			if errors.As(err, &apiErr) {
				fmt.Printf("Bad request: %s\n", apiErr.Message)
			}

		case errors.Is(err, glassnode.ErrNotFound):
			fmt.Println("Metric not found — check the metric path")

		default:
			log.Fatalf("Unexpected error: %v", err)
		}
		return
	}

	fmt.Println("Request succeeded")
	_ = time.Second // keep import
}
