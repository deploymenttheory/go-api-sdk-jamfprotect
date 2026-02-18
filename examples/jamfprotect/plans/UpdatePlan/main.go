package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/plan"
)

func main() {
	// Create client from environment variables
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Update an existing plan
	planID := "plan-id-here" // Replace with actual plan ID
	logLevel := "DEBUG"       // Update log level to DEBUG
	
	request := &plan.UpdatePlanRequest{
		Name:          "Updated Production Security Plan",
		Description:   "Updated security configuration for production systems",
		LogLevel:      &logLevel,
		ActionConfigs: "action-config-id-here", // Replace with actual action config ID
		ExceptionSets: []string{},
		AutoUpdate:    true,
		CommsConfig: plan.CommsConfigInput{
			FQDN:     "protect.example.com",
			Protocol: "HTTPS",
		},
		InfoSync: plan.InfoSyncInput{
			Attrs:                []string{"hostname", "osVersion", "ipAddress"},
			InsightsSyncInterval: 1800, // Updated to 30 minutes
		},
		SignaturesFeedConfig: plan.SignaturesFeedConfigInput{
			Mode: "AUTO",
		},
	}

	updated, _, err := client.Plan.UpdatePlan(ctx, planID, request)
	if err != nil {
		log.Fatalf("Failed to update plan: %v", err)
	}

	fmt.Printf("Successfully updated plan:\n")
	fmt.Printf("  ID: %s\n", updated.ID)
	fmt.Printf("  Name: %s\n", updated.Name)
	fmt.Printf("  Description: %s\n", updated.Description)
	fmt.Printf("  Updated: %s\n", updated.Updated)
	fmt.Printf("  Log Level: %s\n", updated.LogLevel)
	fmt.Printf("  Auto Update: %t\n", updated.AutoUpdate)

	os.Exit(0)
}
