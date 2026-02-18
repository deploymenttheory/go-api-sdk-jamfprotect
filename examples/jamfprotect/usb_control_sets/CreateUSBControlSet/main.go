package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/removable_storage_control_set"
)

func main() {
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	request := &removablestoragecontrolset.CreateUSBControlSetRequest{
		Name:                 "Corporate USB Policy",
		Description:          "Default read-only for unknown devices",
		DefaultMountAction:   removablestoragecontrolset.MountActionReadOnly,
		DefaultMessageAction: "",
		Rules:                []removablestoragecontrolset.USBControlRuleInput{},
	}

	created, _, err := client.USBControlSet.CreateUSBControlSet(ctx, request)
	if err != nil {
		log.Fatalf("Failed to create USB control set: %v", err)
	}

	fmt.Printf("Successfully created USB control set:\n")
	fmt.Printf("  ID: %s\n", created.ID)
	fmt.Printf("  Name: %s\n", created.Name)
	fmt.Printf("  Description: %s\n", created.Description)
	fmt.Printf("  DefaultMountAction: %s\n", created.DefaultMountAction)
	fmt.Printf("  Created: %s\n", created.Created)

	os.Exit(0)
}
