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

	// List all analytics
	analytics, _, err := client.Analytic.ListAnalytics(ctx)
	if err != nil {
		log.Fatalf("Failed to list analytics: %v", err)
	}

	fmt.Printf("Found %d analytic(s):\n\n", len(analytics))

	for i, analytic := range analytics {
		fmt.Printf("%d. %s\n", i+1, analytic.Name)
		fmt.Printf("   UUID: %s\n", analytic.UUID)
		fmt.Printf("   Label: %s\n", analytic.Label)
		fmt.Printf("   Input Type: %s\n", analytic.InputType)
		fmt.Printf("   Description: %s\n", analytic.Description)
		fmt.Printf("   Filter: %s\n", analytic.Filter)
		fmt.Printf("   Level: %d\n", analytic.Level)
		fmt.Printf("   Severity: %s\n", analytic.Severity)
		fmt.Printf("   Jamf Managed: %t\n", analytic.Jamf)
		
		if len(analytic.Tags) > 0 {
			fmt.Printf("   Tags: %v\n", analytic.Tags)
		}
		
		if len(analytic.Categories) > 0 {
			fmt.Printf("   Categories: %v\n", analytic.Categories)
		}
		
		if len(analytic.AnalyticActions) > 0 {
			fmt.Printf("   Actions: %d\n", len(analytic.AnalyticActions))
		}
		
		fmt.Println()
	}

	os.Exit(0)
}
