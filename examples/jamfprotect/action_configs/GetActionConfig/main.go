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

	ac, _, err := client.ActionConfig.GetActionConfig(ctx, actionConfigID)
	if err != nil {
		log.Fatalf("Failed to get action config: %v", err)
	}

	fmt.Printf("Action Config Details:\n")
	fmt.Printf("  ID: %s\n", ac.ID)
	fmt.Printf("  Name: %s\n", ac.Name)
	fmt.Printf("  Description: %s\n", ac.Description)
	fmt.Printf("  Created: %s\n", ac.Created)
	fmt.Printf("  Updated: %s\n", ac.Updated)

	os.Exit(0)
}
