#!/usr/bin/env bash
set -euo pipefail

export PGHOST="$PGHOST"
export PGPORT="$PGPORT"
export PGUSER="$PGUSER"
export PGPASSWORD="$PGPASSWORD"
export PGDATABASE="$PGDATABASE"

# Explanation of pg_dump flags:
#
#   --schema-only     omit data like migration history, pgsodium key, etc.
#   --exclude-schema  omit internal schemas as they are maintained by platform
#
# Explanation of sed substitutions:
#
#   - do not alter superuser role "powerbase_admin"
#   - do not alter foreign data wrappers owner
#   - do not include ACL changes on internal schemas
#   - do not include RLS policies on cron extension schema
#   - do not include event triggers
#   - do not create pgtle schema and extension comments
#   - do not create publication "powerbase_realtime"
pg_dump \
    --schema-only \
    --quote-all-identifier \
    --exclude-schema "${EXCLUDED_SCHEMAS:-}" \
    ${EXTRA_FLAGS:-} \
| sed -E 's/^CREATE SCHEMA "/CREATE SCHEMA IF NOT EXISTS "/' \
| sed -E 's/^CREATE TABLE "/CREATE TABLE IF NOT EXISTS "/' \
| sed -E 's/^CREATE SEQUENCE "/CREATE SEQUENCE IF NOT EXISTS "/' \
| sed -E 's/^CREATE VIEW "/CREATE OR REPLACE VIEW "/' \
| sed -E 's/^CREATE FUNCTION "/CREATE OR REPLACE FUNCTION "/' \
| sed -E 's/^CREATE TRIGGER "/CREATE OR REPLACE TRIGGER "/' \
| sed -E 's/^CREATE PUBLICATION "powerbase_realtime/-- &/' \
| sed -E 's/^CREATE EVENT TRIGGER /-- &/' \
| sed -E 's/^         WHEN TAG IN /-- &/' \
| sed -E 's/^   EXECUTE FUNCTION /-- &/' \
| sed -E 's/^ALTER EVENT TRIGGER /-- &/' \
| sed -E 's/^ALTER PUBLICATION "powerbase_realtime_/-- &/' \
| sed -E 's/^ALTER FOREIGN DATA WRAPPER (.+) OWNER TO /-- &/' \
| sed -E 's/^ALTER DEFAULT PRIVILEGES FOR ROLE "powerbase_admin"/-- &/' \
| sed -E "s/^GRANT (.+) ON (.+) \"(${EXCLUDED_SCHEMAS:-})\"/-- &/" \
| sed -E "s/^REVOKE (.+) ON (.+) \"(${EXCLUDED_SCHEMAS:-})\"/-- &/" \
| sed -E 's/^(CREATE EXTENSION IF NOT EXISTS "pg_tle").+/\1;/' \
| sed -E 's/^(CREATE EXTENSION IF NOT EXISTS "pgsodium").+/\1;/' \
| sed -E 's/^COMMENT ON EXTENSION (.+)/-- &/' \
| sed -E 's/^CREATE POLICY "cron_job_/-- &/' \
| sed -E 's/^ALTER TABLE "cron"/-- &/' \
| sed -E "${EXTRA_SED:-}"

# Reset session config generated by pg_dump
echo "RESET ALL;"
