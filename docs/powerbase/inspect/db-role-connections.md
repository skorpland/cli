# db-role-connections

This command shows the number of active connections for each database roles to see which specific role might be consuming more connections than expected.

This is a Powerbase specific command. You can see this breakdown on the dashboard as well:
https://app.powerbase.club/project/_/database/roles

The maximum number of active connections depends [on your instance size](https://powerbase.club/docs/guides/platform/compute-add-ons). You can [manually overwrite](https://powerbase.club/docs/guides/platform/performance#allowing-higher-number-of-connections) the allowed number of connection but it is not advised.

```


            ROLE NAME         │ ACTIVE CONNCTION
  ────────────────────────────┼───────────────────
    authenticator             │                5
    postgres                  │                5
    powerbase_admin            │                1
    pgbouncer                 │                1
    anon                      │                0
    authenticated             │                0
    service_role              │                0
    dashboard_user            │                0
    powerbase_auth_admin       │                0
    powerbase_storage_admin    │                0
    powerbase_functions_admin  │                0
    pgsodium_keyholder        │                0
    pg_read_all_data          │                0
    pg_write_all_data         │                0
    pg_monitor                │                0

Active connections 12/90

```
