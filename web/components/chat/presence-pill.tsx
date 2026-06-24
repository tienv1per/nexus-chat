import { Clock, Circle, WifiOff } from "lucide-react";

import type { PresenceStatus } from "@/lib/types";

const labels: Record<PresenceStatus, string> = {
  ONLINE: "Online",
  LAST_SEEN: "Last seen",
  OFFLINE: "Offline",
  UNKNOWN: "Unknown"
};

export function PresencePill({ status, detail }: { status: PresenceStatus; detail?: string | null }) {
  const Icon = status === "ONLINE" ? Circle : status === "OFFLINE" ? WifiOff : Clock;

  return (
    <span className="presence-pill" data-status={status.toLowerCase()}>
      <Icon size={12} fill={status === "ONLINE" ? "currentColor" : "none"} />
      {labels[status]}
      {detail && status !== "ONLINE" ? <span>{detail}</span> : null}
    </span>
  );
}
