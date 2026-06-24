export type UserID = string;
export type ConversationID = string;
export type MessageID = string;

export type PresenceStatus = "ONLINE" | "LAST_SEEN" | "OFFLINE" | "UNKNOWN";
export type ConversationType = "ONE_TO_ONE" | "GROUP";
export type MessageKind = "TEXT" | "IMAGE" | "VIDEO" | "FILE";
export type DeliveryStatus = "PENDING" | "SENT" | "DELIVERED" | "FAILED";

export type User = {
  id: UserID;
  username: string;
  displayName: string;
  email: string;
  initials: string;
  avatarColor: string;
  role: "Owner" | "Member" | "Test";
  lastSeenAt: string | null;
  presence: PresenceStatus;
};

export type ConversationMember = {
  userId: UserID;
  role: "ADMIN" | "MEMBER";
  lastSeenSequence: number;
};

export type MediaObject = {
  id: string;
  filename: string;
  mimeType: string;
  sizeLabel: string;
  progress?: number;
};

export type Message = {
  messageId?: MessageID;
  clientMsgId?: string;
  conversationId: ConversationID;
  senderId: UserID;
  sequence: number;
  messageKind: MessageKind;
  body: string;
  media?: MediaObject | null;
  createdAt: string;
  status: DeliveryStatus;
};

export type Conversation = {
  id: ConversationID;
  type: ConversationType;
  title: string;
  description: string;
  members: ConversationMember[];
  lastMessagePreview: string;
  lastMessageAt: string;
  unreadCount: number;
  systemTag: string;
};

export type ChatInitialData = {
  currentUser: User;
  users: User[];
  conversations: Conversation[];
  messagesByConversation: Record<ConversationID, Message[]>;
  activeConversationId: ConversationID;
};

export type SendDraft = {
  conversationId: ConversationID;
  senderId: UserID;
  body: string;
  media?: MediaObject | null;
  clientMsgId: string;
  sequenceHint: number;
};

export type MessageAck = {
  clientMsgId: string;
  messageId: MessageID;
  sequence: number;
  status: Extract<DeliveryStatus, "SENT">;
};
