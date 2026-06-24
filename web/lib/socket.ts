import type { Message, MessageAck, SendDraft } from "@/lib/types";

export type SocketStatus = "connecting" | "connected" | "reconnecting" | "offline";

export function makeClientMessageId() {
  const random = Math.random().toString(36).slice(2, 9);
  return `client_${Date.now()}_${random}`;
}

export function createAck(draft: SendDraft): MessageAck {
  return {
    clientMsgId: draft.clientMsgId,
    messageId: `msg_${draft.sequenceHint}_${draft.clientMsgId.slice(-4)}`,
    sequence: draft.sequenceHint,
    status: "SENT"
  };
}

export function createRecipientEcho(draft: SendDraft): Message {
  return {
    messageId: `msg_echo_${draft.sequenceHint + 1}`,
    conversationId: draft.conversationId,
    senderId: draft.conversationId === "conv_media" ? "usr_charlie" : "usr_bob",
    sequence: draft.sequenceHint + 1,
    messageKind: "TEXT",
    body: draft.conversationId === "conv_infra"
      ? "Got it. I will verify this against the local event flow."
      : "Received. The mock socket inserted this after ACK to test ordering.",
    media: null,
    createdAt: new Date().toISOString(),
    status: "DELIVERED"
  };
}

export function sortMessagesBySequence(messages: Message[]) {
  return [...messages].sort((a, b) => {
    if (a.sequence === b.sequence) {
      return a.createdAt.localeCompare(b.createdAt);
    }

    return a.sequence - b.sequence;
  });
}
