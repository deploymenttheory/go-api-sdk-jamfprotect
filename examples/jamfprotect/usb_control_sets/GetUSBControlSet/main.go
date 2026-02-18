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

	usbControlSetID := "usb-control-set-id-here" // Replace with actual USB control set ID

	usb, _, err := client.USBControlSet.GetUSBControlSet(ctx, usbControlSetID)
	if err != nil {
		log.Fatalf("Failed to get USB control set: %v", err)
	}

	fmt.Printf("USB Control Set Details:\n")
	fmt.Printf("  ID: %s\n", usb.ID)
	fmt.Printf("  Name: %s\n", usb.Name)
	fmt.Printf("  Description: %s\n", usb.Description)
	fmt.Printf("  DefaultMountAction: %s\n", usb.DefaultMountAction)
	fmt.Printf("  Created: %s\n", usb.Created)
	fmt.Printf("  Updated: %s\n", usb.Updated)
	fmt.Printf("  Rules: %d\n", len(usb.Rules))

	os.Exit(0)
}
