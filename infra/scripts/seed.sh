#!/usr/bin/env bash
set -euo pipefail

script_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
# shellcheck source=infra/scripts/db-env.sh
source "${script_dir}/db-env.sh"

psql_cmd="$(require_psql)"
dsn="$(database_url)"
seeds_dir="${repo_root}/infra/postgres/seeds"

if [[ ! -d "${seeds_dir}" ]]; then
  printf 'Seed directory not found: %s\n' "${seeds_dir}" >&2
  exit 1
fi

for seed in "${seeds_dir}"/*.sql; do
  [[ -e "${seed}" ]] || continue
  printf 'Applying seed: %s\n' "${seed#${repo_root}/}"
  "${psql_cmd}" "${dsn}" -v ON_ERROR_STOP=1 -f "${seed}"
done

printf 'Seed data complete.\n'
