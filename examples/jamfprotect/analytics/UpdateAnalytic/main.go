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

	// Update an existing analytic
	analyticUUID := "analytic-uuid-here" // Replace with actual analytic UUID
	severity := "CRITICAL"
	
	request := &analytics.UpdateAnalyticRequest{
		Name:        "Updated Suspicious Process Execution",
		InputType:   "es_event",
		Description: "Updated detection for suspicious processes with enhanced filtering",
		Filter:      "event.type = 'exec' AND (process.name = 'suspicious_binary' OR process.path CONTAINS '/tmp/')",
		Level:       5,
		Severity:    &severity,
		Tags:        []string{"malware", "execution", "threat", "updated"},
		Categories:  []string{"malware", "execution", "persistence"},
		AnalyticActions: []analytics.AnalyticActionInput{
			{
				Name:       "alert",
				Parameters: []string{"security_team", "soc"},
			},
			{
				Name:       "terminate",
				Parameters: []string{"process"},
			},
			{
				Name:       "isolate",
				Parameters: []string{"device"},
			},
		},
		Context: []analytics.AnalyticContextInput{
			{
				Name:  "process_info",
				Type:  "process",
				Exprs: []string{"process.name", "process.path", "process.pid", "process.ppid"},
			},
			{
				Name:  "user_info",
				Type:  "user",
				Exprs: []string{"user.name", "user.uid"},
			},
			{
				Name:  "file_info",
				Type:  "file",
				Exprs: []string{"file.path", "file.sha256"},
			},
		},
		SnapshotFiles: []string{"/var/log/system.log", "/var/log/security.log"},
	}

	analytic, err := client.Analytics.UpdateAnalytic(ctx, analyticUUID, request)
	if err != nil {
		log.Fatalf("Failed to update analytic: %v", err)
	}

	fmt.Printf("Successfully updated analytic:\n")
	fmt.Printf("  UUID: %s\n", analytic.UUID)
	fmt.Printf("  Name: %s\n", analytic.Name)
	fmt.Printf("  Description: %s\n", analytic.Description)
	fmt.Printf("  Updated: %s\n", analytic.Updated)
	fmt.Printf("  Level: %d\n", analytic.Level)
	fmt.Printf("  Severity: %s\n", analytic.Severity)
	fmt.Printf("  Tags: %v\n", analytic.Tags)

	os.Exit(0)
}
