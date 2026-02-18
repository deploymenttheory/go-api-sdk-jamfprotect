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

	lists, _, err := client.PreventList.ListPreventLists(ctx)
	if err != nil {
		log.Fatalf("Failed to list prevent lists: %v", err)
	}

	fmt.Printf("Found %d prevent list(s):\n\n", len(lists))

	for i, pl := range lists {
		fmt.Printf("%d. %s\n", i+1, pl.Name)
		fmt.Printf("   ID: %s\n", pl.ID)
		fmt.Printf("   Description: %s\n", pl.Description)
		fmt.Printf("   Type: %s\n", pl.Type)
		fmt.Printf("   Count: %d\n", pl.Count)
		fmt.Println()
	}

	os.Exit(0)
}
