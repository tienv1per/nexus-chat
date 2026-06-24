import type { Metadata, Viewport } from "next";
import Image from "next/image";
import Link from "next/link";
import { Activity, MessageSquare, Search, Settings, ShieldCheck } from "lucide-react";

import { ThemeToggle } from "@/components/app/theme-toggle";
import "./globals.css";

export const metadata: Metadata = {
  title: "NexusChat",
  description: "A local-first real-time chat workspace for learning production messaging systems."
};

export const viewport: Viewport = {
  width: "device-width",
  initialScale: 1,
  colorScheme: "light dark"
};

const themeScript = `
(() => {
  try {
    const saved = localStorage.getItem("nexuschat-theme");
    const systemDark = window.matchMedia("(prefers-color-scheme: dark)").matches;
    document.documentElement.dataset.theme = saved || (systemDark ? "dark" : "light");
  } catch {
    document.documentElement.dataset.theme = "dark";
  }
})();
`;

const navItems = [
  { href: "/chat", label: "Messages", icon: MessageSquare },
  { href: "/dashboard", label: "Activity", icon: Activity },
  { href: "/admin", label: "Admin", icon: ShieldCheck },
  { href: "/settings", label: "Settings", icon: Settings }
];

export default function RootLayout({ children }: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="en" suppressHydrationWarning>
      <body>
        <script dangerouslySetInnerHTML={{ __html: themeScript }} />
        <div className="app-chrome">
          <aside className="global-sidebar" aria-label="Primary navigation">
            <Link href="/chat" className="brand-mark" aria-label="NexusChat home">
              <Image src="/brand/nexuschat-logo-mark.svg" alt="" width={32} height={32} priority className="brand-image" />
            </Link>
            <nav className="global-nav" aria-label="Workspace modules">
              {navItems.map((item) => {
                const Icon = item.icon;
                return (
                  <Link key={item.href} href={item.href} className="icon-button" title={item.label} aria-label={item.label}>
                    <Icon size={18} strokeWidth={1.9} />
                  </Link>
                );
              })}
            </nav>
          </aside>

          <div className="workspace">
            <header className="workspace-topbar">
              <Link href="/chat" className="wordmark" aria-label="NexusChat chat">
                <Image src="/brand/nexuschat-logo.svg" alt="NexusChat" width={142} height={40} priority className="logo-image" />
              </Link>
              <label className="command-search">
                <Search size={16} aria-hidden="true" />
                <span className="sr-only">Search conversations</span>
                <input placeholder="Search messages, users, sequences" />
              </label>
              <div className="topbar-actions">
                <span className="system-pill">
                  <span className="pulse-dot" />
                  Local stack
                </span>
                <ThemeToggle />
              </div>
            </header>
            <main className="workspace-main">{children}</main>
          </div>
        </div>
      </body>
    </html>
  );
}
