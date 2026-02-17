package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
)

func main() {
	// Create client from environment variables
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Delete a plan by ID
	planID := "plan-id-here" // Replace with actual plan ID

	err = client.Plans.DeletePlan(ctx, planID)
	if err != nil {
		log.Fatalf("Failed to delete plan: %v", err)
	}

	fmt.Printf("Successfully deleted plan with ID: %s\n", planID)

	os.Exit(0)
}
