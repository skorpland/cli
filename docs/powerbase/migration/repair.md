## powerbase-migration-repair

Repairs the remote migration history table.

Requires your local project to be linked to a remote database by running `powerbase link`.

If your local and remote migration history goes out of sync, you can repair the remote history by marking specific migrations as `--status applied` or `--status reverted`. Marking as `reverted` will delete an existing record from the migration history table while marking as `applied` will insert a new record.

For example, your migration history may look like the table below, with missing entries in either local or remote.

```bash
$ powerbase migration list
        LOCAL      │     REMOTE     │     TIME (UTC)
  ─────────────────┼────────────────┼──────────────────────
                   │ 20230103054303 │ 2023-01-03 05:43:03
   20230103054315  │                │ 2023-01-03 05:43:15
```

To reset your migration history to a clean state, first delete your local migration file.

```bash
$ rm powerbase/migrations/20230103054315_remote_commit.sql

$ powerbase migration list
        LOCAL      │     REMOTE     │     TIME (UTC)
  ─────────────────┼────────────────┼──────────────────────
                   │ 20230103054303 │ 2023-01-03 05:43:03
```

Then mark the remote migration `20230103054303` as reverted.

```bash
$ powerbase migration repair 20230103054303 --status reverted
Connecting to remote database...
Repaired migration history: [20220810154537] => reverted
Finished powerbase migration repair.

$ powerbase migration list
        LOCAL      │     REMOTE     │     TIME (UTC)
  ─────────────────┼────────────────┼──────────────────────
```

Now you can run `db pull` again to dump the remote schema as a local migration file.

```bash
$ powerbase db pull
Connecting to remote database...
Schema written to powerbase/migrations/20240414044403_remote_schema.sql
Update remote migration history table? [Y/n]
Repaired migration history: [20240414044403] => applied
Finished powerbase db pull.

$ powerbase migration list
        LOCAL      │     REMOTE     │     TIME (UTC)
  ─────────────────┼────────────────┼──────────────────────
    20240414044403 │ 20240414044403 │ 2024-04-14 04:44:03
```
