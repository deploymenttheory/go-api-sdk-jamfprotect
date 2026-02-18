package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/action_configuration"
)

func main() {
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	actionConfigID := "action-config-id-here" // Replace with actual action config ID

	request := &actionconfiguration.UpdateActionConfigRequest{
		Name:        "Default Alert Config (Updated)",
		Description: "Updated description",
		AlertConfig: map[string]any{
			"data": map[string]any{},
		},
		Clients: []map[string]any{},
	}

	updated, _, err := client.ActionConfig.UpdateActionConfig(ctx, actionConfigID, request)
	if err != nil {
		log.Fatalf("Failed to update action config: %v", err)
	}

	fmt.Printf("Successfully updated action config:\n")
	fmt.Printf("  ID: %s\n", updated.ID)
	fmt.Printf("  Name: %s\n", updated.Name)
	fmt.Printf("  Updated: %s\n", updated.Updated)

	os.Exit(0)
}
