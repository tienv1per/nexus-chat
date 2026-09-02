package application

import (
	"context"

	"mini-grpc/backend/internal/domain"
)

// ConversationQueries is the inbound port used by REST handlers for metadata reads.
type ConversationQueries interface {
	ListConversations(ctx context.Context, input ListConversationsInput) (ListConversationsOutput, error)
	ListMessages(ctx context.Context, input ListMessagesInput) (ListMessagesOutput, error)
}

// MessageCommands is the inbound port used by REST, gRPC, and WebSocket send paths.
type MessageCommands interface {
	SendMessage(ctx context.Context, input SendMessageInput) (SendMessageOutput, error)
	MarkLastSeen(ctx context.Context, input MarkLastSeenInput) error
}

// PresenceQueries is the inbound port used by REST and WebSocket session checks.
type PresenceQueries interface {
	GetPresence(ctx context.Context, input GetPresenceInput) (GetPresenceOutput, error)
}

// ListConversationsInput requests the current user's conversation list.
type ListConversationsInput struct {
	UserID domain.UserID
	Limit  int
}

// ListConversationsOutput returns metadata sorted by recent activity.
type ListConversationsOutput struct {
	Conversations []domain.Conversation
}

// ListMessagesInput requests a bounded page before a sequence cursor.
type ListMessagesInput struct {
	UserID         domain.UserID
	ConversationID domain.ConversationID
	Before         domain.Sequence
	Limit          int
}

// ListMessagesOutput returns timeline messages in descending sequence order.
type ListMessagesOutput struct {
	Messages []domain.Message
}

// SendMessageInput is the application command behind the realtime send path.
type SendMessageInput struct {
	ConversationID domain.ConversationID
	SenderID       domain.UserID
	ClientMsgID    domain.ClientMessageID
	Kind           domain.MessageKind
	Body           string
	MediaID        domain.MediaID
}

// SendMessageOutput is returned to callers as the ACK payload.
type SendMessageOutput struct {
	Message domain.Message
}

// MarkLastSeenInput updates read-position metadata.
type MarkLastSeenInput struct {
	UserID         domain.UserID
	ConversationID domain.ConversationID
	Sequence       domain.Sequence
}

// GetPresenceInput requests the best current presence state.
type GetPresenceInput struct {
	RequesterID domain.UserID
	UserID      domain.UserID
}

// GetPresenceOutput contains Redis presence or PostgreSQL last-seen fallback.
type GetPresenceOutput struct {
	Presence domain.Presence
}
