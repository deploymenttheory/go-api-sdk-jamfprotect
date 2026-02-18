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

	request := &unifiedloggingfilter.CreateUnifiedLoggingFilterRequest{
		Name:        "Security subsystem filter",
		Description: "Filters security-related log messages",
		Filter:      "subsystem == \"com.apple.security\"",
		Enabled:     true,
		Tags:        []string{"security"},
	}

	created, _, err := client.UnifiedLoggingFilter.CreateUnifiedLoggingFilter(ctx, request)
	if err != nil {
		log.Fatalf("Failed to create unified logging filter: %v", err)
	}

	fmt.Printf("Successfully created unified logging filter:\n")
	fmt.Printf("  UUID: %s\n", created.UUID)
	fmt.Printf("  Name: %s\n", created.Name)
	fmt.Printf("  Description: %s\n", created.Description)
	fmt.Printf("  Filter: %s\n", created.Filter)
	fmt.Printf("  Enabled: %t\n", created.Enabled)
	fmt.Printf("  Created: %s\n", created.Created)

	os.Exit(0)
}
