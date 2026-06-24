import { Database, Radio, ShieldCheck, Users } from "lucide-react";

import { PresencePill } from "@/components/chat/presence-pill";
import type { Conversation, User } from "@/lib/types";

type DetailsRailProps = {
  conversation: Conversation;
  users: User[];
};

export function DetailsRail({ conversation, users }: DetailsRailProps) {
  const members = conversation.members
    .map((member) => users.find((user) => user.id === member.userId))
    .filter((user): user is User => Boolean(user));

  return (
    <aside className="details-rail" aria-label="Conversation details">
      <section className="rail-card">
        <div className="rail-card-heading">
          <Users size={16} />
          <h3>Participants</h3>
        </div>
        <div className="member-stack">
          {members.map((user) => (
            <div className="member-row" key={user.id}>
              <span className={`avatar avatar-${user.avatarColor}`}>{user.initials}</span>
              <span>
                <strong>{user.displayName}</strong>
                <small>{user.role}</small>
              </span>
              <PresencePill status={user.presence} detail={user.lastSeenAt ? "recent" : null} />
            </div>
          ))}
        </div>
      </section>

      <section className="rail-card">
        <div className="rail-card-heading">
          <Database size={16} />
          <h3>Storage contract</h3>
        </div>
        <dl className="contract-list">
          <div>
            <dt>History</dt>
            <dd>PostgreSQL messages</dd>
          </div>
          <div>
            <dt>Ordering</dt>
            <dd>sequence ASC</dd>
          </div>
          <div>
            <dt>Dedup</dt>
            <dd>client_msg_id</dd>
          </div>
        </dl>
      </section>

      <section className="rail-card">
        <div className="rail-card-heading">
          <Radio size={16} />
          <h3>Realtime state</h3>
        </div>
        <div className="health-row">
          <span>WS push</span>
          <strong>Ready</strong>
        </div>
        <div className="health-row">
          <span>Redis presence</span>
          <strong>60s TTL</strong>
        </div>
        <div className="health-row">
          <span>Kafka fan-out</span>
          <strong>Mocked</strong>
        </div>
      </section>

      <section className="rail-card subdued">
        <div className="rail-card-heading">
          <ShieldCheck size={16} />
          <h3>Local note</h3>
        </div>
        <p>Backend APIs are not wired yet. This UI uses mock data with the same contracts planned for V1.</p>
      </section>
    </aside>
  );
}
