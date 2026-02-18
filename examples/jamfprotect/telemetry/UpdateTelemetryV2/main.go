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

	telemetryID := "telemetry-id-here" // Replace with actual telemetry ID

	request := &telemetry.UpdateTelemetryV2Request{
		Name:               "Endpoint Telemetry (Updated)",
		Description:        "Updated description",
		LogFiles:           []string{"/var/log/system.log", "/var/log/syslog"},
		LogFileCollection:  true,
		PerformanceMetrics: false,
		Events:             []string{},
	}

	updated, _, err := client.TelemetryV2.UpdateTelemetryV2(ctx, telemetryID, request)
	if err != nil {
		log.Fatalf("Failed to update telemetry v2: %v", err)
	}

	fmt.Printf("Successfully updated telemetry v2:\n")
	fmt.Printf("  ID: %s\n", updated.ID)
	fmt.Printf("  Name: %s\n", updated.Name)
	fmt.Printf("  Updated: %s\n", updated.Updated)

	os.Exit(0)
}
