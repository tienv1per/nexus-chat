package domain

import (
	"strings"
	"testing"
)

func TestIDValidationRejectsBlankValues(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{name: "user", err: UserID(" ").Validate()},
		{name: "conversation", err: ConversationID("").Validate()},
		{name: "message", err: MessageID("").Validate()},
		{name: "device", err: DeviceID("\t").Validate()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				t.Fatal("Validate returned nil error")
			}
			if !strings.Contains(tt.err.Error(), "required") {
				t.Fatalf("error = %q; want required", tt.err)
			}
		})
	}
}

func TestEnumValidation(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{name: "conversation type", err: ConversationTypeGroup.Validate()},
		{name: "member role", err: MemberRoleAdmin.Validate()},
		{name: "message kind", err: MessageKindFile.Validate()},
		{name: "delivery status", err: DeliveryStatusDelivered.Validate()},
		{name: "presence status", err: PresenceStatusLastSeen.Validate()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err != nil {
				t.Fatalf("Validate returned error: %v", tt.err)
			}
		})
	}
}

func TestSequenceValidation(t *testing.T) {
	if err := Sequence(41).Validate(); err != nil {
		t.Fatalf("valid sequence returned error: %v", err)
	}

	if err := Sequence(0).Validate(); err == nil {
		t.Fatal("zero sequence returned nil error")
	}
}
