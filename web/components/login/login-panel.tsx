"use client";

import { ArrowRight, Laptop, Loader2 } from "lucide-react";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

import { loginSeedUser } from "@/lib/api";
import type { User } from "@/lib/types";

function getOrCreateDeviceId() {
  const existing = window.localStorage.getItem("nexuschat-device-id");
  if (existing) {
    return existing;
  }

  const next = `web-${Math.random().toString(36).slice(2, 10)}`;
  window.localStorage.setItem("nexuschat-device-id", next);
  return next;
}

export function LoginPanel({ users }: { users: User[] }) {
  const router = useRouter();
  const [selectedUserId, setSelectedUserId] = useState(users[0]?.id ?? "");
  const [deviceId, setDeviceId] = useState("web-local");
  const [error, setError] = useState("");
  const [isSubmitting, setIsSubmitting] = useState(false);

  useEffect(() => {
    setDeviceId(getOrCreateDeviceId());
  }, []);

  async function submitLogin() {
    setError("");
    setIsSubmitting(true);

    try {
      const response = await loginSeedUser(selectedUserId, deviceId);
      window.localStorage.setItem("nexuschat-user-id", response.data.user.id);
      document.cookie = `chat_token=${response.data.token}; path=/; max-age=86400; samesite=lax`;
      router.push("/chat");
    } catch (loginError) {
      setError(loginError instanceof Error ? loginError.message : "Unable to login with this seed user.");
    } finally {
      setIsSubmitting(false);
    }
  }

  return (
    <div className="login-panel">
      <div className="panel-heading">
        <p className="eyebrow">Seed users</p>
        <h2>Start a two-browser demo</h2>
      </div>

      <div className="seed-user-list" role="radiogroup" aria-label="Seed users">
        {users.map((user) => (
          <button
            className="seed-user"
            data-selected={selectedUserId === user.id}
            key={user.id}
            type="button"
            role="radio"
            aria-checked={selectedUserId === user.id}
            onClick={() => setSelectedUserId(user.id)}
          >
            <span className={`avatar avatar-${user.avatarColor}`}>{user.initials}</span>
            <span>
              <strong>{user.displayName}</strong>
              <small>{user.email}</small>
            </span>
            <span className="role-badge">{user.role}</span>
          </button>
        ))}
      </div>

      <label className="field-label" htmlFor="device-id">
        Device ID
      </label>
      <div className="device-field">
        <Laptop size={16} aria-hidden="true" />
        <input id="device-id" value={deviceId} onChange={(event) => setDeviceId(event.target.value)} />
      </div>

      {error ? <p className="inline-error">{error}</p> : null}

      <button className="button button-primary login-submit" type="button" disabled={isSubmitting} onClick={submitLogin}>
        {isSubmitting ? <Loader2 className="spin" size={16} /> : <ArrowRight size={16} />}
        Continue to chat
      </button>
    </div>
  );
}
