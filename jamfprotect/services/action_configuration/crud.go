package actionconfiguration

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/interfaces"
)

// Service provides operations for Jamf Protect Action Configurations.
type Service struct {
	client interfaces.GraphQLClient
}

// NewService creates a new Action Configurations service.
func NewService(client interfaces.GraphQLClient) *Service {
	return &Service{client: client}
}

// CreateActionConfig creates a new action configuration.
func (s *Service) CreateActionConfig(ctx context.Context, req *CreateActionConfigRequest) (*ActionConfig, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.AlertConfig == nil {
		return nil, nil, fmt.Errorf("%w: alertConfig is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := buildActionConfigVariables(req)
	var result struct {
		CreateActionConfigs *ActionConfig `json:"createActionConfigs"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, createActionConfigMutation, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create action config: %w", err)
	}

	return result.CreateActionConfigs, resp, nil
}

// GetActionConfig retrieves an action configuration by ID.
func (s *Service) GetActionConfig(ctx context.Context, id string) (*ActionConfig, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := map[string]any{"id": id}
	var result struct {
		GetActionConfigs *ActionConfig `json:"getActionConfigs"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, getActionConfigQuery, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get action config: %w", err)
	}

	return result.GetActionConfigs, resp, nil
}

// UpdateActionConfig updates an existing action configuration.
func (s *Service) UpdateActionConfig(ctx context.Context, id string, req *UpdateActionConfigRequest) (*ActionConfig, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.AlertConfig == nil {
		return nil, nil, fmt.Errorf("%w: alertConfig is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := buildActionConfigVariables(req)
	vars["id"] = id
	var result struct {
		UpdateActionConfigs *ActionConfig `json:"updateActionConfigs"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, updateActionConfigMutation, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update action config: %w", err)
	}

	return result.UpdateActionConfigs, resp, nil
}

// DeleteActionConfig deletes an action configuration by ID.
func (s *Service) DeleteActionConfig(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := map[string]any{"id": id}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, deleteActionConfigMutation, vars, nil, headers)
	if err != nil {
		return resp, fmt.Errorf("failed to delete action config: %w", err)
	}

	return resp, nil
}

// ListActionConfigs retrieves all action configurations with automatic pagination.
func (s *Service) ListActionConfigs(ctx context.Context) ([]ActionConfigListItem, *interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	allItems := make([]ActionConfigListItem, 0)
	var nextToken *string
	var lastResp *interfaces.Response

	for {
		vars := map[string]any{
			"direction": "ASC",
			"field":     "NAME",
		}
		if nextToken != nil {
			vars["nextToken"] = *nextToken
		}

		var result struct {
			ListActionConfigs *ListActionConfigsResponse `json:"listActionConfigs"`
		}

		resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, listActionConfigsQuery, vars, &result, headers)
		lastResp = resp
		if err != nil {
			return nil, lastResp, fmt.Errorf("failed to list action configs: %w", err)
		}

		if result.ListActionConfigs != nil {
			allItems = append(allItems, result.ListActionConfigs.Items...)
			if result.ListActionConfigs.PageInfo.Next == nil {
				break
			}
			nextToken = result.ListActionConfigs.PageInfo.Next
		} else {
			break
		}
	}

	return allItems, lastResp, nil
}

// buildActionConfigVariables builds the GraphQL variables map from a request struct.
func buildActionConfigVariables(req any) map[string]any {
	var (
		name        string
		description string
		alertConfig map[string]any
		clients     []map[string]any
	)

	switch r := req.(type) {
	case *CreateActionConfigRequest:
		name = r.Name
		description = r.Description
		alertConfig = r.AlertConfig
		clients = r.Clients
	case *UpdateActionConfigRequest:
		name = r.Name
		description = r.Description
		alertConfig = r.AlertConfig
		clients = r.Clients
	}

	vars := map[string]any{
		"name":        name,
		"description": description,
		"alertConfig": alertConfig,
		"clients":     clients,
	}

	return vars
}
