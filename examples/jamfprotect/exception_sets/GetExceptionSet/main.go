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

	exceptionSetUUID := "aaaaaaaa-bbbb-4ccc-8ddd-eeeeeeeeeeee" // Replace with actual exception set UUID

	exceptionSet, _, err := client.ExceptionSet.GetExceptionSet(ctx, exceptionSetUUID)
	if err != nil {
		log.Fatalf("Failed to get exception set: %v", err)
	}

	fmt.Printf("Exception Set Details:\n")
	fmt.Printf("  UUID: %s\n", exceptionSet.UUID)
	fmt.Printf("  Name: %s\n", exceptionSet.Name)
	fmt.Printf("  Description: %s\n", exceptionSet.Description)
	fmt.Printf("  Managed: %t\n", exceptionSet.Managed)
	fmt.Printf("  Created: %s\n", exceptionSet.Created)
	fmt.Printf("  Updated: %s\n", exceptionSet.Updated)
	fmt.Printf("  Exceptions: %d\n", len(exceptionSet.Exceptions))

	os.Exit(0)
}
