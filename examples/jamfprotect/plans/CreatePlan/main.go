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

	// Create a new plan
	logLevel := "INFO"
	request := &plans.CreatePlanRequest{
		Name:          "Production Security Plan",
		Description:   "Security configuration for production systems",
		LogLevel:      &logLevel,
		ActionConfigs: "action-config-id-here",  // Replace with actual action config ID
		ExceptionSets: []string{},
		AutoUpdate:    true,
		CommsConfig: plans.CommsConfigInput{
			FQDN:     "protect.example.com",
			Protocol: "HTTPS",
		},
		InfoSync: plans.InfoSyncInput{
			Attrs:                []string{"hostname", "osVersion"},
			InsightsSyncInterval: 3600,
		},
		SignaturesFeedConfig: plans.SignaturesFeedConfigInput{
			Mode: "AUTO",
		},
	}

	plan, err := client.Plans.CreatePlan(ctx, request)
	if err != nil {
		log.Fatalf("Failed to create plan: %v", err)
	}

	fmt.Printf("Successfully created plan:\n")
	fmt.Printf("  ID: %s\n", plan.ID)
	fmt.Printf("  Name: %s\n", plan.Name)
	fmt.Printf("  Description: %s\n", plan.Description)
	fmt.Printf("  Created: %s\n", plan.Created)
	fmt.Printf("  Log Level: %s\n", plan.LogLevel)
	fmt.Printf("  Auto Update: %t\n", plan.AutoUpdate)

	os.Exit(0)
}
