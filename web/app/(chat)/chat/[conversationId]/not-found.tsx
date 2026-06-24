import Link from "next/link";

export default function ConversationNotFound() {
  return (
    <div className="route-error">
      <h2>Conversation not found</h2>
      <p>This local conversation id is not part of the seeded workspace.</p>
      <Link className="button button-primary" href="/chat">
        Back to chat
      </Link>
    </div>
  );
}
