package errors

import (
	"errors"
	"fmt"
)

// AuthenticationError represents an error encountered during the authentication process.
// It provides details about the error message, root cause, current retry attempt, and maximum retries allowed.
type AuthenticationError struct {
	Message string
	Cause   error
	// RetryAttempt indicates the current retry attempt
	RetryAttempt int

	// MaxRetries defines the maximum number of retry attempts allowed for an operation before failing.
	MaxRetries int
}

// Error implements the error interface
func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("authentication error: %s", e.Message)
}

// Unwrap returns the underlying cause of the AuthenticationError, enabling error unwrapping for further analysis.
func (e *AuthenticationError) Unwrap() error {
	return e.Cause
}

// IsRetryable determines if the error condition is retryable based on the current retry attempt and the maximum retries allowed.
func (e *AuthenticationError) IsRetryable() bool {
	return e.RetryAttempt < e.MaxRetries
}

// IncrementRetryAttempt increments the count of retry attempts for the current authentication error instance.
func (e *AuthenticationError) IncrementRetryAttempt() {
	e.RetryAttempt++
}

// NewAuthenticationError creates a new instance of AuthenticationError with the specified message and cause.
// RetryAttempt is initialized to 0, and MaxRetries is set to the default value of 3.
func NewAuthenticationError(message string, cause error) *AuthenticationError {
	return &AuthenticationError{
		Message:      message,
		Cause:        cause,
		RetryAttempt: 0,
		// The default of the MaxRetries is 3
		MaxRetries: 3,
	}
}

// NewAuthenticationErrorWithMaxRetries creates a new AuthenticationError initialized with the given message, cause, and maxRetries.
func NewAuthenticationErrorWithMaxRetries(message string, cause error, maxRetries int) *AuthenticationError {
	return &AuthenticationError{
		Message:      message,
		Cause:        cause,
		RetryAttempt: 0,
		MaxRetries:   maxRetries,
	}
}

// IsAuthenticationError checks if the provided error is of type AuthenticationError.
func IsAuthenticationError(err error) bool {
	var authenticationError *AuthenticationError
	ok := errors.As(err, &authenticationError)

	return ok
}
