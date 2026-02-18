package main

import (
	"context"
	"fmt"
	"log"
	"os"

	jamfprotect "github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect"
)

func main() {
	// Create client from environment variables
	client, err := jamfprotect.NewClientFromEnv()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	ctx := context.Background()

	// Get an analytic by UUID
	analyticUUID := "analytic-uuid-here" // Replace with actual analytic UUID
	
	analytic, _, err := client.Analytic.GetAnalytic(ctx, analyticUUID)
	if err != nil {
		log.Fatalf("Failed to get analytic: %v", err)
	}

	fmt.Printf("Analytic Details:\n")
	fmt.Printf("  UUID: %s\n", analytic.UUID)
	fmt.Printf("  Name: %s\n", analytic.Name)
	fmt.Printf("  Label: %s\n", analytic.Label)
	fmt.Printf("  Input Type: %s\n", analytic.InputType)
	fmt.Printf("  Description: %s\n", analytic.Description)
	fmt.Printf("  Long Description: %s\n", analytic.LongDescription)
	fmt.Printf("  Filter: %s\n", analytic.Filter)
	fmt.Printf("  Level: %d\n", analytic.Level)
	fmt.Printf("  Severity: %s\n", analytic.Severity)
	fmt.Printf("  Tenant Severity: %s\n", analytic.TenantSeverity)
	fmt.Printf("  Jamf Managed: %t\n", analytic.Jamf)
	fmt.Printf("  Created: %s\n", analytic.Created)
	fmt.Printf("  Updated: %s\n", analytic.Updated)

	if len(analytic.Tags) > 0 {
		fmt.Printf("  Tags: %v\n", analytic.Tags)
	}

	if len(analytic.Categories) > 0 {
		fmt.Printf("  Categories: %v\n", analytic.Categories)
	}

	if len(analytic.AnalyticActions) > 0 {
		fmt.Printf("  Analytic Actions:\n")
		for _, action := range analytic.AnalyticActions {
			fmt.Printf("    - %s (parameters: %v)\n", action.Name, action.Parameters)
		}
	}

	if len(analytic.TenantActions) > 0 {
		fmt.Printf("  Tenant Actions:\n")
		for _, action := range analytic.TenantActions {
			fmt.Printf("    - %s (parameters: %v)\n", action.Name, action.Parameters)
		}
	}

	if len(analytic.Context) > 0 {
		fmt.Printf("  Context:\n")
		for _, ctx := range analytic.Context {
			fmt.Printf("    - %s (type: %s)\n", ctx.Name, ctx.Type)
			if len(ctx.Exprs) > 0 {
				fmt.Printf("      Expressions: %v\n", ctx.Exprs)
			}
		}
	}

	if len(analytic.SnapshotFiles) > 0 {
		fmt.Printf("  Snapshot Files: %v\n", analytic.SnapshotFiles)
	}

	if analytic.Remediation != "" {
		fmt.Printf("  Remediation: %s\n", analytic.Remediation)
	}

	os.Exit(0)
}
