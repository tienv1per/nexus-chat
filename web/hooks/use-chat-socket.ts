"use client";

import { useCallback, useEffect, useRef, useState } from "react";

import { createAck, createRecipientEcho, type SocketStatus } from "@/lib/socket";
import type { Message, MessageAck, SendDraft } from "@/lib/types";

type UseChatSocketInput = {
  conversationId: string;
  onAck: (ack: MessageAck) => void;
  onIncoming: (message: Message) => void;
};

export function useChatSocket({ conversationId, onAck, onIncoming }: UseChatSocketInput) {
  const [status, setStatus] = useState<SocketStatus>("connecting");
  const timers = useRef<number[]>([]);

  const clearTimers = useCallback(() => {
    timers.current.forEach((timer) => window.clearTimeout(timer));
    timers.current = [];
  }, []);

  useEffect(() => {
    setStatus("connecting");
    const connectTimer = window.setTimeout(() => setStatus("connected"), 420);
    timers.current.push(connectTimer);

    return () => {
      clearTimers();
    };
  }, [clearTimers, conversationId]);

  const sendMessage = useCallback(
    (draft: SendDraft) => {
      const ackTimer = window.setTimeout(() => {
        onAck(createAck(draft));
      }, 640);

      const echoTimer = window.setTimeout(() => {
        if (draft.body.trim().length > 0) {
          onIncoming(createRecipientEcho(draft));
        }
      }, 1800);

      timers.current.push(ackTimer, echoTimer);
    },
    [onAck, onIncoming]
  );

  const reconnect = useCallback(() => {
    setStatus("reconnecting");
    const reconnectTimer = window.setTimeout(() => setStatus("connected"), 700);
    timers.current.push(reconnectTimer);
  }, []);

  const goOffline = useCallback(() => {
    setStatus("offline");
  }, []);

  return {
    status,
    sendMessage,
    reconnect,
    goOffline
  };
}
