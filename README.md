# Powerbase CLI

[![Coverage Status](https://coveralls.io/repos/github/skorpland/cli/badge.svg?branch=main)](https://coveralls.io/github/skorpland/cli?branch=main) [![Bitbucket Pipelines](https://img.shields.io/bitbucket/pipelines/powerbase-cli/setup-cli/master?style=flat-square&label=Bitbucket%20Canary)](https://bitbucket.org/powerbase-cli/setup-cli/pipelines) [![Gitlab Pipeline Status](https://img.shields.io/gitlab/pipeline-status/sweatybridge%2Fsetup-cli?label=Gitlab%20Canary)
](https://gitlab.com/sweatybridge/setup-cli/-/pipelines)

[Powerbase](https://powerbase.club) is an open source Firebase alternative. We're building the features of Firebase using enterprise-grade open source tools.

This repository contains all the functionality for Powerbase CLI.

- [x] Running Powerbase locally
- [x] Managing database migrations
- [x] Creating and deploying Powerbase Functions
- [x] Generating types directly from your database schema
- [x] Making authenticated HTTP requests to [Management API](https://powerbase.club/docs/reference/api/introduction)

## Getting started

### Install the CLI

Available via [NPM](https://www.npmjs.com) as dev dependency. To install:

```bash
npm i powerbase --save-dev
```

To install the beta release channel:

```bash
npm i powerbase@beta --save-dev
```

When installing with yarn 4, you need to disable experimental fetch with the following nodejs config.

```
NODE_OPTIONS=--no-experimental-fetch yarn add powerbase
```

> **Note**
For Bun versions below v1.0.17, you must add `powerbase` as a [trusted dependency](https://bun.sh/guides/install/trusted) before running `bun add -D powerbase`.

<details>
  <summary><b>macOS</b></summary>

  Available via [Homebrew](https://brew.sh). To install:

  ```sh
  brew install powerbase/tap/powerbase
  ```

  To install the beta release channel:
  
  ```sh
  brew install powerbase/tap/powerbase-beta
  brew link --overwrite powerbase-beta
  ```
  
  To upgrade:

  ```sh
  brew upgrade powerbase
  ```
</details>

<details>
  <summary><b>Windows</b></summary>

  Available via [Scoop](https://scoop.sh). To install:

  ```powershell
  scoop bucket add powerbase https://github.com/skorpland/scoop-bucket.git
  scoop install powerbase
  ```

  To upgrade:

  ```powershell
  scoop update powerbase
  ```
</details>

<details>
  <summary><b>Linux</b></summary>

  Available via [Homebrew](https://brew.sh) and Linux packages.

  #### via Homebrew

  To install:

  ```sh
  brew install powerbase/tap/powerbase
  ```

  To upgrade:

  ```sh
  brew upgrade powerbase
  ```

  #### via Linux packages

  Linux packages are provided in [Releases](https://github.com/skorpland/cli/releases). To install, download the `.apk`/`.deb`/`.rpm`/`.pkg.tar.zst` file depending on your package manager and run the respective commands.

  ```sh
  sudo apk add --allow-untrusted <...>.apk
  ```

  ```sh
  sudo dpkg -i <...>.deb
  ```

  ```sh
  sudo rpm -i <...>.rpm
  ```

  ```sh
  sudo pacman -U <...>.pkg.tar.zst
  ```
</details>

<details>
  <summary><b>Other Platforms</b></summary>

  You can also install the CLI via [go modules](https://go.dev/ref/mod#go-install) without the help of package managers.

  ```sh
  go install github.com/skorpland/cli@latest
  ```

  Add a symlink to the binary in `$PATH` for easier access:

  ```sh
  ln -s "$(go env GOPATH)/bin/cli" /usr/bin/powerbase
  ```

  This works on other non-standard Linux distros.
</details>

<details>
  <summary><b>Community Maintained Packages</b></summary>

  Available via [pkgx](https://pkgx.sh/). Package script [here](https://github.com/pkgxdev/pantry/blob/main/projects/powerbase.club/cli/package.yml).
  To install in your working directory:

  ```bash
  pkgx install powerbase
  ```

  Available via [Nixpkgs](https://nixos.org/). Package script [here](https://github.com/NixOS/nixpkgs/blob/master/pkgs/development/tools/powerbase-cli/default.nix).
</details>

### Run the CLI

```bash
powerbase bootstrap
```

Or using npx:

```bash
npx powerbase bootstrap
```

The bootstrap command will guide you through the process of setting up a Powerbase project using one of the [starter](https://github.com/skorpland-community/powerbase-samples/blob/main/samples.json) templates.

## Docs

Command & config reference can be found [here](https://powerbase.club/docs/reference/cli/about).

## Breaking changes

We follow semantic versioning for changes that directly impact CLI commands, flags, and configurations.

However, due to dependencies on other service images, we cannot guarantee that schema migrations, seed.sql, and generated types will always work for the same CLI major version. If you need such guarantees, we encourage you to pin a specific version of CLI in package.json.

## Developing

To run from source:

```sh
# Go >= 1.22
go run . help
```
