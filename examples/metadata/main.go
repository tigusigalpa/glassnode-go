package main

import (
	"context"
	"encoding/json"
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

	// List all available metric paths
	paths, err := client.Metadata.Metrics(ctx, &glassnode.MetricQuery{})
	if err != nil {
		log.Fatalf("Failed to list metrics: %v", err)
	}
	fmt.Printf("Available metrics: %d\n", len(paths))
	for i, p := range paths {
		if i >= 10 {
			fmt.Println("  ...")
			break
		}
		fmt.Printf("  %s\n", p)
	}

	// Get metadata for a specific metric
	meta, err := client.Metadata.Metric(ctx, "/indicators/sopr", &glassnode.MetricQuery{})
	if err != nil {
		log.Fatalf("Failed to fetch metric metadata: %v", err)
	}
	b, _ := json.MarshalIndent(meta, "", "  ")
	fmt.Printf("Metric metadata:\n%s\n", b)
}
