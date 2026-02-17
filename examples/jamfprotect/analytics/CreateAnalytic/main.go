package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytics"
)

func main() {
	// Create client from environment variables
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Create a new analytic for detecting suspicious process execution
	request := &analytics.CreateAnalyticRequest{
		Name:        "Suspicious Process Execution",
		InputType:   "es_event",
		Description: "Detects execution of suspicious processes",
		Filter:      "event.type = 'exec' AND process.name = 'suspicious_binary'",
		Level:       4,
		Severity:    "HIGH",
		Tags:        []string{"malware", "execution", "threat"},
		Categories:  []string{"malware", "execution"},
		AnalyticActions: []analytics.AnalyticActionInput{
			{
				Name:       "alert",
				Parameters: []string{"security_team"},
			},
			{
				Name:       "terminate",
				Parameters: []string{"process"},
			},
		},
		Context: []analytics.AnalyticContextInput{
			{
				Name:  "process_info",
				Type:  "process",
				Exprs: []string{"process.name", "process.path", "process.pid"},
			},
			{
				Name:  "user_info",
				Type:  "user",
				Exprs: []string{"user.name", "user.uid"},
			},
		},
		SnapshotFiles: []string{"/var/log/system.log"},
	}

	analytic, err := client.Analytics.CreateAnalytic(ctx, request)
	if err != nil {
		log.Fatalf("Failed to create analytic: %v", err)
	}

	fmt.Printf("Successfully created analytic:\n")
	fmt.Printf("  UUID: %s\n", analytic.UUID)
	fmt.Printf("  Name: %s\n", analytic.Name)
	fmt.Printf("  Label: %s\n", analytic.Label)
	fmt.Printf("  Input Type: %s\n", analytic.InputType)
	fmt.Printf("  Description: %s\n", analytic.Description)
	fmt.Printf("  Filter: %s\n", analytic.Filter)
	fmt.Printf("  Level: %d\n", analytic.Level)
	fmt.Printf("  Severity: %s\n", analytic.Severity)
	fmt.Printf("  Tags: %v\n", analytic.Tags)
	fmt.Printf("  Categories: %v\n", analytic.Categories)
	fmt.Printf("  Created: %s\n", analytic.Created)
	
	if len(analytic.AnalyticActions) > 0 {
		fmt.Printf("  Actions:\n")
		for _, action := range analytic.AnalyticActions {
			fmt.Printf("    - %s (params: %v)\n", action.Name, action.Parameters)
		}
	}

	os.Exit(0)
}
