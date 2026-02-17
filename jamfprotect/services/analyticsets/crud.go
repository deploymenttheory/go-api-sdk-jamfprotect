package analyticsets

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/interfaces"
)

// Service provides operations for Jamf Protect Analytic Sets
type Service struct {
	client interfaces.GraphQLClient
}

// NewService creates a new Analytic Sets service
func NewService(client interfaces.GraphQLClient) *Service {
	return &Service{client: client}
}

// CreateAnalyticSet creates a new analytic set
func (s *Service) CreateAnalyticSet(ctx context.Context, req *CreateAnalyticSetRequest) (*AnalyticSet, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.Analytics == nil {
		return nil, nil, fmt.Errorf("%w: analytics is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := buildAnalyticSetVariables(req, "")
	var result struct {
		CreateAnalyticSet *AnalyticSet `json:"createAnalyticSet"`
	}

	resp, err := s.client.DoGraphQL(ctx, client.EndpointApp, createAnalyticSetMutation, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create analytic set: %w", err)
	}

	return result.CreateAnalyticSet, resp, nil
}

// GetAnalyticSet retrieves an analytic set by UUID
func (s *Service) GetAnalyticSet(ctx context.Context, uuid string) (*AnalyticSet, *interfaces.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("%w: uuid is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := map[string]any{
		"uuid":             uuid,
		"RBAC_Plan":        true,
		"excludeAnalytics": false,
	}
	var result struct {
		GetAnalyticSet *AnalyticSet `json:"getAnalyticSet"`
	}

	resp, err := s.client.DoGraphQL(ctx, client.EndpointApp, getAnalyticSetQuery, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get analytic set: %w", err)
	}

	return result.GetAnalyticSet, resp, nil
}

// UpdateAnalyticSet updates an existing analytic set
func (s *Service) UpdateAnalyticSet(ctx context.Context, uuid string, req *UpdateAnalyticSetRequest) (*AnalyticSet, *interfaces.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("%w: uuid is required", client.ErrInvalidInput)
	}
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.Analytics == nil {
		return nil, nil, fmt.Errorf("%w: analytics is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := buildAnalyticSetVariables(req, uuid)
	var result struct {
		UpdateAnalyticSet *AnalyticSet `json:"updateAnalyticSet"`
	}

	resp, err := s.client.DoGraphQL(ctx, client.EndpointApp, updateAnalyticSetMutation, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update analytic set: %w", err)
	}

	return result.UpdateAnalyticSet, resp, nil
}

// DeleteAnalyticSet deletes an analytic set by UUID
func (s *Service) DeleteAnalyticSet(ctx context.Context, uuid string) (*interfaces.Response, error) {
	if uuid == "" {
		return nil, fmt.Errorf("%w: uuid is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := map[string]any{"uuid": uuid}

	resp, err := s.client.DoGraphQL(ctx, client.EndpointApp, deleteAnalyticSetMutation, vars, nil, headers)
	if err != nil {
		return resp, fmt.Errorf("failed to delete analytic set: %w", err)
	}

	return resp, nil
}

// ListAnalyticSets retrieves all analytic sets with automatic pagination
func (s *Service) ListAnalyticSets(ctx context.Context) ([]AnalyticSet, *interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	allItems := make([]AnalyticSet, 0)
	var nextToken *string
	var lastResp *interfaces.Response

	for {
		vars := map[string]any{
			"RBAC_Plan":        true,
			"excludeAnalytics": false,
		}
		if nextToken != nil {
			vars["nextToken"] = *nextToken
		}

		var result struct {
			ListAnalyticSets *ListAnalyticSetsResponse `json:"listAnalyticSets"`
		}

		resp, err := s.client.DoGraphQL(ctx, client.EndpointApp, listAnalyticSetsQuery, vars, &result, headers)
		lastResp = resp
		if err != nil {
			return nil, lastResp, fmt.Errorf("failed to list analytic sets: %w", err)
		}

		if result.ListAnalyticSets != nil {
			allItems = append(allItems, result.ListAnalyticSets.Items...)
			if result.ListAnalyticSets.PageInfo.Next == nil {
				break
			}
			nextToken = result.ListAnalyticSets.PageInfo.Next
		} else {
			break
		}
	}

	return allItems, lastResp, nil
}

// buildAnalyticSetVariables builds the GraphQL variables map from a request struct
func buildAnalyticSetVariables(req any, uuid string) map[string]any {
	var (
		name        string
		description string
		types       []string
		analytics   []string
	)

	switch r := req.(type) {
	case *CreateAnalyticSetRequest:
		name = r.Name
		description = r.Description
		types = r.Types
		analytics = r.Analytics
	case *UpdateAnalyticSetRequest:
		name = r.Name
		description = r.Description
		types = r.Types
		analytics = r.Analytics
	}

	vars := map[string]any{
		"name":             name,
		"description":      description,
		"types":            types,
		"analytics":        analytics,
		"RBAC_Plan":        true,
		"excludeAnalytics": false,
	}

	if uuid != "" {
		vars["uuid"] = uuid
	}

	return vars
}
