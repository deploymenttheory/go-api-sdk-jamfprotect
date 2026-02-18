package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytic_set"
)

func main() {
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	request := &analyticset.CreateAnalyticSetRequest{
		Name:        "Production Analytic Set",
		Description: "Analytics for production endpoints",
		Analytics:   []string{"analytic-uuid-1"}, // Replace with actual analytic UUIDs from ListAnalytics
	}

	created, _, err := client.AnalyticSet.CreateAnalyticSet(ctx, request)
	if err != nil {
		log.Fatalf("Failed to create analytic set: %v", err)
	}

	fmt.Printf("Successfully created analytic set:\n")
	fmt.Printf("  UUID: %s\n", created.UUID)
	fmt.Printf("  Name: %s\n", created.Name)
	fmt.Printf("  Description: %s\n", created.Description)
	fmt.Printf("  Created: %s\n", created.Created)

	os.Exit(0)
}
