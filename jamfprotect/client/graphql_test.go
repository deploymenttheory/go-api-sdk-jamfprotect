package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapGraphQLErrors_NoErrors(t *testing.T) {
	err := MapGraphQLErrors(nil)
	assert.NoError(t, err)

	err = MapGraphQLErrors([]GraphQLError{})
	assert.NoError(t, err)
}

func TestMapGraphQLErrors_SingleError(t *testing.T) {
	errors := []GraphQLError{
		{
			Message: "Field validation failed",
		},
	}

	err := MapGraphQLErrors(errors)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Field validation failed")
}

func TestMapGraphQLErrors_NotFound(t *testing.T) {
	errors := []GraphQLError{
		{
			Message: "Resource not found",
		},
	}

	err := MapGraphQLErrors(errors)
	assert.Error(t, err)
	assert.True(t, IsNotFound(err), "expected IsNotFound to be true")
	assert.True(t, IsGraphQL(err), "expected IsGraphQL to be true")
}

func TestMapGraphQLErrors_WithPath(t *testing.T) {
	errors := []GraphQLError{
		{
			Message: "Invalid field value",
			Path:    []any{"createPlan", "input", "name"},
		},
	}

	err := MapGraphQLErrors(errors)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "path: createPlan.input.name")
}

func TestMapGraphQLErrors_WithLocations(t *testing.T) {
	errors := []GraphQLError{
		{
			Message: "Syntax error",
			Locations: []GraphQLLocation{
				{Line: 10, Column: 25},
			},
		},
	}

	err := MapGraphQLErrors(errors)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "locations: 10:25")
}

func TestMapGraphQLErrors_WithExtensions(t *testing.T) {
	errors := []GraphQLError{
		{
			Message: "Validation error",
			Extensions: map[string]any{
				"code":      "VALIDATION_FAILED",
				"fieldName": "email",
			},
		},
	}

	err := MapGraphQLErrors(errors)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "extensions:")
	assert.Contains(t, err.Error(), "VALIDATION_FAILED")
}

func TestMapGraphQLErrors_MultipleErrors(t *testing.T) {
	errors := []GraphQLError{
		{Message: "Error 1"},
		{Message: "Error 2"},
		{Message: "Error 3"},
	}

	err := MapGraphQLErrors(errors)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Error 1")
	assert.Contains(t, err.Error(), "Error 2")
	assert.Contains(t, err.Error(), "Error 3")
}

func TestMapGraphQLErrors_EmptyMessages(t *testing.T) {
	errors := []GraphQLError{
		{Message: ""},
		{Message: ""},
	}

	err := MapGraphQLErrors(errors)
	assert.Error(t, err)
	assert.True(t, IsGraphQL(err), "expected IsGraphQL to be true")
}

func TestFormatGraphQLPath_Mixed(t *testing.T) {
	tests := []struct {
		name string
		path []any
		want string
	}{
		{
			name: "String path",
			path: []any{"createPlan", "input", "name"},
			want: "createPlan.input.name",
		},
		{
			name: "Numeric index",
			path: []any{"items", float64(0), "name"},
			want: "items.0.name",
		},
		{
			name: "Mixed path",
			path: []any{"data", float64(5), "field"},
			want: "data.5.field",
		},
		{
			name: "Empty path",
			path: []any{},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatGraphQLPath(tt.path)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFormatGraphQLLocations(t *testing.T) {
	tests := []struct {
		name      string
		locations []GraphQLLocation
		want      string
	}{
		{
			name: "Single location",
			locations: []GraphQLLocation{
				{Line: 10, Column: 5},
			},
			want: "10:5",
		},
		{
			name: "Multiple locations",
			locations: []GraphQLLocation{
				{Line: 10, Column: 5},
				{Line: 20, Column: 15},
			},
			want: "10:5, 20:15",
		},
		{
			name:      "Empty locations",
			locations: []GraphQLLocation{},
			want:      "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatGraphQLLocations(tt.locations)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestFormatGraphQLExtensions(t *testing.T) {
	tests := []struct {
		name string
		ext  map[string]any
		want string
	}{
		{
			name: "Simple extensions",
			ext: map[string]any{
				"code": "ERROR",
			},
			want: `{"code":"ERROR"}`,
		},
		{
			name: "Empty extensions",
			ext:  map[string]any{},
			want: "",
		},
		{
			name: "Nil extensions",
			ext:  nil,
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := formatGraphQLExtensions(tt.ext)
			assert.Equal(t, tt.want, got)
		})
	}
}
