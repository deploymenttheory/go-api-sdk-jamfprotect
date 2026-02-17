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

	// List all plans (automatically handles pagination)
	plans, err := client.Plans.ListPlans(ctx)
	if err != nil {
		log.Fatalf("Failed to list plans: %v", err)
	}

	fmt.Printf("Found %d plan(s):\n\n", len(plans))

	for i, plan := range plans {
		fmt.Printf("%d. %s\n", i+1, plan.Name)
		fmt.Printf("   ID: %s\n", plan.ID)
		fmt.Printf("   Description: %s\n", plan.Description)
		fmt.Printf("   Log Level: %s\n", plan.LogLevel)
		fmt.Printf("   Auto Update: %t\n", plan.AutoUpdate)
		fmt.Printf("   Created: %s\n", plan.Created)
		fmt.Printf("   Updated: %s\n", plan.Updated)
		
		if plan.ActionConfigs != nil {
			fmt.Printf("   Action Config: %s\n", plan.ActionConfigs.Name)
		}
		
		if len(plan.ExceptionSets) > 0 {
			fmt.Printf("   Exception Sets: %d\n", len(plan.ExceptionSets))
		}
		
		if len(plan.AnalyticSets) > 0 {
			fmt.Printf("   Analytic Sets: %d\n", len(plan.AnalyticSets))
		}
		
		fmt.Println()
	}

	os.Exit(0)
}
