package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/analytic"
)

func main() {
	// Create client from environment variables
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Create a new analytic for detecting suspicious process execution
	request := &analytic.CreateAnalyticRequest{
		Name:        "Suspicious Process Execution",
		InputType:   "es_event",
		Description: "Detects execution of suspicious processes",
		Filter:      "event.type = 'exec' AND process.name = 'suspicious_binary'",
		Level:       4,
		Severity:    "HIGH",
		Tags:        []string{"malware", "execution", "threat"},
		Categories:  []string{"malware", "execution"},
		AnalyticActions: []analytic.AnalyticActionInput{
			{
				Name:       "alert",
				Parameters: []string{"security_team"},
			},
			{
				Name:       "terminate",
				Parameters: []string{"process"},
			},
		},
		Context: []analytic.AnalyticContextInput{
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

	created, _, err := client.Analytic.CreateAnalytic(ctx, request)
	if err != nil {
		log.Fatalf("Failed to create analytic: %v", err)
	}

	fmt.Printf("Successfully created analytic:\n")
	fmt.Printf("  UUID: %s\n", created.UUID)
	fmt.Printf("  Name: %s\n", created.Name)
	fmt.Printf("  Label: %s\n", created.Label)
	fmt.Printf("  Input Type: %s\n", created.InputType)
	fmt.Printf("  Description: %s\n", created.Description)
	fmt.Printf("  Filter: %s\n", created.Filter)
	fmt.Printf("  Level: %d\n", created.Level)
	fmt.Printf("  Severity: %s\n", created.Severity)
	fmt.Printf("  Tags: %v\n", created.Tags)
	fmt.Printf("  Categories: %v\n", created.Categories)
	fmt.Printf("  Created: %s\n", created.Created)

	if len(created.AnalyticActions) > 0 {
		fmt.Printf("  Actions:\n")
		for _, action := range created.AnalyticActions {
			fmt.Printf("    - %s (params: %v)\n", action.Name, action.Parameters)
		}
	}

	os.Exit(0)
}
