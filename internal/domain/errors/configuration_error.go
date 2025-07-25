package errors

import (
	"errors"
	"fmt"
)

// ConfigurationError represents an error related to invalid or missing configuration fields.
type ConfigurationError struct {
	Field   string
	Message string
	Cause   error
}

// Error implements the error interface
func (e *ConfigurationError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("configuration error in field '%s': %s", e.Field, e.Message)
	}

	return fmt.Sprintf("configuration error: %s", e.Message)
}

// Unwrap returns the underlying cause of the ConfigurationError, enabling error unwrapping for more detailed error inspection.
func (e *ConfigurationError) Unwrap() error {
	return e.Cause
}

// NewConfigurationError creates and returns a new ConfigurationError with the specified field, message, and optional cause.
func NewConfigurationError(field, message string, cause error) *ConfigurationError {
	return &ConfigurationError{
		Field:   field,
		Message: message,
		Cause:   cause,
	}
}

// IsConfigurationError checks if the given error is of type ConfigurationError or wraps a ConfigurationError.
func IsConfigurationError(err error) bool {
	var configurationError *ConfigurationError
	ok := errors.As(err, &configurationError)

	return ok
}
