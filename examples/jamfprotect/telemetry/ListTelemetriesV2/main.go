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

	items, _, err := client.TelemetryV2.ListTelemetriesV2(ctx)
	if err != nil {
		log.Fatalf("Failed to list telemetries v2: %v", err)
	}

	fmt.Printf("Found %d telemetry v2 configuration(s):\n\n", len(items))

	for i, t := range items {
		fmt.Printf("%d. %s\n", i+1, t.Name)
		fmt.Printf("   ID: %s\n", t.ID)
		fmt.Printf("   Description: %s\n", t.Description)
		fmt.Printf("   Created: %s\n", t.Created)
		fmt.Println()
	}

	os.Exit(0)
}
