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

	analyticSetUUID := "aaaaaaaa-bbbb-4ccc-8ddd-eeeeeeeeeeee" // Replace with actual analytic set UUID

	analyticSet, _, err := client.AnalyticSet.GetAnalyticSet(ctx, analyticSetUUID)
	if err != nil {
		log.Fatalf("Failed to get analytic set: %v", err)
	}

	fmt.Printf("Analytic Set Details:\n")
	fmt.Printf("  UUID: %s\n", analyticSet.UUID)
	fmt.Printf("  Name: %s\n", analyticSet.Name)
	fmt.Printf("  Description: %s\n", analyticSet.Description)
	fmt.Printf("  Managed: %t\n", analyticSet.Managed)
	fmt.Printf("  Created: %s\n", analyticSet.Created)
	fmt.Printf("  Updated: %s\n", analyticSet.Updated)
	if len(analyticSet.Analytics) > 0 {
		fmt.Printf("  Analytics:\n")
		for _, a := range analyticSet.Analytics {
			fmt.Printf("    - %s (%s)\n", a.Name, a.UUID)
		}
	}

	os.Exit(0)
}
