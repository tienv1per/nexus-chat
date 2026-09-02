package websocket

import "mini-grpc/backend/internal/application"

const (
	// CloseUnsupportedData mirrors RFC 6455 close code 1003.
	CloseUnsupportedData = 1003
	// ClosePolicyViolation mirrors RFC 6455 close code 1008.
	ClosePolicyViolation = 1008
	// CloseInternalError mirrors RFC 6455 close code 1011.
	CloseInternalError = 1011
)

// CloseCode maps application errors to WebSocket close categories.
func CloseCode(err error) int {
	switch application.CodeOf(err) {
	case application.ErrorCodeValidation:
		return CloseUnsupportedData
	case application.ErrorCodeUnauthorized, application.ErrorCodeForbidden:
		return ClosePolicyViolation
	case application.ErrorCodeNotFound, application.ErrorCodeConflict, application.ErrorCodeUnavailable:
		return ClosePolicyViolation
	default:
		return CloseInternalError
	}
}
