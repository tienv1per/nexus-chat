"use client";

import { CheckCheck, Clock3, FileText, Loader2, RefreshCcw } from "lucide-react";

import { sortMessagesBySequence } from "@/lib/socket";
import type { Message, User } from "@/lib/types";

type MessageListProps = {
  messages: Message[];
  currentUserId: string;
  users: User[];
  onLoadOlder: () => void;
  isLoadingOlder: boolean;
};

function formatTime(value: string) {
  return new Intl.DateTimeFormat("en", {
    hour: "2-digit",
    minute: "2-digit"
  }).format(new Date(value));
}

function senderFor(users: User[], senderId: string) {
  return users.find((user) => user.id === senderId) ?? users[0];
}

export function MessageList({ messages, currentUserId, users, onLoadOlder, isLoadingOlder }: MessageListProps) {
  const sortedMessages = sortMessagesBySequence(messages);

  return (
    <div className="message-list" aria-live="polite">
      <button className="load-older" type="button" onClick={onLoadOlder} disabled={isLoadingOlder}>
        {isLoadingOlder ? <Loader2 className="spin" size={15} /> : <RefreshCcw size={15} />}
        Load earlier history
      </button>

      {sortedMessages.length === 0 ? (
        <div className="empty-thread">
          <h3>No messages yet</h3>
          <p>Start with a short note. The mock socket will ACK the send and insert by sequence.</p>
        </div>
      ) : null}

      {sortedMessages.map((message) => {
        const sender = senderFor(users, message.senderId);
        const isMine = message.senderId === currentUserId;

        return (
          <article className="message-row" data-own={isMine} key={message.messageId ?? message.clientMsgId}>
            <span className={`avatar avatar-${sender.avatarColor}`}>{sender.initials}</span>
            <div className="message-bubble">
              <div className="message-meta-row">
                <strong>{isMine ? "You" : sender.displayName}</strong>
                <span>seq {message.sequence}</span>
                <span>{formatTime(message.createdAt)}</span>
              </div>
              <p>{message.body}</p>
              {message.media ? (
                <div className="attachment-card">
                  <FileText size={17} />
                  <span>
                    <strong>{message.media.filename}</strong>
                    <small>{message.media.mimeType} · {message.media.sizeLabel}</small>
                  </span>
                </div>
              ) : null}
              <div className="delivery-row">
                {message.status === "PENDING" ? <Clock3 size={13} /> : <CheckCheck size={13} />}
                {message.status}
              </div>
            </div>
          </article>
        );
      })}
    </div>
  );
}
