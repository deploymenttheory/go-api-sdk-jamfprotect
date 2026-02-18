package removablestoragecontrolset

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/interfaces"
)

// Service provides operations for Jamf Protect USB Control Sets
type Service struct {
	client interfaces.GraphQLClient
}

// NewService creates a new USB Control Sets service
func NewService(client interfaces.GraphQLClient) *Service {
	return &Service{client: client}
}

// CreateUSBControlSet creates a new USB control set
func (s *Service) CreateUSBControlSet(ctx context.Context, req *CreateUSBControlSetRequest) (*USBControlSet, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.DefaultMountAction == "" {
		return nil, nil, fmt.Errorf("%w: defaultMountAction is required", client.ErrInvalidInput)
	}
	if req.Rules == nil {
		return nil, nil, fmt.Errorf("%w: rules is required", client.ErrInvalidInput)
	}
	if err := ValidateCreateUSBControlSetRequest(req); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := usbControlSetMutationVariables(req, "")
	var result struct {
		CreateUSBControlSet *USBControlSet `json:"createUSBControlSet"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, createUSBControlSetMutation, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create USB control set: %w", err)
	}

	return result.CreateUSBControlSet, resp, nil
}

// GetUSBControlSet retrieves a USB control set by ID
func (s *Service) GetUSBControlSet(ctx context.Context, id string) (*USBControlSet, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := map[string]any{"id": id}
	var result struct {
		GetUSBControlSet *USBControlSet `json:"getUSBControlSet"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, getUSBControlSetQuery, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get USB control set: %w", err)
	}

	return result.GetUSBControlSet, resp, nil
}

// UpdateUSBControlSet updates an existing USB control set
func (s *Service) UpdateUSBControlSet(ctx context.Context, id string, req *UpdateUSBControlSetRequest) (*USBControlSet, *interfaces.Response, error) {
	if id == "" {
		return nil, nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.DefaultMountAction == "" {
		return nil, nil, fmt.Errorf("%w: defaultMountAction is required", client.ErrInvalidInput)
	}
	if req.Rules == nil {
		return nil, nil, fmt.Errorf("%w: rules is required", client.ErrInvalidInput)
	}
	if err := ValidateUpdateUSBControlSetRequest(req); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := usbControlSetMutationVariables(req, id)
	var result struct {
		UpdateUSBControlSet *USBControlSet `json:"updateUSBControlSet"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, updateUSBControlSetMutation, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update USB control set: %w", err)
	}

	return result.UpdateUSBControlSet, resp, nil
}

// DeleteUSBControlSet deletes a USB control set by ID
func (s *Service) DeleteUSBControlSet(ctx context.Context, id string) (*interfaces.Response, error) {
	if id == "" {
		return nil, fmt.Errorf("%w: id is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := map[string]any{"id": id}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, deleteUSBControlSetMutation, vars, nil, headers)
	if err != nil {
		return resp, fmt.Errorf("failed to delete USB control set: %w", err)
	}

	return resp, nil
}

// ListUSBControlSets retrieves all USB control sets with automatic pagination
func (s *Service) ListUSBControlSets(ctx context.Context) ([]USBControlSet, *interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	allItems := make([]USBControlSet, 0)
	var nextToken *string
	var lastResp *interfaces.Response

	for {
		vars := map[string]any{
			"direction": "ASC",
			"field":     "created",
		}
		if nextToken != nil {
			vars["nextToken"] = *nextToken
		}

		var result struct {
			ListUSBControlSets *ListUSBControlSetsResponse `json:"listUSBControlSets"`
		}

		resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, listUSBControlSetsQuery, vars, &result, headers)
		lastResp = resp
		if err != nil {
			return nil, lastResp, fmt.Errorf("failed to list USB control sets: %w", err)
		}

		if result.ListUSBControlSets != nil {
			allItems = append(allItems, result.ListUSBControlSets.Items...)
			if result.ListUSBControlSets.PageInfo.Next == nil {
				break
			}
			nextToken = result.ListUSBControlSets.PageInfo.Next
		} else {
			break
		}
	}

	return allItems, lastResp, nil
}

// ListUSBControlSetNames retrieves only the names of all USB control sets
func (s *Service) ListUSBControlSetNames(ctx context.Context) ([]string, *interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	var result struct {
		ListUsbControlNames *ListUSBControlSetNamesResponse `json:"listUsbControlNames"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, listUSBControlSetNamesQuery, nil, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list USB control set names: %w", err)
	}

	names := []string{}
	if result.ListUsbControlNames != nil {
		for _, item := range result.ListUsbControlNames.Items {
			names = append(names, item.Name)
		}
	}

	return names, resp, nil
}
