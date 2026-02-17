package analytics

import (
	"context"
	"fmt"

	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/client"
	"github.com/deploymenttheory/go-api-sdk-jamfprotect/jamfprotect/interfaces"
)

// Service provides operations for Jamf Protect Analytics
type Service struct {
	client interfaces.GraphQLClient
}

// NewService creates a new Analytics service
func NewService(client interfaces.GraphQLClient) *Service {
	return &Service{client: client}
}

// CreateAnalytic creates a new analytic
func (s *Service) CreateAnalytic(ctx context.Context, req *CreateAnalyticRequest) (*Analytic, *interfaces.Response, error) {
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}
	if req.InputType == "" {
		return nil, nil, fmt.Errorf("%w: inputType is required", client.ErrInvalidInput)
	}
	if req.Filter == "" {
		return nil, nil, fmt.Errorf("%w: filter is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := buildAnalyticVariables(req, false)
	var result struct {
		CreateAnalytic *Analytic `json:"createAnalytic"`
	}

	resp, err := s.client.DoGraphQL(ctx, client.EndpointApp, createAnalyticMutation, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create analytic: %w", err)
	}

	return result.CreateAnalytic, resp, nil
}

// GetAnalytic retrieves an analytic by UUID
func (s *Service) GetAnalytic(ctx context.Context, uuid string) (*Analytic, *interfaces.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("%w: uuid is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := map[string]any{"uuid": uuid}
	var result struct {
		GetAnalytic *Analytic `json:"getAnalytic"`
	}

	resp, err := s.client.DoGraphQL(ctx, client.EndpointApp, getAnalyticQuery, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get analytic: %w", err)
	}

	return result.GetAnalytic, resp, nil
}

// UpdateAnalytic updates an existing analytic
func (s *Service) UpdateAnalytic(ctx context.Context, uuid string, req *UpdateAnalyticRequest) (*Analytic, *interfaces.Response, error) {
	if uuid == "" {
		return nil, nil, fmt.Errorf("%w: uuid is required", client.ErrInvalidInput)
	}
	if req == nil {
		return nil, nil, fmt.Errorf("%w: request cannot be nil", client.ErrInvalidInput)
	}
	if req.Name == "" {
		return nil, nil, fmt.Errorf("%w: name is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := buildAnalyticVariables(req, true)
	vars["uuid"] = uuid
	var result struct {
		UpdateAnalytic *Analytic `json:"updateAnalytic"`
	}

	resp, err := s.client.DoGraphQL(ctx, client.EndpointApp, updateAnalyticMutation, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update analytic: %w", err)
	}

	return result.UpdateAnalytic, resp, nil
}

// DeleteAnalytic deletes an analytic by UUID
func (s *Service) DeleteAnalytic(ctx context.Context, uuid string) (*interfaces.Response, error) {
	if uuid == "" {
		return nil, fmt.Errorf("%w: uuid is required", client.ErrInvalidInput)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := map[string]any{"uuid": uuid}

	resp, err := s.client.DoGraphQL(ctx, client.EndpointApp, deleteAnalyticMutation, vars, nil, headers)
	if err != nil {
		return resp, fmt.Errorf("failed to delete analytic: %w", err)
	}

	return resp, nil
}

// ListAnalytics retrieves all analytics
func (s *Service) ListAnalytics(ctx context.Context) ([]Analytic, *interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	var result struct {
		ListAnalytics *ListAnalyticsResponse `json:"listAnalytics"`
	}

	resp, err := s.client.DoGraphQL(ctx, client.EndpointApp, listAnalyticsQuery, nil, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list analytics: %w", err)
	}

	if result.ListAnalytics != nil {
		return result.ListAnalytics.Items, resp, nil
	}

	return []Analytic{}, resp, nil
}

// buildAnalyticVariables builds the GraphQL variables map from a request struct
func buildAnalyticVariables(req any, isUpdate bool) map[string]any {
	var (
		name            string
		inputType       string
		description     string
		actions         []string
		analyticActions []AnalyticActionInput
		tags            []string
		categories      []string
		filter          string
		context         []AnalyticContextInput
		level           int
		severity        *string
		snapshotFiles   []string
	)

	switch r := req.(type) {
	case *CreateAnalyticRequest:
		name = r.Name
		inputType = r.InputType
		description = r.Description
		actions = r.Actions
		analyticActions = r.AnalyticActions
		tags = r.Tags
		categories = r.Categories
		filter = r.Filter
		context = r.Context
		level = r.Level
		sev := r.Severity
		severity = &sev
		snapshotFiles = r.SnapshotFiles
	case *UpdateAnalyticRequest:
		name = r.Name
		inputType = r.InputType
		description = r.Description
		actions = r.Actions
		analyticActions = r.AnalyticActions
		tags = r.Tags
		categories = r.Categories
		filter = r.Filter
		context = r.Context
		level = r.Level
		severity = r.Severity
		snapshotFiles = r.SnapshotFiles
	}

	vars := map[string]any{
		"name":          name,
		"inputType":     inputType,
		"description":   description,
		"actions":       actions,
		"tags":          tags,
		"categories":    categories,
		"filter":        filter,
		"level":         level,
		"snapshotFiles": snapshotFiles,
	}

	// Build analytic actions
	analyticActionsVars := make([]map[string]any, 0, len(analyticActions))
	for _, action := range analyticActions {
		analyticActionsVars = append(analyticActionsVars, map[string]any{
			"name":       action.Name,
			"parameters": action.Parameters,
		})
	}
	vars["analyticActions"] = analyticActionsVars

	// Build context
	contextVars := make([]map[string]any, 0, len(context))
	for _, ctx := range context {
		contextVars = append(contextVars, map[string]any{
			"name":  ctx.Name,
			"type":  ctx.Type,
			"exprs": ctx.Exprs,
		})
	}
	vars["context"] = contextVars

	// Severity is required for create, optional for update
	if !isUpdate || severity != nil {
		if severity != nil {
			vars["severity"] = *severity
		}
	}

	return vars
}
