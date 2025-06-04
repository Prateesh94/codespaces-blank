# broadcast-cli

A command-line interface (CLI) tool written in Go for broadcasting and connecting to WebSocket servers, built using [Cobra](https://github.com/spf13/cobra) and other robust libraries. This project is set up for development inside a Codespace/devcontainer.

## Features

- Start a broadcast server with WebSocket support.
- Connect as a client to WebSocket servers.
- Flexible configuration via command line and configuration files.
- Built-in commands using Cobra CLI framework.

## Project Structure

```
.devcontainer/broadcast-cli/
├── LICENSE
├── go.mod
├── go.sum
├── main.go
└── cmd/
    ├── connect.go    # Implements the connect command
    ├── root.go       # Main CLI entry point
    └── start.go      # Implements the start command (broadcast server)
```

## Getting Started

### Prerequisites

- Go 1.24 or later (as specified in `go.mod`)
- [Git](https://git-scm.com/) for cloning the repository

### Installation

Clone this repository and navigate to the CLI source:

```bash
git clone https://github.com/Prateesh94/codespaces-blank.git
cd codespaces-blank/.devcontainer/broadcast-cli
go build -o broadcast-cli
```

### Usage

Run the CLI tool and explore available commands:

```bash
./broadcast-cli --help
```

Example: Start a broadcast server

```bash
./broadcast-cli start
```

Example: Connect to a broadcast server

```bash
./broadcast-cli connect ws://localhost:8080
```

### Configuration

- The CLI supports configuration via flags and config files (see `cmd/root.go`).
- Common flags and configuration options can be found using the `--help` flag on each command.

## Dependencies

Key dependencies from `go.mod`:
- [Cobra](https://github.com/spf13/cobra): Command-line framework
- [Viper](https://github.com/spf13/viper): Configuration management
- [Gorilla Mux](https://github.com/gorilla/mux): HTTP routing
- [Gorilla WebSocket](https://github.com/gorilla/websocket): WebSocket support

## License

This project is licensed under the terms of the LICENSE file in this directory.

---

**Note:** This CLI is intended for development and prototyping inside a Codespace/devcontainer setup. For production usage, review and extend as needed.
