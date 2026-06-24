import type { Metadata } from "next";
import Image from "next/image";

import { LoginPanel } from "@/components/login/login-panel";
import { getSeedUsers } from "@/lib/api";

export const metadata: Metadata = {
  title: "Dev Login | NexusChat"
};

export default async function LoginPage() {
  const users = await getSeedUsers();

  return (
    <section className="login-page">
      <div className="login-hero">
        <Image src="/brand/nexuschat-logo.svg" alt="NexusChat" width={192} height={54} priority className="login-logo" />
        <div>
          <p className="eyebrow">Local development</p>
          <h1>Choose a seeded user and jump into the workspace.</h1>
          <p className="section-copy">
            The browser stores a local device id and mock chat token now. When backend auth arrives, this panel can switch
            from mock actions to the real `/v1/dev/login` contract.
          </p>
        </div>
        <div className="login-stack">
          <span>Cookie: `chat_token`</span>
          <span>Device: local storage</span>
          <span>Transport: WebSocket-ready</span>
        </div>
      </div>
      <LoginPanel users={users} />
    </section>
  );
}
