"use client";

import clsx from "clsx";
import Link from "next/link";
import { Search } from "lucide-react";
import { useMemo, useState } from "react";

import { PresencePill } from "@/components/chat/presence-pill";
import type { Conversation, User } from "@/lib/types";

type ConversationListProps = {
  activeConversationId: string;
  conversations: Conversation[];
  users: User[];
};

function getConversationPresence(conversation: Conversation, users: User[]) {
  const member = conversation.members
    .map((conversationMember) => users.find((user) => user.id === conversationMember.userId))
    .find((user) => user?.presence === "ONLINE" || user?.presence === "LAST_SEEN");

  return member?.presence ?? "UNKNOWN";
}

export function ConversationList({ activeConversationId, conversations, users }: ConversationListProps) {
  const [query, setQuery] = useState("");

  const filteredConversations = useMemo(() => {
    const normalized = query.trim().toLowerCase();
    if (!normalized) {
      return conversations;
    }

    return conversations.filter((conversation) => {
      return `${conversation.title} ${conversation.description} ${conversation.systemTag}`
        .toLowerCase()
        .includes(normalized);
    });
  }, [conversations, query]);

  return (
    <aside className="conversation-pane" aria-label="Conversations">
      <div className="pane-header">
        <div>
          <p className="eyebrow">Workspace</p>
          <h2>Messages</h2>
        </div>
        <span className="count-pill">{conversations.length}</span>
      </div>

      <label className="conversation-search">
        <Search size={15} aria-hidden="true" />
        <span className="sr-only">Filter conversations</span>
        <input value={query} onChange={(event) => setQuery(event.target.value)} placeholder="Filter threads" />
      </label>

      <div className="conversation-list">
        {filteredConversations.map((conversation) => (
          <Link
            className={clsx("conversation-row", activeConversationId === conversation.id && "is-active")}
            href={`/chat/${conversation.id}`}
            key={conversation.id}
          >
            <div className="conversation-avatar" data-kind={conversation.type.toLowerCase()}>
              {conversation.type === "GROUP" ? "GR" : conversation.title.slice(0, 2).toUpperCase()}
            </div>
            <div className="conversation-copy">
              <div className="conversation-title-row">
                <strong>{conversation.title}</strong>
                <span>{conversation.lastMessageAt}</span>
              </div>
              <p>{conversation.lastMessagePreview}</p>
              <div className="conversation-meta">
                <PresencePill status={getConversationPresence(conversation, users)} />
                <span>{conversation.systemTag}</span>
              </div>
            </div>
            {conversation.unreadCount > 0 ? <span className="unread-badge">{conversation.unreadCount}</span> : null}
          </Link>
        ))}
      </div>
    </aside>
  );
}
