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

	usbControlSetID := "usb-control-set-id-here" // Replace with actual USB control set ID

	request := &removablestoragecontrolset.UpdateUSBControlSetRequest{
		Name:                 "Corporate USB Policy (Updated)",
		Description:          "Updated description",
		DefaultMountAction:   removablestoragecontrolset.MountActionReadOnly,
		DefaultMessageAction: "",
		Rules:                []removablestoragecontrolset.USBControlRuleInput{},
	}

	updated, _, err := client.USBControlSet.UpdateUSBControlSet(ctx, usbControlSetID, request)
	if err != nil {
		log.Fatalf("Failed to update USB control set: %v", err)
	}

	fmt.Printf("Successfully updated USB control set:\n")
	fmt.Printf("  ID: %s\n", updated.ID)
	fmt.Printf("  Name: %s\n", updated.Name)
	fmt.Printf("  Updated: %s\n", updated.Updated)

	os.Exit(0)
}
