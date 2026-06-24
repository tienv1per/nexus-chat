import { ShieldCheck, UserRoundCheck, Users } from "lucide-react";

export default function AdminPage() {
  return (
    <section className="placeholder-page">
      <div className="placeholder-heading">
        <p className="eyebrow">Admin</p>
        <h1>Seeded workspace controls</h1>
        <p>Member tables, reset actions, and local diagnostics stay queued behind the main chat UI.</p>
      </div>
      <div className="placeholder-grid">
        <article className="placeholder-card">
          <Users size={18} />
          <span>Seed users</span>
          <strong>Alice, Bob, Charlie</strong>
        </article>
        <article className="placeholder-card">
          <UserRoundCheck size={18} />
          <span>Roles</span>
          <strong>Owner, Member, Test</strong>
        </article>
        <article className="placeholder-card">
          <ShieldCheck size={18} />
          <span>Controls</span>
          <strong>Phase 2 preview</strong>
        </article>
      </div>
    </section>
  );
}
