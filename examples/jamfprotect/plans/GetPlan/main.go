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

	// Get a plan by ID
	planID := "plan-id-here" // Replace with actual plan ID
	
	plan, err := client.Plans.GetPlan(ctx, planID)
	if err != nil {
		log.Fatalf("Failed to get plan: %v", err)
	}

	fmt.Printf("Plan Details:\n")
	fmt.Printf("  ID: %s\n", plan.ID)
	fmt.Printf("  Name: %s\n", plan.Name)
	fmt.Printf("  Description: %s\n", plan.Description)
	fmt.Printf("  Created: %s\n", plan.Created)
	fmt.Printf("  Updated: %s\n", plan.Updated)
	fmt.Printf("  Log Level: %s\n", plan.LogLevel)
	fmt.Printf("  Auto Update: %t\n", plan.AutoUpdate)

	if plan.CommsConfig != nil {
		fmt.Printf("  Communications:\n")
		fmt.Printf("    FQDN: %s\n", plan.CommsConfig.FQDN)
		fmt.Printf("    Protocol: %s\n", plan.CommsConfig.Protocol)
	}

	if plan.InfoSync != nil {
		fmt.Printf("  Info Sync:\n")
		fmt.Printf("    Attributes: %v\n", plan.InfoSync.Attrs)
		fmt.Printf("    Sync Interval: %d seconds\n", plan.InfoSync.InsightsSyncInterval)
	}

	if plan.ActionConfigs != nil {
		fmt.Printf("  Action Config: %s (%s)\n", plan.ActionConfigs.Name, plan.ActionConfigs.ID)
	}

	if len(plan.ExceptionSets) > 0 {
		fmt.Printf("  Exception Sets:\n")
		for _, set := range plan.ExceptionSets {
			fmt.Printf("    - %s (UUID: %s, Managed: %t)\n", set.Name, set.UUID, set.Managed)
		}
	}

	if len(plan.AnalyticSets) > 0 {
		fmt.Printf("  Analytic Sets:\n")
		for _, set := range plan.AnalyticSets {
			fmt.Printf("    - Type: %s, Name: %s\n", set.Type, set.AnalyticSet.Name)
		}
	}

	os.Exit(0)
}
