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

	// AlertConfig structure depends on your Jamf Protect alert configuration. This is a minimal example.
	request := &actionconfiguration.CreateActionConfigRequest{
		Name:        "Default Alert Config",
		Description: "Sends alerts to configured destinations",
		AlertConfig: map[string]any{
			"data": map[string]any{},
		},
		Clients: []map[string]any{},
	}

	created, _, err := client.ActionConfig.CreateActionConfig(ctx, request)
	if err != nil {
		log.Fatalf("Failed to create action config: %v", err)
	}

	fmt.Printf("Successfully created action config:\n")
	fmt.Printf("  ID: %s\n", created.ID)
	fmt.Printf("  Name: %s\n", created.Name)
	fmt.Printf("  Description: %s\n", created.Description)
	fmt.Printf("  Created: %s\n", created.Created)

	os.Exit(0)
}
