# rbranch

CLI tool built to simplify Git branches.

> **Status:** Complete — no further development planned, but ready to download and use.

## Overview

### What is rbranch?

rbranch is a command-line tool, built with Go and Bubble Tea, for the Git branch operations you do all day without typing long branch names. Pick a branch from an interactive list to check it out, copy its name or delete it, instead of memorizing lengthy commands, hunting through inactive branches or switching between a GUI and the terminal.

### Key Features

- **Checkout Branches**: Switch to another branch without typing its full name.
- **Delete Branches**: Safely clean up unused local branches.
- **Copy Branch Names**: Copy an entire branch name to your clipboard instantly.

## Getting Started

### Prerequisites

1. **Go** (1.22+) is required to build the tool. Check it by running:

    ```bash
    go version
    ```

### Installation

1. Install the project's dependencies:

    ```bash
    go mod tidy
    ```

2. Build the executable:

    ```bash
    go build
    ```

3. Move it somewhere on your `PATH`:

    ```bash
    sudo mv rbranch /usr/local/bin
    ```

## Usage

### Check Out a Branch

```bash
rbranch
```

Shows the branches in the current Git repository for you to choose from, then checks out the one you pick.

### Copy a Branch Name

```bash
rbranch -c
```

Prompts you to select a branch and copies its name to your clipboard.

### Delete a Branch

```bash
rbranch -d
```

Prompts you to select a branch from the available options and deletes it.

## Contributing

No new features are planned, but it's meant to be used: bug reports through [Issues](../../issues) are welcome, and small fixes may be merged when I get to them. For anything bigger, you're welcome to fork it and take it in your own direction under the license below.

## License

Released under the [MIT License](LICENSE).
