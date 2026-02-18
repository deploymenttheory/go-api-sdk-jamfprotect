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

	_, err = client.AnalyticSet.DeleteAnalyticSet(ctx, analyticSetUUID)
	if err != nil {
		log.Fatalf("Failed to delete analytic set: %v", err)
	}

	fmt.Printf("Successfully deleted analytic set: %s\n", analyticSetUUID)

	os.Exit(0)
}
