package entities

import (
	"encoding/json"
	"fmt"
	"github.com/SHIMA0111/kanjo/internal/domain/utils"
	"slices"
)

// Config represents the configuration for a calculation
// Type represents the DataSource type; csv, googlesheets, etc.
// Source represents the identifier for the data source, such as the sheet ID for a Google Sheets source, filepath for csv.
type Config struct {
	Name         string              `json:"name"`
	Description  string              `json:"description"`
	Creator      string              `json:"creator"`
	Type         string              `json:"type"`
	Source       string              `json:"source"`
	Filters      []FilterConfig      `json:"filters,omitempty"`
	MergeColumns []MergeConfig       `json:"mergeColumns,omitempty"`
	Aggregations []AggregationConfig `json:"aggregations,omitempty"`
	OutputFormat string              `json:"outputFormat"`
}

// FilterConfig defines the structure for filtering operations based on a column, its value, and a specified operator.
type FilterConfig struct {
	Column          string `json:"column"`
	Value           string `json:"value"`
	Operator        string `json:"operator"`
	LogicalOperator string `json:"logicalOperator"` // LogicalOperator represents the way how to combine the next filter
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
	ResultName      string `json:"resultName,omitempty"`
}

// Validate checks the Config object for required fields and sets default values where applicable.
// It validates nested MergeColumns and Aggregations configurations as well. Errors are returned for invalid cases.
func (c *Config) Validate() error {
	if c.Name == "" {
		c.Name = "UntitledConfig_" + utils.RandomString(10)
	}
	if c.Type == "" {
		return fmt.Errorf("type is required, valid values are: %s", "csv, googlesheets")
	}
	if c.Source == "" {
		return fmt.Errorf("source is required")
	}
	if c.OutputFormat == "" {
		c.OutputFormat = "csv"
	}

	// This may be an implicit conversion and cause bugs. So commented out.
	//if len(c.Filters) == 1 && !slices.Contains([]string{"and", "or"}, c.Filters[0].LogicalOperator) {
	//	c.Filters[0].LogicalOperator = "and"
	//}

	for i, filter := range c.Filters {
		if err := filter.Validate(); err != nil {
			return fmt.Errorf("filter[%d]: %w", i, err)
		}
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

func (fc *FilterConfig) Validate() error {
	if fc.Column == "" {
		return fmt.Errorf("column is required")
	}

	if fc.Value == "" {
		return fmt.Errorf("value is required")
	}

	validateOperators := []string{"eq", "neq", "gt", "gte", "lt", "lte"}
	if !slices.Contains(validateOperators, fc.Operator) {
		return fmt.Errorf("invalid operator '%s', operator must be one of %v", fc.Operator, validateOperators)
	}

	validateLogicalOperators := []string{"and", "or"}
	if !slices.Contains(validateLogicalOperators, fc.LogicalOperator) {
		return fmt.Errorf("invalid logical operator '%s', operator must be one of %v", fc.LogicalOperator, validateLogicalOperators)
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
