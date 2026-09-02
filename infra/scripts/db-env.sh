#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
repo_root="$(cd "${script_dir}/../.." && pwd)"

strip_quotes() {
  local value="$1"

  case "${value}" in
    \"*\")
      value="${value#\"}"
      value="${value%\"}"
      ;;
    \'*\')
      value="${value#\'}"
      value="${value%\'}"
      ;;
  esac

  printf '%s' "${value}"
}

read_env_value() {
  local key="$1"
  local file
  local line
  local value

  for file in "${repo_root}/.env" "${repo_root}/.env.local"; do
    if [[ ! -f "${file}" ]]; then
      continue
    fi

    line="$(grep -E "^${key}=" "${file}" | tail -n 1 || true)"
    if [[ -z "${line}" ]]; then
      continue
    fi

    value="${line#*=}"
    strip_quotes "${value}"
    return 0
  done

  return 1
}

database_url() {
  if [[ -n "${POSTGRES_DSN:-}" ]]; then
    printf '%s' "${POSTGRES_DSN}"
    return 0
  fi

  if [[ -n "${DATABASE_URL:-}" ]]; then
    printf '%s' "${DATABASE_URL}"
    return 0
  fi

  if read_env_value "POSTGRES_DSN"; then
    return 0
  fi

  if read_env_value "DATABASE_URL"; then
    return 0
  fi

  printf 'POSTGRES_DSN or DATABASE_URL is required. Put your Neon connection string in .env.\n' >&2
  return 1
}

psql_bin() {
  local candidate

  if command -v psql >/dev/null 2>&1; then
    command -v psql
    return 0
  fi

  for candidate in \
    "/opt/homebrew/opt/libpq/bin/psql" \
    "/usr/local/opt/libpq/bin/psql" \
    "/opt/homebrew/opt/postgresql@16/bin/psql" \
    "/usr/local/opt/postgresql@16/bin/psql"; do
    if [[ -x "${candidate}" ]]; then
      printf '%s\n' "${candidate}"
      return 0
    fi
  done

  return 1
}

require_psql() {
  if ! psql_bin; then
    cat >&2 <<'EOF'
psql is required to run database scripts.

Install PostgreSQL client tools, then retry:
  macOS: brew install libpq
  or:    brew install postgresql@16

Neon usage:
  POSTGRES_DSN="postgresql://role:password@ep-...neon.tech/dbname?sslmode=require&channel_binding=require"
EOF
    return 1
  fi
}
