package entities

import (
	"encoding/json"
	"fmt"
	"github.com/SHIMA0111/kanjo/internal/domain/utils"
	"slices"
)

// Config represents the configuration for a calculation
type Config struct {
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	Creator      string              `json:"creator"`
	Source       string              `json:"source"`
	SheetID      string              `json:"sheetID,omitempty"`
	Filter       string              `json:"filter,omitempty"`
	MergeColumns []MergeConfig       `json:"mergeColumns,omitempty"`
	Aggregations []AggregationConfig `json:"aggregations,omitempty"`
	OutputFormat string              `json:"outputFormat"`
}

// MergeConfig defines how to merge columns
type MergeConfig struct {
	FirstColumn      string   `json:"firstColumn"`
	SecondColumn     string   `json:"secondColumn"`
	Strategy         string   `json:"strategy"`
	DefaultValues    []string `json:"defaultValues,omitempty"`
	ResultColumnName string   `json:"resultColumnName,omitempty"`
}

// AggregationConfig defines how to aggregate data
type AggregationConfig struct {
	GroupingColumns []string      `json:"groupingColumns"`
	Aggregations    []Aggregation `json:"aggregations"`
}

// Aggregation defines a specific aggregation operation
type Aggregation struct {
	Column          string `json:"column"`
	AggregateMethod string `json:"aggregateMethod"`
	ResultName      string `json:"resultName"`
}

// Validate checks the Config object for required fields and sets default values where applicable.
// It validates nested MergeColumns and Aggregations configurations as well. Errors are returned for invalid cases.
func (c *Config) Validate() error {
	if c.Name == "" {
		c.Name = "UntitledConfig_" + utils.RandomString(10)
	}
	if c.Source == "" {
		return fmt.Errorf("source is required")
	}
	if c.SheetID == "" && c.Source == "googlesheets" {
		return fmt.Errorf("sheetID is required for Google Sheets source")
	}
	if c.OutputFormat == "" {
		c.OutputFormat = "csv"
	}

	// Validate all mergeColumns setting
	for i, mergeColumn := range c.MergeColumns {
		if err := mergeColumn.Validate(); err != nil {
			return fmt.Errorf("mergeColumn[%d]: %w", i, err)
		}
	}

	// Validate all aggregations setting
	for i, aggregation := range c.Aggregations {
		if err := aggregation.Validate(); err != nil {
			return fmt.Errorf("aggregation[%d]: %w", i, err)
		}
	}

	return nil
}

// Validate checks the MergeConfig for required fields, sets appropriate defaults, and validates the strategy field.
func (m *MergeConfig) Validate() error {
	if m.FirstColumn == "" {
		return fmt.Errorf("firstColumn is required")
	}
	if m.SecondColumn == "" {
		return fmt.Errorf("secondColumn is required")
	}
	if m.Strategy == "" {
		m.Strategy = "concat"
	}
	if m.ResultColumnName == "" {
		m.ResultColumnName = m.FirstColumn + "_" + m.SecondColumn
	}

	validateStrategies := []string{"concat", "sum", "first", "second"}
	if !slices.Contains(validateStrategies, m.Strategy) {
		return fmt.Errorf("invalid strategy '%s', strategy must be one of %v", m.Strategy, validateStrategies)
	}

	return nil
}

// Validate checks if the AggregationConfig instance has valid GroupingColumns and Aggregations and validates each aggregation.
func (ac *AggregationConfig) Validate() error {
	if len(ac.GroupingColumns) == 0 {
		return fmt.Errorf("groupingColumns cannot be empty")
	}
	if len(ac.Aggregations) == 0 {
		return fmt.Errorf("aggregations cannot be empty")
	}

	for i, aggregation := range ac.Aggregations {
		if err := aggregation.Validate(); err != nil {
			return fmt.Errorf("aggregation[%d]: %w", i, err)
		}
	}

	return nil
}

// Validate ensures that the Aggregation instance has valid values and performs the necessary validations on its fields.
func (a *Aggregation) Validate() error {
	if a.Column == "" {
		return fmt.Errorf("column is required")
	}
	if a.AggregateMethod == "" {
		return fmt.Errorf("aggregateMethod is required")
	}
	if a.ResultName == "" {
		a.ResultName = a.Column + "_" + a.AggregateMethod
	}

	validateAggregateMethods := []string{"sum", "avg", "min", "max", "count", "median"}
	if !slices.Contains(validateAggregateMethods, a.AggregateMethod) {
		return fmt.Errorf("invalid aggregateMethod '%s', aggregateMethod must be one of %v", a.AggregateMethod, validateAggregateMethods)
	}

	return nil
}

// ToJSON converts the Config object into a formatted JSON string. Returns an error if marshaling fails.
func (c *Config) ToJSON() (string, error) {
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal Config to JSON: %w", err)
	}

	return string(data), nil
}

// FromJSON parses a JSON string and populates the Config struct. Returns an error if unmarshalling or validation fails.
func (c *Config) FromJSON(jsonString string) error {
	if err := json.Unmarshal([]byte(jsonString), c); err != nil {
		return fmt.Errorf("failed to unmarshal JSON to Config: %w", err)
	}

	return c.Validate()
}
