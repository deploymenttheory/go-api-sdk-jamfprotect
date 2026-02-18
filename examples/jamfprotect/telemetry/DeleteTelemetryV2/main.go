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

	_, err = client.TelemetryV2.DeleteTelemetryV2(ctx, telemetryID)
	if err != nil {
		log.Fatalf("Failed to delete telemetry v2: %v", err)
	}

	fmt.Printf("Successfully deleted telemetry v2: %s\n", telemetryID)

	os.Exit(0)
}
