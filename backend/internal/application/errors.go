// Package application contains use-case contracts and application-level errors.
package application

import (
	"errors"
	"fmt"
)

// ErrorCode is stable across transports so adapters can map errors consistently.
type ErrorCode string

const (
	ErrorCodeValidation   ErrorCode = "VALIDATION_ERROR"
	ErrorCodeUnauthorized ErrorCode = "UNAUTHORIZED"
	ErrorCodeForbidden    ErrorCode = "FORBIDDEN"
	ErrorCodeNotFound     ErrorCode = "NOT_FOUND"
	ErrorCodeConflict     ErrorCode = "CONFLICT"
	ErrorCodeUnavailable  ErrorCode = "UNAVAILABLE"
	ErrorCodeInternal     ErrorCode = "INTERNAL"
)

// Error is returned by use cases when a transport needs stable error semantics.
type Error struct {
	Code    ErrorCode
	Message string
	Cause   error
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	if e.Cause == nil {
		return e.Message
	}

	return fmt.Sprintf("%s: %v", e.Message, e.Cause)
}

func (e *Error) Unwrap() error {
	if e == nil {
		return nil
	}

	return e.Cause
}

// NewError creates an application error with optional cause wrapping.
func NewError(code ErrorCode, message string, cause error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Cause:   cause,
	}
}

// CodeOf returns the stable application error code for err.
func CodeOf(err error) ErrorCode {
	var appErr *Error
	if errors.As(err, &appErr) {
		return appErr.Code
	}

	return ErrorCodeInternal
}
