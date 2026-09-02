package application

import (
	"errors"
	"testing"
)

func TestCodeOfReturnsApplicationCode(t *testing.T) {
	err := NewError(ErrorCodeNotFound, "conversation not found", nil)

	if got := CodeOf(err); got != ErrorCodeNotFound {
		t.Fatalf("CodeOf = %q; want %q", got, ErrorCodeNotFound)
	}
}

func TestErrorWrapsCause(t *testing.T) {
	cause := errors.New("postgres timeout")
	err := NewError(ErrorCodeUnavailable, "storage unavailable", cause)

	if !errors.Is(err, cause) {
		t.Fatal("application error does not wrap cause")
	}
}
