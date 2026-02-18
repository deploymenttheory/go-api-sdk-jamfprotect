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

	_, err = client.UnifiedLoggingFilter.DeleteUnifiedLoggingFilter(ctx, filterUUID)
	if err != nil {
		log.Fatalf("Failed to delete unified logging filter: %v", err)
	}

	fmt.Printf("Successfully deleted unified logging filter: %s\n", filterUUID)

	os.Exit(0)
}
