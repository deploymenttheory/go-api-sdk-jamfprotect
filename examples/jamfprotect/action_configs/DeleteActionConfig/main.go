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

	actionConfigID := "action-config-id-here" // Replace with actual action config ID

	_, err = client.ActionConfig.DeleteActionConfig(ctx, actionConfigID)
	if err != nil {
		log.Fatalf("Failed to delete action config: %v", err)
	}

	fmt.Printf("Successfully deleted action config: %s\n", actionConfigID)

	os.Exit(0)
}
