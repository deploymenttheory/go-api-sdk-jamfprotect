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

	filters, _, err := client.UnifiedLoggingFilter.ListUnifiedLoggingFilters(ctx)
	if err != nil {
		log.Fatalf("Failed to list unified logging filters: %v", err)
	}

	fmt.Printf("Found %d unified logging filter(s):\n\n", len(filters))

	for i, f := range filters {
		fmt.Printf("%d. %s\n", i+1, f.Name)
		fmt.Printf("   UUID: %s\n", f.UUID)
		fmt.Printf("   Description: %s\n", f.Description)
		fmt.Printf("   Filter: %s\n", f.Filter)
		fmt.Printf("   Enabled: %t\n", f.Enabled)
		fmt.Println()
	}

	os.Exit(0)
}
