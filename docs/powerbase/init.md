## powerbase-init

Initialize configurations for Powerbase local development.

A `powerbase/config.toml` file is created in your current working directory. This configuration is specific to each local project.

> You may override the directory path by specifying the `POWERBASE_WORKDIR` environment variable or `--workdir` flag.

In addition to `config.toml`, the `powerbase` directory may also contain other Powerbase objects, such as `migrations`, `functions`, `tests`, etc.
