package application

import (
	"context"
	"io"
	"time"

	"mini-grpc/backend/internal/domain"
)

// ConversationMetadataStore owns conversation metadata and membership reads/writes.
type ConversationMetadataStore interface {
	GetConversation(ctx context.Context, id domain.ConversationID) (domain.Conversation, error)
	ListConversations(ctx context.Context, userID domain.UserID) ([]domain.Conversation, error)
	ListMembers(ctx context.Context, conversationID domain.ConversationID) ([]domain.ConversationMember, error)
	IsActiveMember(
		ctx context.Context,
		conversationID domain.ConversationID,
		userID domain.UserID,
	) (bool, error)
	MarkLastSeen(
		ctx context.Context,
		conversationID domain.ConversationID,
		userID domain.UserID,
		sequence domain.Sequence,
	) error
}

// MessageTimelineStore owns durable message history and delivery state.
type MessageTimelineStore interface {
	InsertMessage(ctx context.Context, message domain.Message) (domain.Message, error)
	FindMessageByClientID(
		ctx context.Context,
		conversationID domain.ConversationID,
		senderID domain.UserID,
		clientMsgID domain.ClientMessageID,
	) (domain.Message, error)
	ListMessagesBefore(
		ctx context.Context,
		conversationID domain.ConversationID,
		before domain.Sequence,
		limit int,
	) ([]domain.Message, error)
	UpsertDelivery(ctx context.Context, delivery domain.MessageDelivery) error
}

// SequenceAllocator allocates monotonically increasing sequence numbers per conversation.
type SequenceAllocator interface {
	NextSequence(ctx context.Context, conversationID domain.ConversationID) (domain.Sequence, error)
}

// DedupStore protects send-message retries from creating duplicate durable rows.
type DedupStore interface {
	ReserveClientMessage(
		ctx context.Context,
		conversationID domain.ConversationID,
		senderID domain.UserID,
		clientMsgID domain.ClientMessageID,
	) (bool, error)
	ReleaseClientMessage(
		ctx context.Context,
		conversationID domain.ConversationID,
		senderID domain.UserID,
		clientMsgID domain.ClientMessageID,
	) error
}

// EventPublisher publishes durable application events to the async pipeline.
type EventPublisher interface {
	Publish(ctx context.Context, event domain.Event) error
}

// PresenceRegistry owns realtime session and presence state.
type PresenceRegistry interface {
	RegisterSession(ctx context.Context, session domain.Session) error
	RefreshSession(ctx context.Context, connectionID domain.ConnectionID) error
	RemoveSession(ctx context.Context, connectionID domain.ConnectionID) error
	GetPresence(ctx context.Context, userID domain.UserID) (domain.Presence, error)
	MarkLastSeen(ctx context.Context, userID domain.UserID) error
}

// MediaStorage stores media bytes and returns metadata suitable for PostgreSQL.
type MediaStorage interface {
	Save(ctx context.Context, input SaveMediaInput) (domain.MediaObject, error)
	Open(ctx context.Context, mediaID domain.MediaID) (io.ReadCloser, domain.MediaObject, error)
	Delete(ctx context.Context, mediaID domain.MediaID) error
}

// SaveMediaInput describes one local media write.
type SaveMediaInput struct {
	OwnerUserID      domain.UserID
	OriginalFilename string
	MimeType         string
	SizeBytes        int64
	Reader           io.Reader
}

// Clock makes time deterministic in use-case tests.
type Clock interface {
	Now() time.Time
}

// IDGenerator centralizes ID generation for messages, media, connections, and events.
type IDGenerator interface {
	NewMessageID() domain.MessageID
	NewMediaID() domain.MediaID
	NewConnectionID() domain.ConnectionID
	NewEventID() domain.EventID
}
