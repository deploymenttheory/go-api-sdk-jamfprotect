package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
)

func main() {
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	telemetryID := "telemetry-id-here" // Replace with actual telemetry ID

	t, _, err := client.TelemetryV2.GetTelemetryV2(ctx, telemetryID)
	if err != nil {
		log.Fatalf("Failed to get telemetry v2: %v", err)
	}

	fmt.Printf("Telemetry V2 Details:\n")
	fmt.Printf("  ID: %s\n", t.ID)
	fmt.Printf("  Name: %s\n", t.Name)
	fmt.Printf("  Description: %s\n", t.Description)
	fmt.Printf("  Created: %s\n", t.Created)
	fmt.Printf("  Updated: %s\n", t.Updated)

	os.Exit(0)
}
