"use client";

export default function ConversationError({ reset }: { reset: () => void }) {
  return (
    <div className="route-error">
      <h2>Conversation unavailable</h2>
      <p>Mock history could not be normalized for this conversation.</p>
      <button className="button button-primary" onClick={reset}>
        Retry
      </button>
    </div>
  );
}
