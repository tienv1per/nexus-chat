// Package domain contains chat business types without transport or storage dependencies.
package domain

import (
	"fmt"
	"strings"
)

// UserID identifies an application user.
type UserID string

// ConversationID identifies a direct or group conversation.
type ConversationID string

// MessageID identifies a durable message.
type MessageID string

// MediaID identifies stored media metadata.
type MediaID string

// ClientMessageID identifies a client retry/deduplication key.
type ClientMessageID string

// DeviceID identifies one client device or browser tab family.
type DeviceID string

// ConnectionID identifies one active WebSocket connection.
type ConnectionID string

// EventID identifies an event published to the async pipeline.
type EventID string

// Sequence is a per-conversation message ordering value.
type Sequence int64

// ConversationType describes how membership is interpreted.
type ConversationType string

const (
	ConversationTypeOneToOne ConversationType = "ONE_TO_ONE"
	ConversationTypeGroup    ConversationType = "GROUP"
)

// MemberRole describes a user's role inside a conversation.
type MemberRole string

const (
	MemberRoleAdmin  MemberRole = "ADMIN"
	MemberRoleMember MemberRole = "MEMBER"
)

// MessageKind describes the payload shape of a message.
type MessageKind string

const (
	MessageKindText  MessageKind = "TEXT"
	MessageKindImage MessageKind = "IMAGE"
	MessageKindVideo MessageKind = "VIDEO"
	MessageKindFile  MessageKind = "FILE"
)

// DeliveryStatus describes best-effort delivery progress.
type DeliveryStatus string

const (
	DeliveryStatusPending   DeliveryStatus = "PENDING"
	DeliveryStatusSent      DeliveryStatus = "SENT"
	DeliveryStatusDelivered DeliveryStatus = "DELIVERED"
	DeliveryStatusFailed    DeliveryStatus = "FAILED"
)

// PresenceStatus describes the online state surfaced to clients.
type PresenceStatus string

const (
	PresenceStatusOnline   PresenceStatus = "ONLINE"
	PresenceStatusLastSeen PresenceStatus = "LAST_SEEN"
	PresenceStatusOffline  PresenceStatus = "OFFLINE"
	PresenceStatusUnknown  PresenceStatus = "UNKNOWN"
)

func (id UserID) Validate() error {
	return validateNonEmpty("user_id", string(id))
}

func (id ConversationID) Validate() error {
	return validateNonEmpty("conversation_id", string(id))
}

func (id MessageID) Validate() error {
	return validateNonEmpty("message_id", string(id))
}

func (id MediaID) Validate() error {
	return validateNonEmpty("media_id", string(id))
}

func (id ClientMessageID) Validate() error {
	return validateNonEmpty("client_msg_id", string(id))
}

func (id DeviceID) Validate() error {
	return validateNonEmpty("device_id", string(id))
}

func (id ConnectionID) Validate() error {
	return validateNonEmpty("connection_id", string(id))
}

func (id EventID) Validate() error {
	return validateNonEmpty("event_id", string(id))
}

func (s Sequence) Validate() error {
	if s <= 0 {
		return fmt.Errorf("sequence must be positive")
	}

	return nil
}

func (t ConversationType) Validate() error {
	switch t {
	case ConversationTypeOneToOne, ConversationTypeGroup:
		return nil
	default:
		return fmt.Errorf("invalid conversation_type %q", t)
	}
}

func (r MemberRole) Validate() error {
	switch r {
	case MemberRoleAdmin, MemberRoleMember:
		return nil
	default:
		return fmt.Errorf("invalid member_role %q", r)
	}
}

func (k MessageKind) Validate() error {
	switch k {
	case MessageKindText, MessageKindImage, MessageKindVideo, MessageKindFile:
		return nil
	default:
		return fmt.Errorf("invalid message_kind %q", k)
	}
}

func (s DeliveryStatus) Validate() error {
	switch s {
	case DeliveryStatusPending, DeliveryStatusSent, DeliveryStatusDelivered, DeliveryStatusFailed:
		return nil
	default:
		return fmt.Errorf("invalid delivery_status %q", s)
	}
}

func (s PresenceStatus) Validate() error {
	switch s {
	case PresenceStatusOnline, PresenceStatusLastSeen, PresenceStatusOffline, PresenceStatusUnknown:
		return nil
	default:
		return fmt.Errorf("invalid presence_status %q", s)
	}
}

func validateNonEmpty(name string, value string) error {
	if strings.TrimSpace(value) == "" {
		return fmt.Errorf("%s is required", name)
	}

	return nil
}
