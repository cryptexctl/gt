# gt – Gitea CLI

[![Go Version](https://img.shields.io/badge/go-1.21+-blue)](https://golang.org) [![License](https://img.shields.io/badge/license-0BSD-green)](LICENSE)

`gt` is a fast, minimal command-line tool to work with a self-hosted **Gitea** instance: repositories, issues, pull requests and more.

```bash
$ gt repo list
NAME                     DESCRIPTION            PRIVATE   FORK
cryptexctl/infra         Infrastructure code    yes       no
cryptexctl/teah          Gitea CLI              no        no
```

---

## Installation

### From source

```bash
git clone https://github.com/cryptexctl/teah.git
cd teah-cli
make build      # binary in bin/gt
sudo make install  # optional, copies to /usr/local/bin
```

### Pre-built binaries

Releases provide archives for Linux, macOS and Windows. Unpack and put `gt` (or `gt.exe`) in your `PATH`.

---

## Quick start

```bash
# Configure credentials (stores in ~/.config/gt/config.yaml)
gt auth login --host https://git.example.com --token YOUR_TOKEN

# List own repositories
gt repo list

# Create repository
gt repo create my-project --description "Demo project" --private

# Issues
gt issue list owner/project

# Pull requests
gt pr create owner/project --title "Fix bug" --head fix/bug --base main
```

---

## Commands overview

| Command    | Description                     |
|------------|---------------------------------|
| `auth`     | Manage authentication           |
| `repo`     | List / create repositories      |
| `issue`    | List / view / create issues     |
| `pr`       | List / view / create PRs        |

Run `gt <command> --help` for sub-commands and flags.

---

## Configuration

* default path: `~/.config/gt/config.yaml`
* env variables override values:
  * `GT_HOST`
  * `GT_TOKEN`

---

## Contributing

1. Fork and clone the repo.
2. `make dev` – build & run in watch mode.
3. Send pull request.

[*] Keep code straightforward, comments minimal.

---

## License

0BSD © Lain 