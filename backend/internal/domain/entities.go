package domain

import "time"

// User is the identity shape shared by dev auth, conversation, and presence flows.
type User struct {
	ID          UserID
	Username    string
	DisplayName string
	Email       string
	Initials    string
	AvatarColor string
	Role        string
	LastSeenAt  *time.Time
	Presence    PresenceStatus
}

// ConversationMember is a user's membership state inside a conversation.
type ConversationMember struct {
	UserID           UserID
	Role             MemberRole
	LastSeenSequence Sequence
}

// Conversation contains metadata needed by conversation list and thread views.
type Conversation struct {
	ID                 ConversationID
	Type               ConversationType
	Title              string
	Description        string
	Members            []ConversationMember
	LastMessagePreview string
	LastMessageAt      *time.Time
	UnreadCount        int
	SystemTag          string
}

// MediaObject stores metadata for a file payload. Bytes stay outside PostgreSQL.
type MediaObject struct {
	ID               MediaID
	OwnerUserID      UserID
	OriginalFilename string
	MimeType         string
	SizeBytes        int64
	SizeLabel        string
	StoragePath      string
	CreatedAt        time.Time
}

// Message is the durable timeline record recovered after reconnect.
type Message struct {
	ID             MessageID
	ClientMsgID    ClientMessageID
	ConversationID ConversationID
	SenderID       UserID
	Sequence       Sequence
	Kind           MessageKind
	Body           string
	Media          *MediaObject
	CreatedAt      time.Time
	Status         DeliveryStatus
}

// MessageDelivery is per-recipient delivery state.
type MessageDelivery struct {
	MessageID   MessageID
	RecipientID UserID
	Status      DeliveryStatus
	DeliveredAt *time.Time
	ReadAt      *time.Time
	UpdatedAt   time.Time
}

// Session represents one active realtime connection.
type Session struct {
	ConnectionID ConnectionID
	UserID       UserID
	DeviceID     DeviceID
	ConnectedAt  time.Time
	ExpiresAt    time.Time
}

// Presence is the current or fallback presence view for one user.
type Presence struct {
	UserID     UserID
	Status     PresenceStatus
	LastSeenAt *time.Time
}

// Event is the application-level shape for async Kafka publication.
type Event struct {
	ID         EventID
	Type       string
	Key        string
	Headers    map[string]string
	Payload    []byte
	OccurredAt time.Time
}
