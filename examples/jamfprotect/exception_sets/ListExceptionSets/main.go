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

	items, _, err := client.ExceptionSet.ListExceptionSets(ctx)
	if err != nil {
		log.Fatalf("Failed to list exception sets: %v", err)
	}

	fmt.Printf("Found %d exception set(s):\n\n", len(items))

	for i, item := range items {
		fmt.Printf("%d. %s\n", i+1, item.Name)
		fmt.Printf("   UUID: %s\n", item.UUID)
		fmt.Printf("   Managed: %t\n", item.Managed)
		fmt.Println()
	}

	os.Exit(0)
}
