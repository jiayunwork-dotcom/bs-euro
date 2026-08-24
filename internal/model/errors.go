package model

import "fmt"

// ErrorCode identifies the machine-readable category of a rejection.
type ErrorCode string

const (
	CodeInvalidSpot    ErrorCode = "invalid_spot"
	CodeInvalidStrike  ErrorCode = "invalid_strike"
	CodeInvalidTime    ErrorCode = "invalid_time"
	CodeInvalidVol     ErrorCode = "invalid_volatility"
	CodeInvalidRate    ErrorCode = "invalid_rate"
	CodeInvalidFlag    ErrorCode = "invalid_flag"
	CodeInvalidJSON    ErrorCode = "invalid_json"
	CodeMissingField   ErrorCode = "missing_field"
	CodeOutOfRange     ErrorCode = "out_of_range"
	CodeUnknownCommand ErrorCode = "unknown_command"
)

// Error is a structured validation or calculation error.
type Error struct {
	Code    ErrorCode
	Field   string
	Value   float64
	Message string
}

// Error implements the error interface.
func (e *Error) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("%s: %s (field %s=%g)", e.Code, e.Message, e.Field, e.Value)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// NewError builds a structured error.
func NewError(code ErrorCode, message string) *Error {
	return &Error{Code: code, Message: message}
}

// WithField attaches the invalid field name and value.
func (e *Error) WithField(field string, value float64) *Error {
	e.Field = field
	e.Value = value
	return e
}

// IsCode checks whether an error has the given code.
func IsCode(err error, code ErrorCode) bool {
	if err == nil {
		return false
	}
	me, ok := err.(*Error)
	if !ok {
		return false
	}
	return me.Code == code
}

// Code extracts a stable code from any error.
func Code(err error) ErrorCode {
	if err == nil {
		return ""
	}
	if me, ok := err.(*Error); ok {
		return me.Code
	}
	return CodeUnknownCommand
}

// MissingField builds a missing-field error.
func MissingField(name string) *Error {
	return NewError(CodeMissingField, fmt.Sprintf("missing required field %q", name))
}
