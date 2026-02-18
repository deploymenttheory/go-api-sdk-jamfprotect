package validate

import (
	"fmt"
	"slices"
)

// OneOf returns nil if value is empty (optional field) or if value is one of allowed.
// Use for allowed-value validation of enum-like fields. Returns an error when value
// is non-empty and not in the allowed set.
func OneOf(fieldName, value string, allowed ...string) error {
	if value == "" {
		return nil
	}
	if slices.Contains(allowed, value) {
		return nil
	}
	return fmt.Errorf("%s must be one of %v, got %q", fieldName, allowed, value)
}

// IntBetween returns nil if value is between min and max inclusive.
// Use for numeric range validation (e.g. level 0-10).
func IntBetween(fieldName string, value, min, max int) error {
	if value >= min && value <= max {
		return nil
	}
	return fmt.Errorf("%s must be between %d and %d inclusive, got %d", fieldName, min, max, value)
}
