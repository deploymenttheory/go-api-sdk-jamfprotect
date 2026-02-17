package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/plans"
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
	
	request := &plans.UpdatePlanRequest{
		Name:          "Updated Production Security Plan",
		Description:   "Updated security configuration for production systems",
		LogLevel:      &logLevel,
		ActionConfigs: "action-config-id-here", // Replace with actual action config ID
		ExceptionSets: []string{},
		AutoUpdate:    true,
		CommsConfig: plans.CommsConfigInput{
			FQDN:     "protect.example.com",
			Protocol: "HTTPS",
		},
		InfoSync: plans.InfoSyncInput{
			Attrs:                []string{"hostname", "osVersion", "ipAddress"},
			InsightsSyncInterval: 1800, // Updated to 30 minutes
		},
		SignaturesFeedConfig: plans.SignaturesFeedConfigInput{
			Mode: "AUTO",
		},
	}

	plan, err := client.Plans.UpdatePlan(ctx, planID, request)
	if err != nil {
		log.Fatalf("Failed to update plan: %v", err)
	}

	fmt.Printf("Successfully updated plan:\n")
	fmt.Printf("  ID: %s\n", plan.ID)
	fmt.Printf("  Name: %s\n", plan.Name)
	fmt.Printf("  Description: %s\n", plan.Description)
	fmt.Printf("  Updated: %s\n", plan.Updated)
	fmt.Printf("  Log Level: %s\n", plan.LogLevel)
	fmt.Printf("  Auto Update: %t\n", plan.AutoUpdate)

	os.Exit(0)
}
