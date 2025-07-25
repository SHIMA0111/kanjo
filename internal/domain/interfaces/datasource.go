package interfaces

import (
	"context"
	"github.com/go-gota/gota/dataframe"
)

// DataSourceConfig represents the configuration required for retrieving data from a specific source.
type DataSourceConfig struct {
	Type   string `json:"type"`
	Source string `json:"source"`
	Range  string `json:"range"`
}

// DataSource abstracts data retrieval from various sources
// Implementations should handle authentication, connection management,
// and data format conversion to DataFrame
type DataSource interface {
	// Fetch retrieves data from the configured source and return it as a DataFrame
	// ctx: context for cancellation and timeout control
	// config: source-specific configuration (SheetID, file path, etc.)
	// Returns: DataFrame containing the fetched data, or error if fetch fails
	//
	// Implementation notes:
	// - Should handle authentication automatically (OAuth2, API keys, etc.)
	// - Should validate source configuration before attempting fetch
	// - Should return descriptive errors for common failure scenarios
	// - Should support context cancellation for long-running operations
	Fetch(ctx context.Context, config DataSourceConfig) (*dataframe.DataFrame, error)

	// Validate checks if the source configuration is valid
	// config: source configuration to valid
	// Returns: error if configuration is invalid, nil if valid
	//
	// Should validate:
	// - Required fields are present
	// - a Source type is supported
	// - Credentials/authentication are available
	// - Source is accessible (optional connectivity check)
	Validate(config DataSourceConfig) error

	// GetSourceInfo returns human-readable information about the data source
	// config: source configuration
	// Returns: descriptive string about the source (e.g., "Google Sheets: MySheet (100 rows)")
	GetSourceInfo(config DataSourceConfig) string

	// SupportedTypes returns a list of the source types this implementation supports
	// Returns: slice of supported type strings (e.g., ["googlesheets", "csv])
	SupportedTypes() []string
}
