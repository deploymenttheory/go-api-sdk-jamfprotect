package preventlists

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/interfaces"
)

// Service provides operations for Jamf Protect Prevent Lists
type Service struct {
	client interfaces.GraphQLClient
}

// NewService creates a new Prevent Lists service
func NewService(client interfaces.GraphQLClient) *Service {
	return &Service{client: client}
}

// CreatePreventList creates a new prevent list
func (s *Service) CreatePreventList(ctx context.Context, req *CreatePreventListRequest) (*PreventList, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.Type == "" {
		return nil, nil, fmt.Errorf("%w: type is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := buildPreventListVariables(req)
	var result struct {
		CreatePreventList *PreventList `json:"createPreventList"`
	}

	resp, err := s.client.DoGraphQL(ctx, client.EndpointApp, createPreventListMutation, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create prevent list: %w", err)
	}

	return result.CreatePreventList, resp, nil
}

// GetPreventList retrieves a prevent list by ID
func (s *Service) GetPreventList(ctx context.Context, id string) (*PreventList, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := map[string]any{"id": id}
	var result struct {
		GetPreventList *PreventList `json:"getPreventList"`
	}

	resp, err := s.client.DoGraphQL(ctx, client.EndpointApp, getPreventListQuery, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get prevent list: %w", err)
	}

	return result.GetPreventList, resp, nil
}

// UpdatePreventList updates an existing prevent list
func (s *Service) UpdatePreventList(ctx context.Context, id string, req *UpdatePreventListRequest) (*PreventList, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.Type == "" {
		return nil, nil, fmt.Errorf("%w: type is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := buildPreventListVariables(req)
	vars["id"] = id
	var result struct {
		UpdatePreventList *PreventList `json:"updatePreventList"`
	}

	resp, err := s.client.DoGraphQL(ctx, client.EndpointApp, updatePreventListMutation, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update prevent list: %w", err)
	}

	return result.UpdatePreventList, resp, nil
}

// DeletePreventList deletes a prevent list by ID
func (s *Service) DeletePreventList(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := map[string]any{"id": id}

	resp, err := s.client.DoGraphQL(ctx, client.EndpointApp, deletePreventListMutation, vars, nil, headers)
	if err != nil {
		return resp, fmt.Errorf("failed to delete prevent list: %w", err)
	}

	return resp, nil
}

// ListPreventLists retrieves all prevent lists with automatic pagination
func (s *Service) ListPreventLists(ctx context.Context) ([]PreventList, *interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	allItems := make([]PreventList, 0)
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
			ListPreventLists *ListPreventListsResponse `json:"listPreventLists"`
		}

		resp, err := s.client.DoGraphQL(ctx, client.EndpointApp, listPreventListsQuery, vars, &result, headers)
		lastResp = resp
		if err != nil {
			return nil, lastResp, fmt.Errorf("failed to list prevent lists: %w", err)
		}

		if result.ListPreventLists != nil {
			allItems = append(allItems, result.ListPreventLists.Items...)
			if result.ListPreventLists.PageInfo.Next == nil {
				break
			}
			nextToken = result.ListPreventLists.PageInfo.Next
		} else {
			break
		}
	}

	return allItems, lastResp, nil
}

// buildPreventListVariables builds the GraphQL variables map from a request struct
func buildPreventListVariables(req any) map[string]any {
	var (
		name        string
		description string
		typ         string
		tags        []string
		list        []string
	)

	switch r := req.(type) {
	case *CreatePreventListRequest:
		name = r.Name
		description = r.Description
		typ = r.Type
		tags = r.Tags
		list = r.List
	case *UpdatePreventListRequest:
		name = r.Name
		description = r.Description
		typ = r.Type
		tags = r.Tags
		list = r.List
	}

	vars := map[string]any{
		"name":        name,
		"description": description,
		"type":        typ,
		"tags":        tags,
		"list":        list,
	}

	return vars
}
