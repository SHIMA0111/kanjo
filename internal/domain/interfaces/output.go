package interfaces

import (
	"context"
	"github.com/SHIMA0111/kanjo/internal/domain/entities"
	"github.com/go-gota/gota/dataframe"
)

// OutputConfig represents configuration for output formatting and destination
type OutputConfig struct {
	Format      string                 `json:"format"`                // "csv", "console", "json"
	Destination string                 `json:"destination,omitempty"` // file path for file outputs
	Options     map[string]interface{} `json:"options,omitempty"`     // format-specific options
}

// Output handles result output in various formats and destinations
// Implementations should handle file creation, formatting, and error recovery
type Output interface {
	// Write outputs the process result according to the provided configuration
	// ctx: context for cancellation and timeout control
	// result: process result containing data and metadata
	// config: output configuration specifying format and destination
	// Returns: error if write operation fails
	//
	// Implementation notes:
	// - Should create output directories if they don't exist
	// - Should handle file permissions and disk space issues gracefully
	// - Should validate output configuration before attempting writing
	// - Should support context cancellation for log-running operations
	// - Should preserve data formatting and types when possible
	Write(ctx context.Context, df *dataframe.DataFrame, config OutputConfig) error

	// Validate checks if the output configuration is valid for this output target
	// config: output configuration to validate
	// Returns: error if configuration is invalid, nil if valid
	//
	// Should validate:
	// - Output format is supported
	// - Destination path is writable (for file outputs)
	// - Required options are provided
	// - Option values are valid for the format
	Validate(config OutputConfig) error

	// SupportedFormats returns a list of output formats this supports
	// Returns: slice of supported format strings (e.g., ["csv", "json"])
	SupportedFormats() []string

	// GetFormatOptions returns available options for a specific output format
	// format: output format to get options for
	// Returns: map of option names to their descriptions
	GetFormatOptions(format string) map[string]string

	// Preview generates a preview of how the output would look without writing
	// result: process result to preview
	// config: output configuration
	// maxRows: maximum number of rows to include in preview (0 for all)
	// Returns: string preview of the output or error if preview fails
	Preview(result *entities.Processing, config OutputConfig, maxRows int) (string, error)
}
