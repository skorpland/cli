## powerbase-login

Connect the Powerbase CLI to your Powerbase account by logging in with your [personal access token](https://powerbase.club/dashboard/account/tokens).

Your access token is stored securely in [native credentials storage](https://github.com/zalando/go-keyring#dependencies). If native credentials storage is unavailable, it will be written to a plain text file at `~/.powerbase/access-token`.

> If this behavior is not desired, such as in a CI environment, you may skip login by specifying the `POWERBASE_ACCESS_TOKEN` environment variable in other commands.

The Powerbase CLI uses the stored token to access Management APIs for projects, functions, secrets, etc.
