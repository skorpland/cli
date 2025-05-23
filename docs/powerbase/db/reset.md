## powerbase-db-reset

Resets the local database to a clean state.

Requires the local development stack to be started by running `powerbase start`.

Recreates the local Postgres container and applies all local migrations found in `powerbase/migrations` directory. If test data is defined in `powerbase/seed.sql`, it will be seeded after the migrations are run. Any other data or schema changes made during local development will be discarded.

When running db reset with `--linked` or `--db-url` flag, a SQL script is executed to identify and drop all user created entities in the remote database. Since Postgres roles are cluster level entities, any custom roles created through the dashboard or `powerbase/roles.sql` will not be deleted by remote reset.
