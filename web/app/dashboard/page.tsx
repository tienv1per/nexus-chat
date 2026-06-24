import { Activity, Database, Radio, Server } from "lucide-react";

const cards = [
  { label: "Chat Service", value: "Mock warmup", icon: Server },
  { label: "WebSocket", value: "Connected UI", icon: Radio },
  { label: "Postgres", value: "Planned Phase 3", icon: Database },
  { label: "Events", value: "Local preview", icon: Activity }
];

export default function DashboardPage() {
  return (
    <section className="placeholder-page">
      <div className="placeholder-heading">
        <p className="eyebrow">Operations</p>
        <h1>Local service dashboard</h1>
        <p>Phase 2 keeps this as a lightweight preview while the chat workspace gets implemented first.</p>
      </div>
      <div className="placeholder-grid">
        {cards.map((card) => {
          const Icon = card.icon;
          return (
            <article className="placeholder-card" key={card.label}>
              <Icon size={18} />
              <span>{card.label}</span>
              <strong>{card.value}</strong>
            </article>
          );
        })}
      </div>
    </section>
  );
}
