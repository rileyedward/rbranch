# rbranch

rbranch is a CLI tool built with Go and Bubble Tea designed to simplify your Git workflow. If you’re tired of typing long and cumbersome branch names, rbranch can help. With just a few commands and flags, you can effortlessly perform common Git branch operations and streamline your development process.

## Getting Started

### Prerequisites

Ensure you have the following prerequisites installed on your system. You can verify each installation by running the provided commands in your terminal.

1. **Go** is required for the application. Check if Go is installed by running:

    ```bash
    go version
    ```

### Installation

1. Install project dependencies:

    ```bash
    go mod tidy
    ```

2. Build the application:

    ```bash
    go build
    ```

3. Install the executable:

    ```bash
    sudo mv rbranch /usr/local/bin
    ```

## Usage

### Checkout a Branch

To checkout a branch from your current Git repository, simply run:

```bash
rbranch
```

This will present a list of available branches for you to choose from.

### Copy a Branch

To copy a branch name to your clipboard, use the following command:

```bash
rbranch -c
```

You’ll be prompted to select a branch, and its name will be copied to your clipboard.

### Delete a Branch

To delete a branch, run:

```bash
rbranch -d
```

You will be prompted to select a branch from the available options for deletion.
