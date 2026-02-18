package unifiedloggingfilter

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/interfaces"
)

// Service provides operations for Jamf Protect Unified Logging Filters
type Service struct {
	client interfaces.GraphQLClient
}

// NewService creates a new Unified Logging Filters service
func NewService(client interfaces.GraphQLClient) *Service {
	return &Service{client: client}
}

// CreateUnifiedLoggingFilter creates a new unified logging filter
func (s *Service) CreateUnifiedLoggingFilter(ctx context.Context, req *CreateUnifiedLoggingFilterRequest) (*UnifiedLoggingFilter, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.Filter == "" {
		return nil, nil, fmt.Errorf("%w: filter is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := unifiedLoggingFilterMutationVariables(req)
	var result struct {
		CreateUnifiedLoggingFilter *UnifiedLoggingFilter `json:"createUnifiedLoggingFilter"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointGraphQL, createUnifiedLoggingFilterMutation, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create unified logging filter: %w", err)
	}

	return result.CreateUnifiedLoggingFilter, resp, nil
}

// GetUnifiedLoggingFilter retrieves a unified logging filter by UUID
func (s *Service) GetUnifiedLoggingFilter(ctx context.Context, uuid string) (*UnifiedLoggingFilter, *interfaces.Response, error) {
	if err := ValidateUnifiedLoggingFilterUUID(uuid); err != nil {
		return nil, nil, err
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := map[string]any{"uuid": uuid}
	var result struct {
		GetUnifiedLoggingFilter *UnifiedLoggingFilter `json:"getUnifiedLoggingFilter"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointGraphQL, getUnifiedLoggingFilterQuery, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get unified logging filter: %w", err)
	}

	return result.GetUnifiedLoggingFilter, resp, nil
}

// UpdateUnifiedLoggingFilter updates an existing unified logging filter
func (s *Service) UpdateUnifiedLoggingFilter(ctx context.Context, uuid string, req *UpdateUnifiedLoggingFilterRequest) (*UnifiedLoggingFilter, *interfaces.Response, error) {
	if err := ValidateUnifiedLoggingFilterUUID(uuid); err != nil {
		return nil, nil, err
	}
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.Filter == "" {
		return nil, nil, fmt.Errorf("%w: filter is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := unifiedLoggingFilterMutationVariables(req)
	vars["uuid"] = uuid
	var result struct {
		UpdateUnifiedLoggingFilter *UnifiedLoggingFilter `json:"updateUnifiedLoggingFilter"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointGraphQL, updateUnifiedLoggingFilterMutation, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update unified logging filter: %w", err)
	}

	return result.UpdateUnifiedLoggingFilter, resp, nil
}

// DeleteUnifiedLoggingFilter deletes a unified logging filter by UUID
func (s *Service) DeleteUnifiedLoggingFilter(ctx context.Context, uuid string) (*interfaces.Response, error) {
	if err := ValidateUnifiedLoggingFilterUUID(uuid); err != nil {
		return nil, err
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := map[string]any{"uuid": uuid}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointGraphQL, deleteUnifiedLoggingFilterMutation, vars, nil, headers)
	if err != nil {
		return resp, fmt.Errorf("failed to delete unified logging filter: %w", err)
	}

	return resp, nil
}

// ListUnifiedLoggingFilters retrieves all unified logging filters with automatic pagination
func (s *Service) ListUnifiedLoggingFilters(ctx context.Context) ([]UnifiedLoggingFilter, *interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	allItems := make([]UnifiedLoggingFilter, 0)
	var nextToken *string
	var lastResp *interfaces.Response

	for {
		vars := map[string]any{
			"direction": "ASC",
			"field":     "NAME",
			"filter":    map[string]any{},
		}
		if nextToken != nil {
			vars["nextToken"] = *nextToken
		}

		var result struct {
			ListUnifiedLoggingFilters *ListUnifiedLoggingFiltersResponse `json:"listUnifiedLoggingFilters"`
		}

		resp, err := s.client.GraphQLPost(ctx, client.EndpointGraphQL, listUnifiedLoggingFiltersQuery, vars, &result, headers)
		lastResp = resp
		if err != nil {
			return nil, lastResp, fmt.Errorf("failed to list unified logging filters: %w", err)
		}

		if result.ListUnifiedLoggingFilters != nil {
			allItems = append(allItems, result.ListUnifiedLoggingFilters.Items...)
			if result.ListUnifiedLoggingFilters.PageInfo.Next == nil {
				break
			}
			nextToken = result.ListUnifiedLoggingFilters.PageInfo.Next
		} else {
			break
		}
	}

	return allItems, lastResp, nil
}

// ListUnifiedLoggingFilterNames retrieves only the names of all unified logging filters
func (s *Service) ListUnifiedLoggingFilterNames(ctx context.Context) ([]string, *interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	var result struct {
		ListUnifiedLoggingFilterNames *ListUnifiedLoggingFilterNamesResponse `json:"listUnifiedLoggingFilterNames"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointGraphQL, listUnifiedLoggingFilterNamesQuery, nil, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list unified logging filter names: %w", err)
	}

	names := []string{}
	if result.ListUnifiedLoggingFilterNames != nil {
		for _, item := range result.ListUnifiedLoggingFilterNames.Items {
			names = append(names, item.Name)
		}
	}

	return names, resp, nil
}
