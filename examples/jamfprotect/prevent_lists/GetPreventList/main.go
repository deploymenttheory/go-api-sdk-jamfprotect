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

	preventListID := "prevent-list-id-here" // Replace with actual prevent list ID

	pl, _, err := client.PreventList.GetPreventList(ctx, preventListID)
	if err != nil {
		log.Fatalf("Failed to get prevent list: %v", err)
	}

	fmt.Printf("Prevent List Details:\n")
	fmt.Printf("  ID: %s\n", pl.ID)
	fmt.Printf("  Name: %s\n", pl.Name)
	fmt.Printf("  Description: %s\n", pl.Description)
	fmt.Printf("  Type: %s\n", pl.Type)
	fmt.Printf("  Count: %d\n", pl.Count)
	fmt.Printf("  Created: %s\n", pl.Created)

	os.Exit(0)
}
