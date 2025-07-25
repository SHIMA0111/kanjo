package interfaces

// Logger defines the interface for logging operations
// Implementations should handle log formatting, levels, and output destinations
type Logger interface {
	// Debug logs a debug-level message
	// msg: log message
	// fields: additional structured data to include in the log entry
	Debug(msg string, fields map[string]interface{})

	// Info logs an info-level message
	// msg: log message
	// fields: additional structured data to include in the log entry
	Info(msg string, fields map[string]interface{})

	// Warn logs a warning-level message
	// msg: log message
	// fields: additional structured data to include in the log entry
	Warn(msg string, fields map[string]interface{})

	// Error logs an error-level message
	// msg: log message
	// fields: additional structured data to include in the log entry
	Error(msg string, fields map[string]interface{})
}
