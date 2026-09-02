package websocket

import (
	"testing"

	"mini-grpc/backend/internal/application"
)

func TestCloseCodeMapsApplicationErrors(t *testing.T) {
	err := application.NewError(application.ErrorCodeUnauthorized, "missing token", nil)

	if got := CloseCode(err); got != ClosePolicyViolation {
		t.Fatalf("CloseCode = %d; want %d", got, ClosePolicyViolation)
	}
}
