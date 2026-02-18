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

	analyticSetUUID := "aaaaaaaa-bbbb-4ccc-8ddd-eeeeeeeeeeee" // Replace with actual analytic set UUID

	request := &analyticset.UpdateAnalyticSetRequest{
		Name:        "Production Analytic Set (Updated)",
		Description: "Updated description",
		Analytics:   []string{"analytic-uuid-1"},
	}

	updated, _, err := client.AnalyticSet.UpdateAnalyticSet(ctx, analyticSetUUID, request)
	if err != nil {
		log.Fatalf("Failed to update analytic set: %v", err)
	}

	fmt.Printf("Successfully updated analytic set:\n")
	fmt.Printf("  UUID: %s\n", updated.UUID)
	fmt.Printf("  Name: %s\n", updated.Name)
	fmt.Printf("  Description: %s\n", updated.Description)
	fmt.Printf("  Updated: %s\n", updated.Updated)

	os.Exit(0)
}
