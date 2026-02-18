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

	filterUUID := "aaaaaaaa-bbbb-4ccc-8ddd-eeeeeeeeeeee" // Replace with actual filter UUID

	filter, _, err := client.UnifiedLoggingFilter.GetUnifiedLoggingFilter(ctx, filterUUID)
	if err != nil {
		log.Fatalf("Failed to get unified logging filter: %v", err)
	}

	fmt.Printf("Unified Logging Filter Details:\n")
	fmt.Printf("  UUID: %s\n", filter.UUID)
	fmt.Printf("  Name: %s\n", filter.Name)
	fmt.Printf("  Description: %s\n", filter.Description)
	fmt.Printf("  Filter: %s\n", filter.Filter)
	fmt.Printf("  Enabled: %t\n", filter.Enabled)
	fmt.Printf("  Created: %s\n", filter.Created)
	fmt.Printf("  Updated: %s\n", filter.Updated)

	os.Exit(0)
}
