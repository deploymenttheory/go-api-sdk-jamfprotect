package analytic

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
	if err := ValidateCreateAnalyticRequest(req); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := buildAnalyticVariables(req, false)
	var result struct {
		CreateAnalytic *Analytic `json:"createAnalytic"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, createAnalyticMutation, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to create analytic: %w", err)
	}

	return result.CreateAnalytic, resp, nil
}

// GetAnalytic retrieves an analytic by UUID
func (s *Service) GetAnalytic(ctx context.Context, uuid string) (*Analytic, *interfaces.Response, error) {
	if err := ValidateAnalyticID(uuid); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := map[string]any{"uuid": uuid}
	var result struct {
		GetAnalytic *Analytic `json:"getAnalytic"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, getAnalyticQuery, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to get analytic: %w", err)
	}

	return result.GetAnalytic, resp, nil
}

// UpdateAnalytic updates an existing analytic
func (s *Service) UpdateAnalytic(ctx context.Context, uuid string, req *UpdateAnalyticRequest) (*Analytic, *interfaces.Response, error) {
	if err := ValidateAnalyticID(uuid); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}
	if err := ValidateUpdateAnalyticRequest(req); err != nil {
		return nil, nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
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

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, updateAnalyticMutation, vars, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to update analytic: %w", err)
	}

	return result.UpdateAnalytic, resp, nil
}

// DeleteAnalytic deletes an analytic by UUID
func (s *Service) DeleteAnalytic(ctx context.Context, uuid string) (*interfaces.Response, error) {
	if err := ValidateAnalyticID(uuid); err != nil {
		return nil, fmt.Errorf("%w: %v", client.ErrInvalidInput, err)
	}

	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	vars := map[string]any{"uuid": uuid}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, deleteAnalyticMutation, vars, nil, headers)
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

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, listAnalyticsQuery, nil, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list analytics: %w", err)
	}

	if result.ListAnalytics != nil {
		return result.ListAnalytics.Items, resp, nil
	}

	return []Analytic{}, resp, nil
}

// ListAnalyticsLite retrieves a lightweight summary of all analytics
func (s *Service) ListAnalyticsLite(ctx context.Context) ([]AnalyticLite, *interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	var result struct {
		ListAnalytics *ListAnalyticsLiteResponse `json:"listAnalytics"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, listAnalyticsLiteQuery, nil, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list analytics lite: %w", err)
	}

	if result.ListAnalytics != nil {
		return result.ListAnalytics.Items, resp, nil
	}

	return []AnalyticLite{}, resp, nil
}

// ListAnalyticsNames retrieves only the names of all analytics
func (s *Service) ListAnalyticsNames(ctx context.Context) ([]string, *interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	var result struct {
		ListAnalyticsNames *struct {
			Items []struct {
				Name string `json:"name"`
			} `json:"items"`
		} `json:"listAnalyticsNames"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, listAnalyticsNamesQuery, nil, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list analytics names: %w", err)
	}

	names := []string{}
	if result.ListAnalyticsNames != nil {
		for _, item := range result.ListAnalyticsNames.Items {
			names = append(names, item.Name)
		}
	}

	return names, resp, nil
}

// ListAnalyticsCategories retrieves all analytics categories with their counts
func (s *Service) ListAnalyticsCategories(ctx context.Context) ([]AnalyticCategory, *interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	var result struct {
		ListAnalyticsCategories []AnalyticCategory `json:"listAnalyticsCategories"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, listAnalyticsCategoriesQuery, nil, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list analytics categories: %w", err)
	}

	if result.ListAnalyticsCategories != nil {
		return result.ListAnalyticsCategories, resp, nil
	}

	return []AnalyticCategory{}, resp, nil
}

// ListAnalyticsTags retrieves all analytics tags with their counts
func (s *Service) ListAnalyticsTags(ctx context.Context) ([]AnalyticTag, *interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	var result struct {
		ListAnalyticsTags []AnalyticTag `json:"listAnalyticsTags"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, listAnalyticsTagsQuery, nil, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list analytics tags: %w", err)
	}

	if result.ListAnalyticsTags != nil {
		return result.ListAnalyticsTags, resp, nil
	}

	return []AnalyticTag{}, resp, nil
}

// ListAnalyticsFilterOptions retrieves both tags and categories for populating filter UIs
func (s *Service) ListAnalyticsFilterOptions(ctx context.Context) (*AnalyticsFilterOptions, *interfaces.Response, error) {
	headers := map[string]string{
		"Accept":       client.AcceptJSON,
		"Content-Type": client.ContentTypeJSON,
	}

	var result struct {
		ListAnalyticsTags       []AnalyticTag      `json:"listAnalyticsTags"`
		ListAnalyticsCategories []AnalyticCategory `json:"listAnalyticsCategories"`
	}

	resp, err := s.client.GraphQLPost(ctx, client.EndpointApp, listAnalyticsFilterOptionsQuery, nil, &result, headers)
	if err != nil {
		return nil, resp, fmt.Errorf("failed to list analytics filter options: %w", err)
	}

	opts := &AnalyticsFilterOptions{
		Tags:       result.ListAnalyticsTags,
		Categories: result.ListAnalyticsCategories,
	}

	if opts.Tags == nil {
		opts.Tags = []AnalyticTag{}
	}
	if opts.Categories == nil {
		opts.Categories = []AnalyticCategory{}
	}

	return opts, resp, nil
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
