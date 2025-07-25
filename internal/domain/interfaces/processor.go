package interfaces

import (
	"context"
	"github.com/SHIMA0111/kanjo/internal/domain/entities"
	"github.com/go-gota/gota/dataframe"
)

// Processor handles data transformations including filtering, merging, and aggregation
// All operations should be performed in a way that preserves data integrity
// and provides meaningful error messages for debugging
type Processor interface {
	// Filter applies filter expressions to the data and returns filtered DataFrame
	// data: input DataFrame to filter
	// config: slice of filter configurations defining how to filter columns
	// Returns: filtered DataFrame or error if filter expression is invalid
	//
	// Supported filter operations:
	// - Equality: ==, !=
	// - Comparison: <, <=, >, >=
	// - String operations: contains, startWith, endWith
	//
	// Supported logical operator combine filters:
	// - or: Combine the next filter with OR condition
	// - and: Combine the next filter with AND condition
	//
	// Implementation notes:
	// - Should validate filter expression syntax before applying
	// - Should handle type conversions automatically (string to date, etc.)
	// - Should provide detailed error messages for invalid expressions
	// - Should preserve original column type in filtered result
	Filter(ctx context.Context, data *dataframe.DataFrame, config []entities.FilterConfig) (*dataframe.DataFrame, error)

	// Merge combines columns according to the provided merge configurations
	// data: input DataFrame to perform merge operations on
	// config: slice of merge configurations defining how to combine columns
	// Returns: DataFrame with merged columns or error if merge fails
	//
	// Supported merge strategies:
	// - concat: Concatenate the columns data
	// - sum: Sum the columns data (if specified non-numeric column, returns error)
	// - first: Prior the first column data, and if the first column is missing, the second value represented
	// - second: Prior the second column data (the thought is the same as the `first` strategy)
	//
	// Implementation notes:
	// - Should validate that source columns exist before merging
	// - Should handle missing values gracefully using default values
	// - Should preserve data type when possible
	// - Should create new result columns without modifying originals
	Merge(ctx context.Context, data *dataframe.DataFrame, config []entities.MergeConfig) (*dataframe.DataFrame, error)

	// Aggregate performs grouping and aggregation operations on the data
	// data: input DataFrame to aggregate
	// config: slice of aggregate
	// Returns: aggregated DataFrame or error if aggregation fails
	//
	// Supported aggregation method:
	// - sum: Total of the specified column data each group
	// - avg: Average of the specified column data each group
	// - min: Minimum data of the specified column data each group
	// - max: Maximum data of the specified column data each group
	// - count: Counting data number of the specified column data in each group
	// - median: Median of the specified column data each group
	//
	// Implementation notes:
	// - Should validate that target columns exist and are appropriate for aggregation method
	// - Should handle multiple grouping columns correctly
	// - Should preserve grouping column values in result
	// - Should handle null/missing values appropriately for each aggregation type
	Aggregate(ctx context.Context, data *dataframe.DataFrame, config []entities.AggregationConfig) (*dataframe.DataFrame, error)

	// ValidateExpression checks if a filter expression is syntactically valid
	// expression: filter expression to validate
	// columnNames: available column names for validate
	// Returns: error if expression is invalid, nil if valid
	ValidateExpression(expression string, columnNames []string) error

	// GetSupportedOperators returns a list of supported filter operators
	// Returns: slice of operator strings (e.g., ["==", "!=", "<", "<=", ">", ">="])
	GetSupportedOperators() []string

	// GetSupportedMergeStrategies returns a list of merge strategies
	// Returns: slice of strategies (e.g., ["concat", "sum", "first", "second"])
	GetSupportedMergeStrategies() []string

	// GetSupportedAggregations returns a list of supported aggregation types
	// Returns: slice of aggregation type strings (e.g., ["sum", "avg", "min", "max", "count"])
	GetSupportedAggregations() []string
}
