package errors

import (
	"errors"
	"fmt"
)

// DataProcessError represents an error occurring during a data processing step, including details about the step and recovery.
// Recoverable indicates if the error can be recovered from, allowing the process to potentially continue execution.
// RecoveryAction specifies the action or steps to be taken to recover from the error, if it is recoverable.
type DataProcessError struct {
	Step           string
	Message        string
	Cause          error
	Recoverable    bool
	RecoveryAction string
}

// Error implements error interface
func (e *DataProcessError) Error() string {
	if e.Step != "" {
		return fmt.Sprintf("data process error in step '%s': %s", e.Step, e.Message)
	}

	return fmt.Sprintf("data process error: %s", e.Message)
}

// Unwrap returns the underlying cause of the DataProcessError, if present.
func (e *DataProcessError) Unwrap() error {
	return e.Cause
}

// IsRecoverable checks if the DataProcessError can be recovered from, allowing the process to potentially continue execution.
func (e *DataProcessError) IsRecoverable() bool {
	return e.Recoverable
}

// GetRecoveryAction returns the recovery action or steps to address the error, if it is recoverable.
func (e *DataProcessError) GetRecoveryAction() string {
	return e.RecoveryAction
}

// NewDataProcessError creates and returns a new DataProcessError with the specified step, message, and underlying cause.
func NewDataProcessError(step, message string, cause error) *DataProcessError {
	return &DataProcessError{
		Step:           step,
		Message:        message,
		Cause:          cause,
		Recoverable:    false,
		RecoveryAction: "",
	}
}

// NewRecoverableDataProcessError creates a new recoverable DataProcessError with the provided step, message, cause, and recovery action.
func NewRecoverableDataProcessError(step, message string, cause error, recoveryAction string) *DataProcessError {
	return &DataProcessError{
		Step:           step,
		Message:        message,
		Cause:          cause,
		Recoverable:    true,
		RecoveryAction: recoveryAction,
	}
}

// IsDataProcessError checks if the given error is of type DataProcessError and returns true if it matches.
func IsDataProcessError(err error) bool {
	var dataProcessError *DataProcessError
	ok := errors.As(err, &dataProcessError)

	return ok
}
