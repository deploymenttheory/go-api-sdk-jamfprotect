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

	_, err = client.USBControlSet.DeleteUSBControlSet(ctx, usbControlSetID)
	if err != nil {
		log.Fatalf("Failed to delete USB control set: %v", err)
	}

	fmt.Printf("Successfully deleted USB control set: %s\n", usbControlSetID)

	os.Exit(0)
}
