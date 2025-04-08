## powerbase-stop

Stops the Powerbase local development stack.

Requires `powerbase/config.toml` to be created in your current working directory by running `powerbase init`.

All Docker resources are maintained across restarts.  Use `--no-backup` flag to reset your local development data between restarts.

Use the `--all` flag to stop all local Powerbase projects instances on the machine. Use with caution with `--no-backup` as it will delete all powerbase local projects data.