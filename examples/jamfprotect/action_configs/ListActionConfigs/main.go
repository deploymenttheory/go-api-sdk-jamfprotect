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

	configs, _, err := client.ActionConfig.ListActionConfigs(ctx)
	if err != nil {
		log.Fatalf("Failed to list action configs: %v", err)
	}

	fmt.Printf("Found %d action config(s):\n\n", len(configs))

	for i, ac := range configs {
		fmt.Printf("%d. %s\n", i+1, ac.Name)
		fmt.Printf("   ID: %s\n", ac.ID)
		fmt.Printf("   Description: %s\n", ac.Description)
		fmt.Printf("   Created: %s\n", ac.Created)
		fmt.Printf("   Updated: %s\n", ac.Updated)
		fmt.Println()
	}

	os.Exit(0)
}
