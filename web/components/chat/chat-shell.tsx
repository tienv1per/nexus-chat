"use client";

import { CircleOff, Loader2, Menu, Radio, RotateCcw } from "lucide-react";
import { useCallback, useEffect, useMemo, useState } from "react";
import { useRouter } from "next/navigation";

import { ConversationList } from "@/components/chat/conversation-list";
import { DetailsRail } from "@/components/chat/details-rail";
import { MessageComposer } from "@/components/chat/message-composer";
import { MessageList } from "@/components/chat/message-list";
import { PresencePill } from "@/components/chat/presence-pill";
import { useChatSocket } from "@/hooks/use-chat-socket";
import { makeClientMessageId, sortMessagesBySequence } from "@/lib/socket";
import type { ChatInitialData, MediaObject, Message, MessageAck, User } from "@/lib/types";

type ChatShellProps = {
  initialData: ChatInitialData;
};

function getUserFromLocalStorage(users: User[], fallback: User) {
  if (typeof window === "undefined") {
    return fallback;
  }

  const stored = window.localStorage.getItem("nexuschat-user-id");
  return users.find((user) => user.id === stored) ?? fallback;
}

export function ChatShell({ initialData }: ChatShellProps) {
  const router = useRouter();
  const [currentUser, setCurrentUser] = useState(initialData.currentUser);
  const [messagesByConversation, setMessagesByConversation] = useState(initialData.messagesByConversation);
  const [isLoadingOlder, setIsLoadingOlder] = useState(false);

  const activeConversation = useMemo(() => {
    return initialData.conversations.find((conversation) => conversation.id === initialData.activeConversationId);
  }, [initialData.activeConversationId, initialData.conversations]);

  const activeMessages = messagesByConversation[initialData.activeConversationId] ?? [];

  useEffect(() => {
    setCurrentUser(getUserFromLocalStorage(initialData.users, initialData.currentUser));
  }, [initialData.currentUser, initialData.users]);

  const handleAck = useCallback(
    (ack: MessageAck) => {
      setMessagesByConversation((current) => {
        const messages = current[initialData.activeConversationId] ?? [];
        const nextMessages = messages.map((message) => {
          if (message.clientMsgId !== ack.clientMsgId) {
            return message;
          }

          return {
            ...message,
            messageId: ack.messageId,
            sequence: ack.sequence,
            status: ack.status
          };
        });

        return {
          ...current,
          [initialData.activeConversationId]: sortMessagesBySequence(nextMessages)
        };
      });
    },
    [initialData.activeConversationId]
  );

  const handleIncoming = useCallback(
    (incoming: Message) => {
      setMessagesByConversation((current) => {
        const messages = current[initialData.activeConversationId] ?? [];
        const exists = messages.some((message) => message.messageId === incoming.messageId);
        if (exists) {
          return current;
        }

        return {
          ...current,
          [initialData.activeConversationId]: sortMessagesBySequence([...messages, incoming])
        };
      });
    },
    [initialData.activeConversationId]
  );

  const socket = useChatSocket({
    conversationId: initialData.activeConversationId,
    onAck: handleAck,
    onIncoming: handleIncoming
  });

  const nextSequence = useMemo(() => {
    return activeMessages.reduce((max, message) => Math.max(max, message.sequence), 0) + 1;
  }, [activeMessages]);

  function sendMessage(body: string, media?: MediaObject | null) {
    const clientMsgId = makeClientMessageId();
    const optimisticMessage: Message = {
      clientMsgId,
      conversationId: initialData.activeConversationId,
      senderId: currentUser.id,
      sequence: nextSequence,
      messageKind: media ? "FILE" : "TEXT",
      body: body || "Attachment",
      media: media ?? null,
      createdAt: new Date().toISOString(),
      status: "PENDING"
    };

    setMessagesByConversation((current) => {
      const messages = current[initialData.activeConversationId] ?? [];
      return {
        ...current,
        [initialData.activeConversationId]: sortMessagesBySequence([...messages, optimisticMessage])
      };
    });

    socket.sendMessage({
      conversationId: initialData.activeConversationId,
      senderId: currentUser.id,
      body,
      media,
      clientMsgId,
      sequenceHint: nextSequence
    });
  }

  function loadOlderMessages() {
    setIsLoadingOlder(true);
    window.setTimeout(() => {
      setMessagesByConversation((current) => {
        const messages = current[initialData.activeConversationId] ?? [];
        const oldest = messages.reduce((min, message) => Math.min(min, message.sequence), Number.MAX_SAFE_INTEGER);
        const olderMessage: Message = {
          messageId: `msg_older_${oldest - 1}`,
          conversationId: initialData.activeConversationId,
          senderId: currentUser.id === "usr_alice" ? "usr_bob" : "usr_alice",
          sequence: oldest - 1,
          messageKind: "TEXT",
          body: "Earlier history loaded from the mock Postgres timeline.",
          media: null,
          createdAt: new Date(Date.now() - 1000 * 60 * 30).toISOString(),
          status: "DELIVERED"
        };

        return {
          ...current,
          [initialData.activeConversationId]: sortMessagesBySequence([olderMessage, ...messages])
        };
      });
      setIsLoadingOlder(false);
    }, 500);
  }

  if (!activeConversation) {
    return null;
  }

  const activeMembers = activeConversation.members
    .map((member) => initialData.users.find((user) => user.id === member.userId))
    .filter((user): user is User => Boolean(user));
  const firstPeer = activeMembers.find((user) => user.id !== currentUser.id) ?? activeMembers[0];
  const isDisconnected = socket.status === "offline" || socket.status === "reconnecting";

  return (
    <div className="chat-shell">
      <ConversationList
        activeConversationId={initialData.activeConversationId}
        conversations={initialData.conversations}
        users={initialData.users}
      />

      <section className="thread-panel" aria-label={`${activeConversation.title} conversation`}>
        <header className="thread-header">
          <button className="icon-button mobile-only" type="button" aria-label="Open conversations" onClick={() => router.push("/chat")}>
            <Menu size={18} />
          </button>
          <div className="thread-title">
            <p className="eyebrow">{activeConversation.systemTag}</p>
            <h1>{activeConversation.title}</h1>
            {firstPeer ? <PresencePill status={firstPeer.presence} detail={firstPeer.lastSeenAt ? "recent" : null} /> : null}
          </div>
          <div className="thread-actions">
            <span className="sequence-pill">next seq {nextSequence}</span>
            <button className="button button-secondary" type="button" onClick={socket.reconnect}>
              {socket.status === "reconnecting" ? <Loader2 className="spin" size={15} /> : <RotateCcw size={15} />}
              Reconnect
            </button>
            <button className="icon-button" type="button" aria-label="Go offline" onClick={socket.goOffline}>
              <CircleOff size={17} />
            </button>
          </div>
        </header>

        <div className="socket-banner" data-status={socket.status}>
          <Radio size={15} />
          <span>Socket {socket.status}</span>
          <strong>ACK reconciliation enabled</strong>
        </div>

        <MessageList
          messages={activeMessages}
          currentUserId={currentUser.id}
          users={initialData.users}
          isLoadingOlder={isLoadingOlder}
          onLoadOlder={loadOlderMessages}
        />

        <MessageComposer disabled={isDisconnected} onSend={sendMessage} />
      </section>

      <DetailsRail conversation={activeConversation} users={initialData.users} />
    </div>
  );
}
