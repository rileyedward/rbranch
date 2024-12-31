# rbranch

CLI tool built to simplify Git branches.

## Overview

### What is rbranch?

rbranch is a CLI tool built with Go and Bubble Tea designed to simplify your Git workflow. If you’re tired of typing long and cumbersome branch names, rbranch can help. With just a few commands and flags, you can effortlessly perform common Git branch operations and streamline your development process.

### Why Use rbranch?

Typing long branch names, searching through inactive branches, and managing Git operations can be tedious. rbranch is designed for developers who want a faster, cleaner way to interact with Git branches—no need to memorize lengthy commands or switch between GUIs and terminals.

### Key Features

- **Checkout branches**: Easily switch to another branch without typing its full name.
- **Delete branches**: Safely cleanup unused local branches.
- **Copy branch names**: Save time ny instantly copying an entire branch name to your clipboard.

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
