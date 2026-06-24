import { MonitorCog, Moon, Radio } from "lucide-react";

export default function SettingsPage() {
  return (
    <section className="placeholder-page">
      <div className="placeholder-heading">
        <p className="eyebrow">Settings</p>
        <h1>Workspace preferences</h1>
        <p>Theme persistence and mock device identity are active. Advanced settings can land after backend contracts settle.</p>
      </div>
      <div className="placeholder-grid">
        <article className="placeholder-card">
          <Moon size={18} />
          <span>Theme</span>
          <strong>Light / dark</strong>
        </article>
        <article className="placeholder-card">
          <MonitorCog size={18} />
          <span>Device</span>
          <strong>Browser local</strong>
        </article>
        <article className="placeholder-card">
          <Radio size={18} />
          <span>Realtime</span>
          <strong>Mock socket</strong>
        </article>
      </div>
    </section>
  );
}
