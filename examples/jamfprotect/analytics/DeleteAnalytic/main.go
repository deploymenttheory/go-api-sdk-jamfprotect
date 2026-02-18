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

	// Delete an analytic by UUID
	analyticUUID := "analytic-uuid-here" // Replace with actual analytic UUID

	_, err = client.Analytic.DeleteAnalytic(ctx, analyticUUID)
	if err != nil {
		log.Fatalf("Failed to delete analytic: %v", err)
	}

	fmt.Printf("Successfully deleted analytic with UUID: %s\n", analyticUUID)

	os.Exit(0)
}
