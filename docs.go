// Package errors provides enhanced error handling capabilities with stack trace support.
//
// This package extends Go's standard error handling by providing automatic stack trace
// collection and enhanced error formatting. It offers drop-in replacements for standard
// error functions while adding valuable debugging information.
//
// Core Functions:
//   - [New]: Creates a new error with stack trace information
//   - [Errorf]: Formats an error with stack trace support and %w verb handling
//   - [Join]: Combines multiple errors into a single error
//   - [Is], [As], [Unwrap]: Standard error inspection functions (delegated to std library)
//   - [SprintStackTrace]: Extracts and formats stack trace from errors
//
// Stack Traces:
//
// Stack traces are automatically added to errors created by this package unless
// the error already has a stack trace or
// the error is created in a sentinel context (like package initialization).
// Stack traces can be formatted using the %+v verb with fmt package or extracted using [SprintStackTrace].
//
// Error Formatting:
//
// Errors support enhanced formatting through the fmt package:
//   - %s, %q: Basic error message
//   - %v: Error message
//   - %+v: Error message with full stack trace
//   - %#v: Go syntax representation
//
// Compatibility:
//
// This package provides deprecated compatibility functions for migration from
// [github.com/pkg/errors], but using the core API is recommended for new code.
package errors
