package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/telemetry"
)

func main() {
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	request := &telemetry.CreateTelemetryV2Request{
		Name:               "Endpoint Telemetry",
		Description:        "System and security event collection",
		LogFiles:           []string{"/var/log/system.log", "/var/log/syslog"},
		LogFileCollection:  true,
		PerformanceMetrics: false,
		Events:             []string{},
	}

	created, _, err := client.TelemetryV2.CreateTelemetryV2(ctx, request)
	if err != nil {
		log.Fatalf("Failed to create telemetry v2: %v", err)
	}

	fmt.Printf("Successfully created telemetry v2:\n")
	fmt.Printf("  ID: %s\n", created.ID)
	fmt.Printf("  Name: %s\n", created.Name)
	fmt.Printf("  Description: %s\n", created.Description)
	fmt.Printf("  Created: %s\n", created.Created)

	os.Exit(0)
}
