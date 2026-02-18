package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/custom_prevent_list"
)

func main() {
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	request := &custompreventlist.CreatePreventListRequest{
		Name:        "Blocked file hashes",
		Description: "Hashes of files to block from execution",
		Type:        custompreventlist.PreventTypeFILEHASH,
		Tags:        []string{"security"},
		List:        []string{},
	}

	created, _, err := client.PreventList.CreatePreventList(ctx, request)
	if err != nil {
		log.Fatalf("Failed to create prevent list: %v", err)
	}

	fmt.Printf("Successfully created prevent list:\n")
	fmt.Printf("  ID: %s\n", created.ID)
	fmt.Printf("  Name: %s\n", created.Name)
	fmt.Printf("  Description: %s\n", created.Description)
	fmt.Printf("  Type: %s\n", created.Type)
	fmt.Printf("  Created: %s\n", created.Created)

	os.Exit(0)
}
