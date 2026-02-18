package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/unified_logging_filter"
)

func main() {
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	filterUUID := "aaaaaaaa-bbbb-4ccc-8ddd-eeeeeeeeeeee" // Replace with actual filter UUID

	request := &unifiedloggingfilter.UpdateUnifiedLoggingFilterRequest{
		Name:        "Security subsystem filter (Updated)",
		Description: "Updated description",
		Filter:      "subsystem == \"com.apple.security\"",
		Enabled:     true,
		Tags:        []string{"security"},
	}

	updated, _, err := client.UnifiedLoggingFilter.UpdateUnifiedLoggingFilter(ctx, filterUUID, request)
	if err != nil {
		log.Fatalf("Failed to update unified logging filter: %v", err)
	}

	fmt.Printf("Successfully updated unified logging filter:\n")
	fmt.Printf("  UUID: %s\n", updated.UUID)
	fmt.Printf("  Name: %s\n", updated.Name)
	fmt.Printf("  Updated: %s\n", updated.Updated)

	os.Exit(0)
}
