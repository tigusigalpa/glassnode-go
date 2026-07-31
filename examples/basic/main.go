package main

import (
	"context"
	"fmt"
	"log"
	"time"

	glassnode "github.com/tigusigalpa/glassnode-go"
)

func main() {
	// Create a client with your API key.
	// In production, use NewClientFromEnv to read GLASSNODE_API_KEY from the environment.
	client := glassnode.NewClient("YOUR_API_KEY",
		glassnode.WithTimeout(15*time.Second),
		glassnode.WithRetry(3, time.Second),
	)

	ctx := context.Background()

	// Fetch BTC price
	price, err := client.Market.Price(ctx, &glassnode.MetricQuery{
		Asset:      "BTC",
		Resolution: glassnode.Resolution24h,
	})
	if err != nil {
		log.Fatalf("Failed to fetch price: %v", err)
	}
	for _, p := range price {
		fmt.Printf("BTC price at %d: $%.2f\n", p.T, p.V)
	}

	// Fetch SOPR indicator
	sopr, err := client.Indicators.SOPR(ctx, &glassnode.MetricQuery{
		Asset:      "BTC",
		Resolution: glassnode.Resolution24h,
	})
	if err != nil {
		log.Fatalf("Failed to fetch SOPR: %v", err)
	}
	if len(sopr) > 0 {
		fmt.Printf("Latest SOPR: %.4f\n", sopr[len(sopr)-1].V)
	}

	// List available assets
	assets, err := client.Metadata.Assets(ctx, "")
	if err != nil {
		log.Fatalf("Failed to fetch assets: %v", err)
	}
	fmt.Printf("Total assets: %d\n", len(assets.Data))

	// Check API credit usage
	usage, err := client.User.APIUsage(ctx)
	if err != nil {
		log.Fatalf("Failed to fetch API usage: %v", err)
	}
	fmt.Printf("Credits used this month: %d\n", usage.CreditsUsed)
}
