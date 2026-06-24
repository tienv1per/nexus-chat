import type { ChatInitialData, Conversation, Message, User } from "@/lib/types";

export const seedUsers: User[] = [
  {
    id: "usr_alice",
    username: "alice",
    displayName: "Alice Tran",
    email: "alice@nexus.local",
    initials: "AT",
    avatarColor: "teal",
    role: "Owner",
    lastSeenAt: null,
    presence: "ONLINE"
  },
  {
    id: "usr_bob",
    username: "bob",
    displayName: "Bob Nguyen",
    email: "bob@nexus.local",
    initials: "BN",
    avatarColor: "indigo",
    role: "Member",
    lastSeenAt: "2026-06-24T08:32:00+07:00",
    presence: "LAST_SEEN"
  },
  {
    id: "usr_charlie",
    username: "charlie",
    displayName: "Charlie Le",
    email: "charlie@nexus.local",
    initials: "CL",
    avatarColor: "orange",
    role: "Test",
    lastSeenAt: "2026-06-23T22:10:00+07:00",
    presence: "OFFLINE"
  }
];

export const seedConversations: Conversation[] = [
  {
    id: "conv_bob",
    type: "ONE_TO_ONE",
    title: "Bob Nguyen",
    description: "1:1 product implementation thread",
    members: [
      { userId: "usr_alice", role: "ADMIN", lastSeenSequence: 41 },
      { userId: "usr_bob", role: "MEMBER", lastSeenSequence: 38 }
    ],
    lastMessagePreview: "Let's keep the ACK payload visible in the UI.",
    lastMessageAt: "10:42",
    unreadCount: 2,
    systemTag: "Direct"
  },
  {
    id: "conv_infra",
    type: "GROUP",
    title: "Infra Review Room",
    description: "Kafka, Redis, Postgres, and WebSocket diagnostics",
    members: [
      { userId: "usr_alice", role: "ADMIN", lastSeenSequence: 57 },
      { userId: "usr_bob", role: "MEMBER", lastSeenSequence: 55 },
      { userId: "usr_charlie", role: "MEMBER", lastSeenSequence: 49 }
    ],
    lastMessagePreview: "Postgres history is now the durable recovery path.",
    lastMessageAt: "09:18",
    unreadCount: 0,
    systemTag: "Group"
  },
  {
    id: "conv_media",
    type: "GROUP",
    title: "Media QA",
    description: "Upload previews, MIME checks, and local storage",
    members: [
      { userId: "usr_alice", role: "ADMIN", lastSeenSequence: 16 },
      { userId: "usr_charlie", role: "MEMBER", lastSeenSequence: 14 }
    ],
    lastMessagePreview: "Image previews should not shift the composer.",
    lastMessageAt: "Yesterday",
    unreadCount: 1,
    systemTag: "Media"
  }
];

export const seedMessages: Record<string, Message[]> = {
  conv_bob: [
    {
      messageId: "msg_39",
      conversationId: "conv_bob",
      senderId: "usr_bob",
      sequence: 39,
      messageKind: "TEXT",
      body: "The conversation list feels stronger when unread and presence are visible at the same time.",
      media: null,
      createdAt: "2026-06-24T10:36:00+07:00",
      status: "DELIVERED"
    },
    {
      messageId: "msg_40",
      conversationId: "conv_bob",
      senderId: "usr_alice",
      sequence: 40,
      messageKind: "TEXT",
      body: "Agreed. I also want sequence numbers exposed so reconnect behavior is obvious during demos.",
      media: null,
      createdAt: "2026-06-24T10:39:00+07:00",
      status: "DELIVERED"
    },
    {
      messageId: "msg_41",
      conversationId: "conv_bob",
      senderId: "usr_bob",
      sequence: 41,
      messageKind: "TEXT",
      body: "Let's keep the ACK payload visible in the UI.",
      media: null,
      createdAt: "2026-06-24T10:42:00+07:00",
      status: "DELIVERED"
    }
  ],
  conv_infra: [
    {
      messageId: "msg_54",
      conversationId: "conv_infra",
      senderId: "usr_charlie",
      sequence: 54,
      messageKind: "TEXT",
      body: "Redis presence TTL is still the fast path. Postgres only answers last-seen fallback now.",
      media: null,
      createdAt: "2026-06-24T09:04:00+07:00",
      status: "DELIVERED"
    },
    {
      messageId: "msg_55",
      conversationId: "conv_infra",
      senderId: "usr_alice",
      sequence: 55,
      messageKind: "FILE",
      body: "Attached the Postgres timeline migration sketch for review.",
      media: {
        id: "media_schema",
        filename: "message-timeline-schema.sql",
        mimeType: "text/sql",
        sizeLabel: "9 KB"
      },
      createdAt: "2026-06-24T09:12:00+07:00",
      status: "DELIVERED"
    },
    {
      messageId: "msg_56",
      conversationId: "conv_infra",
      senderId: "usr_bob",
      sequence: 56,
      messageKind: "TEXT",
      body: "The UI can sort by sequence and keep realtime arrival as just another insert path.",
      media: null,
      createdAt: "2026-06-24T09:16:00+07:00",
      status: "DELIVERED"
    },
    {
      messageId: "msg_57",
      conversationId: "conv_infra",
      senderId: "usr_alice",
      sequence: 57,
      messageKind: "TEXT",
      body: "Postgres history is now the durable recovery path.",
      media: null,
      createdAt: "2026-06-24T09:18:00+07:00",
      status: "SENT"
    }
  ],
  conv_media: [
    {
      messageId: "msg_15",
      conversationId: "conv_media",
      senderId: "usr_charlie",
      sequence: 15,
      messageKind: "TEXT",
      body: "Image previews should not shift the composer.",
      media: null,
      createdAt: "2026-06-23T18:22:00+07:00",
      status: "DELIVERED"
    }
  ]
};

export function buildInitialData(activeConversationId = "conv_bob"): ChatInitialData {
  return {
    currentUser: seedUsers[0],
    users: seedUsers,
    conversations: seedConversations,
    messagesByConversation: seedMessages,
    activeConversationId
  };
}
