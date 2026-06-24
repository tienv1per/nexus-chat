"use client";

export default function ChatError({ reset }: { reset: () => void }) {
  return (
    <div className="route-error">
      <h2>Conversation failed to load</h2>
      <p>The chat workspace hit a local rendering error.</p>
      <button className="button button-primary" onClick={reset}>
        Reload workspace
      </button>
    </div>
  );
}
