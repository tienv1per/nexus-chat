#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=infra/scripts/db-env.sh
source "${script_dir}/db-env.sh"

psql_cmd="$(require_psql)"
dsn="$(database_url)"
smoke_dir="${repo_root}/infra/postgres/smoke"

if [[ ! -d "${smoke_dir}" ]]; then
  printf 'Smoke test directory not found: %s\n' "${smoke_dir}" >&2
  exit 1
fi

for smoke in "${smoke_dir}"/*.sql; do
  [[ -e "${smoke}" ]] || continue
  printf 'Running database smoke test: %s\n' "${smoke#${repo_root}/}"
  "${psql_cmd}" "${dsn}" -v ON_ERROR_STOP=1 -f "${smoke}"
done

printf 'Database smoke tests complete.\n'
