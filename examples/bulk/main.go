package main

import (
	"context"
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

	// Bulk request: fetch MVRV for multiple assets in a single request.
	// Always specify assets explicitly to control credit consumption.
	// Each non-BTC asset costs 2 credits; BTC costs 1 credit.
	since := time.Now().AddDate(0, 0, -7).Unix()
	resp, err := client.Metrics.GetBulk(ctx, "/market/mvrv", &glassnode.BulkQuery{
		Assets:     []string{"BTC", "ETH"},
		Since:      &since,
		Resolution: glassnode.Resolution24h,
	})
	if err != nil {
		log.Fatalf("Failed to fetch bulk MVRV: %v", err)
	}
	fmt.Printf("Bulk data points: %d\n", len(resp.Data))
	for _, dp := range resp.Data {
		fmt.Printf("  t=%d bulk=%s\n", dp.T, string(dp.Bulk))
	}
}
