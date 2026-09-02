#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=infra/scripts/db-env.sh
source "${script_dir}/db-env.sh"

psql_cmd="$(require_psql)"
dsn="$(database_url)"
migrations_dir="${repo_root}/infra/postgres/migrations"

if [[ ! -d "${migrations_dir}" ]]; then
  printf 'Migration directory not found: %s\n' "${migrations_dir}" >&2
  exit 1
fi

for migration in "${migrations_dir}"/*.sql; do
  [[ -e "${migration}" ]] || continue
  printf 'Applying migration: %s\n' "${migration#${repo_root}/}"
  "${psql_cmd}" "${dsn}" -v ON_ERROR_STOP=1 -f "${migration}"
done

printf 'Migrations complete.\n'
