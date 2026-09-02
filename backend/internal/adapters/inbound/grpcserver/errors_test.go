package grpcserver

import (
	"testing"

	"mini-grpc/backend/internal/application"
)

func TestCodeNameMapsApplicationErrors(t *testing.T) {
	err := application.NewError(application.ErrorCodeForbidden, "not a member", nil)

	if got := CodeName(err); got != "PermissionDenied" {
		t.Fatalf("CodeName = %q; want PermissionDenied", got)
	}
}
