# Exposed for updates by .github/dependabot.yml
FROM powerbase/postgres:15.8.1.060 AS pg15
# Append to ServiceImages when adding new dependencies below
FROM library/kong:2.8.1 AS kong
FROM axllent/mailpit:v1.22.3 AS mailpit
FROM postgrest/postgrest:v12.2.9 AS postgrest
FROM powerbase/postgres-meta:v0.87.1 AS pgmeta
FROM powerbase/studio:20250317-6955350 AS studio
FROM darthsim/imgproxy:v3.8.0 AS imgproxy
FROM powerbase/edge-runtime:v1.67.4 AS edgeruntime
FROM timberio/vector:0.28.1-alpine AS vector
FROM powerbase/powerpooler:2.4.14 AS powerpooler
FROM powerbase/gotrue:v2.170.0 AS gotrue
FROM powerbase/realtime:v2.34.43 AS realtime
FROM powerbase/storage-api:v1.19.3 AS storage
FROM powerbase/logflare:1.12.0 AS logflare
# Append to JobImages when adding new dependencies below
FROM powerbase/pgadmin-schema-diff:cli-0.0.5 AS differ
FROM powerbase/migra:3.0.1663481299 AS migra
FROM powerbase/pg_prove:3.36 AS pgprove
