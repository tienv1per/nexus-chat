"use client";

import { Paperclip, SendHorizontal, X } from "lucide-react";
import { useRef, useState } from "react";

import type { MediaObject } from "@/lib/types";

type MessageComposerProps = {
  disabled: boolean;
  onSend: (body: string, media?: MediaObject | null) => void;
};

function createMedia(file: File): MediaObject {
  return {
    id: `media_${Date.now()}`,
    filename: file.name,
    mimeType: file.type || "application/octet-stream",
    sizeLabel: `${Math.max(1, Math.round(file.size / 1024))} KB`,
    progress: 100
  };
}

export function MessageComposer({ disabled, onSend }: MessageComposerProps) {
  const [body, setBody] = useState("");
  const [media, setMedia] = useState<MediaObject | null>(null);
  const fileInputRef = useRef<HTMLInputElement>(null);
  const canSend = !disabled && (body.trim().length > 0 || media !== null);

  function submit() {
    if (!canSend) {
      return;
    }

    onSend(body.trim(), media);
    setBody("");
    setMedia(null);
  }

  function handleKeyDown(event: React.KeyboardEvent<HTMLTextAreaElement>) {
    if (event.key === "Enter" && !event.shiftKey) {
      event.preventDefault();
      submit();
    }
  }

  return (
    <section className="composer" aria-label="Message composer">
      {media ? (
        <div className="upload-preview">
          <Paperclip size={16} />
          <span>
            <strong>{media.filename}</strong>
            <small>{media.mimeType} · {media.sizeLabel}</small>
          </span>
          <button className="icon-button" type="button" aria-label="Remove attachment" onClick={() => setMedia(null)}>
            <X size={15} />
          </button>
        </div>
      ) : null}

      <div className="composer-row">
        <input
          ref={fileInputRef}
          className="sr-only"
          type="file"
          onChange={(event) => {
            const file = event.target.files?.[0];
            if (file) {
              setMedia(createMedia(file));
            }
          }}
        />
        <button className="icon-button" type="button" aria-label="Attach file" onClick={() => fileInputRef.current?.click()}>
          <Paperclip size={18} />
        </button>
        <textarea
          value={body}
          rows={1}
          placeholder={disabled ? "Reconnect before sending" : "Message this conversation"}
          disabled={disabled}
          onChange={(event) => setBody(event.target.value)}
          onKeyDown={handleKeyDown}
        />
        <button className="send-button" type="button" disabled={!canSend} aria-label="Send message" onClick={submit}>
          <SendHorizontal size={18} />
        </button>
      </div>
    </section>
  );
}
