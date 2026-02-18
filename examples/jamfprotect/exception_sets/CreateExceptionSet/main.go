package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/services/exception_set"
)

func main() {
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	request := &exceptionset.CreateExceptionSetRequest{
		Name:        "Development Exceptions",
		Description: "Exceptions for development team",
		Exceptions:  []exceptionset.ExceptionInput{},
		EsExceptions: []exceptionset.EsExceptionInput{},
	}

	created, _, err := client.ExceptionSet.CreateExceptionSet(ctx, request)
	if err != nil {
		log.Fatalf("Failed to create exception set: %v", err)
	}

	fmt.Printf("Successfully created exception set:\n")
	fmt.Printf("  UUID: %s\n", created.UUID)
	fmt.Printf("  Name: %s\n", created.Name)
	fmt.Printf("  Description: %s\n", created.Description)
	fmt.Printf("  Created: %s\n", created.Created)

	os.Exit(0)
}
