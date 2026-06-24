"use client";

export default function LoginError({ reset }: { reset: () => void }) {
  return (
    <div className="route-error">
      <h2>Login could not load</h2>
      <p>Seeded users are unavailable. Try refreshing the local mock state.</p>
      <button className="button button-primary" onClick={reset}>
        Try again
      </button>
    </div>
  );
}
