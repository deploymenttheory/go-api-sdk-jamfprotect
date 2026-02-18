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

	sets, _, err := client.USBControlSet.ListUSBControlSets(ctx)
	if err != nil {
		log.Fatalf("Failed to list USB control sets: %v", err)
	}

	fmt.Printf("Found %d USB control set(s):\n\n", len(sets))

	for i, set := range sets {
		fmt.Printf("%d. %s\n", i+1, set.Name)
		fmt.Printf("   ID: %s\n", set.ID)
		fmt.Printf("   Description: %s\n", set.Description)
		fmt.Printf("   DefaultMountAction: %s\n", set.DefaultMountAction)
		fmt.Printf("   Created: %s\n", set.Created)
		fmt.Println()
	}

	os.Exit(0)
}
