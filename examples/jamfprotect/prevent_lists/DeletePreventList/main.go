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

	_, err = client.PreventList.DeletePreventList(ctx, preventListID)
	if err != nil {
		log.Fatalf("Failed to delete prevent list: %v", err)
	}

	fmt.Printf("Successfully deleted prevent list: %s\n", preventListID)

	os.Exit(0)
}
