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

	_, err = client.ExceptionSet.DeleteExceptionSet(ctx, exceptionSetUUID)
	if err != nil {
		log.Fatalf("Failed to delete exception set: %v", err)
	}

	fmt.Printf("Successfully deleted exception set: %s\n", exceptionSetUUID)

	os.Exit(0)
}
