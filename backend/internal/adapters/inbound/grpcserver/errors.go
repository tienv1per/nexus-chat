package grpcserver

import "mini-grpc/backend/internal/application"

// CodeName returns the gRPC status category that should be used for err.
//
// The concrete google.golang.org/grpc/status dependency is added with generated
// protobuf contracts in Phase 5; keeping this mapping here preserves the adapter
// boundary without pulling gRPC into the core yet.
func CodeName(err error) string {
	switch application.CodeOf(err) {
	case application.ErrorCodeValidation:
		return "InvalidArgument"
	case application.ErrorCodeUnauthorized:
		return "Unauthenticated"
	case application.ErrorCodeForbidden:
		return "PermissionDenied"
	case application.ErrorCodeNotFound:
		return "NotFound"
	case application.ErrorCodeConflict:
		return "AlreadyExists"
	case application.ErrorCodeUnavailable:
		return "Unavailable"
	default:
		return "Internal"
	}
}
