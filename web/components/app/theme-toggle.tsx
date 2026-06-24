"use client";

import { Moon, Sun } from "lucide-react";
import { useEffect, useState } from "react";

type Theme = "light" | "dark";

function readTheme(): Theme {
  if (typeof window === "undefined") {
    return "dark";
  }

  const saved = window.localStorage.getItem("nexuschat-theme");
  if (saved === "light" || saved === "dark") {
    return saved;
  }

  return window.matchMedia("(prefers-color-scheme: dark)").matches ? "dark" : "light";
}

export function ThemeToggle() {
  const [theme, setTheme] = useState<Theme>("dark");

  useEffect(() => {
    const current = readTheme();
    setTheme(current);
    document.documentElement.dataset.theme = current;
  }, []);

  function toggleTheme() {
    const next = theme === "dark" ? "light" : "dark";
    setTheme(next);
    window.localStorage.setItem("nexuschat-theme", next);
    document.documentElement.dataset.theme = next;
  }

  const Icon = theme === "dark" ? Moon : Sun;

  return (
    <button className="icon-button theme-toggle" type="button" aria-label="Toggle color theme" onClick={toggleTheme}>
      <Icon size={17} strokeWidth={1.9} />
    </button>
  );
}
