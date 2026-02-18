package plan

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/interfaces"
)

// Service provides operations for Jamf Protect Plans
type Service struct {
	client interfaces.GraphQLClient
}

// NewService creates a new Plans service
func NewService(client interfaces.GraphQLClient) *Service {
	return &Service{client: client}
}

// CreatePlan creates a new plan
func (s *Service) CreatePlan(ctx context.Context, req *CreatePlanRequest) (*Plan, *interfaces.Response, error) {
	if err := ValidateCreatePlanRequest(req); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := planMutationVariables(req)
	var result struct {
		CreatePlan *Plan `json:"createPlan"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, createPlanMutation, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create plan: %w", err)
	}

	return result.CreatePlan, resp, nil
}

// GetPlan retrieves a plan by ID
func (s *Service) GetPlan(ctx context.Context, id string) (*Plan, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := map[string]any{"id": id}
	var result struct {
		GetPlan *Plan `json:"getPlan"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, getPlanQuery, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get plan: %w", err)
	}

	return result.GetPlan, resp, nil
}

// UpdatePlan updates an existing plan
func (s *Service) UpdatePlan(ctx context.Context, id string, req *UpdatePlanRequest) (*Plan, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}
	if err := ValidateUpdatePlanRequest(req); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := planMutationVariables(req)
	vars["id"] = id
	var result struct {
		UpdatePlan *Plan `json:"updatePlan"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, updatePlanMutation, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update plan: %w", err)
	}

	return result.UpdatePlan, resp, nil
}

// DeletePlan deletes a plan by ID
func (s *Service) DeletePlan(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := map[string]any{"id": id}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, deletePlanMutation, vars, nil, headers)
	if err != nil {
		return resp, fmt.Errorf("failed to delete plan: %w", err)
	}

	return resp, nil
}

// ListPlans retrieves all plans with automatic pagination
func (s *Service) ListPlans(ctx context.Context) ([]Plan, *interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	allItems := make([]Plan, 0)
	var nextToken *string
	var lastResp *interfaces.Response

	for {
		vars := map[string]any{
			"direction": "ASC",
			"field":     "CREATED",
		}
		if nextToken != nil {
			vars["nextToken"] = *nextToken
		}

		var result struct {
			ListPlans *ListPlansResponse `json:"listPlans"`
		}

		resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, listPlansQuery, vars, &result, headers)
		lastResp = resp
		if err != nil {
			return nil, lastResp, fmt.Errorf("failed to list plans: %w", err)
		}

		if result.ListPlans != nil {
			allItems = append(allItems, result.ListPlans.Items...)
			if result.ListPlans.PageInfo.Next == nil {
				break
			}
			nextToken = result.ListPlans.PageInfo.Next
		} else {
			break
		}
	}

	return allItems, lastResp, nil
}

// planMutationVariables returns GraphQL variables for createPlan/updatePlan mutations.
func planMutationVariables(req any) map[string]any {
	var (
		name                 string
		description          string
		logLevel             *string
		actionConfigs        string
		exceptionSets        []string
		telemetry            *string
		telemetryV2          *string
		telemetryV2Null      bool
		analyticSets         []AnalyticSetInput
		usbControlSet        *string
		commsConfig          CommsConfigInput
		infoSync             InfoSyncInput
		autoUpdate           bool
		signaturesFeedConfig SignaturesFeedConfigInput
	)

	switch r := req.(type) {
	case *CreatePlanRequest:
		name = r.Name
		description = r.Description
		logLevel = r.LogLevel
		actionConfigs = r.ActionConfigs
		exceptionSets = r.ExceptionSets
		telemetry = r.Telemetry
		telemetryV2 = r.TelemetryV2
		telemetryV2Null = r.TelemetryV2Null
		analyticSets = r.AnalyticSets
		usbControlSet = r.USBControlSet
		commsConfig = r.CommsConfig
		infoSync = r.InfoSync
		autoUpdate = r.AutoUpdate
		signaturesFeedConfig = r.SignaturesFeedConfig
	case *UpdatePlanRequest:
		name = r.Name
		description = r.Description
		logLevel = r.LogLevel
		actionConfigs = r.ActionConfigs
		exceptionSets = r.ExceptionSets
		telemetry = r.Telemetry
		telemetryV2 = r.TelemetryV2
		telemetryV2Null = r.TelemetryV2Null
		analyticSets = r.AnalyticSets
		usbControlSet = r.USBControlSet
		commsConfig = r.CommsConfig
		infoSync = r.InfoSync
		autoUpdate = r.AutoUpdate
		signaturesFeedConfig = r.SignaturesFeedConfig
	}

	vars := map[string]any{
		"name":          name,
		"description":   description,
		"actionConfigs": actionConfigs,
		"autoUpdate":    autoUpdate,
		"commsConfig": map[string]any{
			"fqdn":     commsConfig.FQDN,
			"protocol": commsConfig.Protocol,
		},
		"infoSync": map[string]any{
			"attrs":                infoSync.Attrs,
			"insightsSyncInterval": infoSync.InsightsSyncInterval,
		},
		"signaturesFeedConfig": map[string]any{
			"mode": signaturesFeedConfig.Mode,
		},
	}

	if logLevel != nil {
		vars["logLevel"] = *logLevel
	}

	if exceptionSets != nil {
		vars["exceptionSets"] = exceptionSets
	}

	if telemetry != nil {
		vars["telemetry"] = *telemetry
	}

	if telemetryV2Null {
		vars["telemetryV2"] = nil
	} else if telemetryV2 != nil {
		vars["telemetryV2"] = *telemetryV2
	}

	if analyticSets != nil {
		analyticSetsVars := make([]map[string]any, 0, len(analyticSets))
		for _, set := range analyticSets {
			analyticSetsVars = append(analyticSetsVars, map[string]any{
				"type": set.Type,
				"uuid": set.UUID,
			})
		}
		vars["analyticSets"] = analyticSetsVars
	}

	if usbControlSet != nil {
		vars["usbControlSet"] = *usbControlSet
	}

	return vars
}
