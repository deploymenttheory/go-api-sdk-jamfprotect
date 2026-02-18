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

	preventListID := "prevent-list-id-here" // Replace with actual prevent list ID

	request := &custompreventlist.UpdatePreventListRequest{
		Name:        "Blocked file hashes (Updated)",
		Description: "Updated description",
		Type:        custompreventlist.PreventTypeFILEHASH,
		Tags:        []string{"security"},
		List:        []string{},
	}

	updated, _, err := client.PreventList.UpdatePreventList(ctx, preventListID, request)
	if err != nil {
		log.Fatalf("Failed to update prevent list: %v", err)
	}

	fmt.Printf("Successfully updated prevent list:\n")
	fmt.Printf("  ID: %s\n", updated.ID)
	fmt.Printf("  Name: %s\n", updated.Name)

	os.Exit(0)
}
