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

	// Create a new plan
	logLevel := "INFO"
	request := &plan.CreatePlanRequest{
		Name:          "Production Security Plan",
		Description:   "Security configuration for production systems",
		LogLevel:      &logLevel,
		ActionConfigs: "action-config-id-here",  // Replace with actual action config ID
		ExceptionSets: []string{},
		AutoUpdate:    true,
		CommsConfig: plan.CommsConfigInput{
			FQDN:     "protect.example.com",
			Protocol: "HTTPS",
		},
		InfoSync: plan.InfoSyncInput{
			Attrs:                []string{"hostname", "osVersion"},
			InsightsSyncInterval: 3600,
		},
		SignaturesFeedConfig: plan.SignaturesFeedConfigInput{
			Mode: "AUTO",
		},
	}

	created, _, err := client.Plan.CreatePlan(ctx, request)
	if err != nil {
		log.Fatalf("Failed to create plan: %v", err)
	}

	fmt.Printf("Successfully created plan:\n")
	fmt.Printf("  ID: %s\n", created.ID)
	fmt.Printf("  Name: %s\n", created.Name)
	fmt.Printf("  Description: %s\n", created.Description)
	fmt.Printf("  Created: %s\n", created.Created)
	fmt.Printf("  Log Level: %s\n", created.LogLevel)
	fmt.Printf("  Auto Update: %t\n", created.AutoUpdate)

	os.Exit(0)
}
