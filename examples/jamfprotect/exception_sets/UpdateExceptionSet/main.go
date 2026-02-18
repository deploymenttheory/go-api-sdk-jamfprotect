package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/exception_set"
)

func main() {
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	exceptionSetUUID := "aaaaaaaa-bbbb-4ccc-8ddd-eeeeeeeeeeee" // Replace with actual exception set UUID

	request := &exceptionset.UpdateExceptionSetRequest{
		Name:        "Development Exceptions (Updated)",
		Description: "Updated description",
		Exceptions:  []exceptionset.ExceptionInput{},
		EsExceptions: []exceptionset.EsExceptionInput{},
	}

	updated, _, err := client.ExceptionSet.UpdateExceptionSet(ctx, exceptionSetUUID, request)
	if err != nil {
		log.Fatalf("Failed to update exception set: %v", err)
	}

	fmt.Printf("Successfully updated exception set:\n")
	fmt.Printf("  UUID: %s\n", updated.UUID)
	fmt.Printf("  Name: %s\n", updated.Name)
	fmt.Printf("  Updated: %s\n", updated.Updated)

	os.Exit(0)
}
