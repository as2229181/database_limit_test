#!/bin/bash
set -euo pipefail

: "${RECORD_PASSWORD:?RECORD_PASSWORD is required}"

export record_password="${RECORD_PASSWORD}"

envsubst '${record_password}' \
	< "$(dirname "$0")/10-record.sql.template" \
	| psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB"
