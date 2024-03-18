# Toolbox CLI
Toolbox is a command-line interface (CLI) tool designed to help you manage and organize your tasks efficiently. It provides a set of commands to handle various things effectively.

# Installation (using binary)
- Linux
  - [linux-amd64](https://github.com/priyank-amagi/toolbox/releases/download/v0.0.1/toolbox-linux-amd64)
  - [linux-386](https://github.com/priyank-amagi/toolbox/releases/download/v0.0.1/toolbox-linux-386)
- MacOS
  - [macos-amd64](https://github.com/priyank-amagi/toolbox/releases/download/v0.0.1/toolbox-macos-amd64)
- Windows
  - [windows-amd64](https://github.com/priyank-amagi/toolbox/releases/download/v0.0.1/toolbox-windows-amd64.exe)
  - [windows-386](https://github.com/priyank-amagi/toolbox/releases/download/v0.0.1/toolbox-windows-386.exe)

NOTE: Please set the environment variables as belows, to use the CLI
```
JSON_API_HOST=https://jsonplaceholder.typicode.com
```

# Usage
Toolbox provides the following subcommands and options:

### todos list
Lists todos based on specified options.

Options
- -c <count>: Specifies the total number of todos to be listed.
- -t <todo-type>: Specifies the type of todos to be listed (complete, incomplete | <u>*to list both, leave it blank*</u>).
- -f <filter-type>: Specifies the filter type for todo index (even, odd | <u>*to list both, leave it blank*</u>).

## Examples
List all todos:
```
toolbox todos list
```

List only complete todos:
```
toolbox todos list -t complete
```

List odd-numbered incomplete todos:
```
toolbox todos list -t incomplete -f odd
```

List even-numbered 5 todos:
```
toolbox todos list -f even -c 5
```

# For local setup
To install Toolbox, follow these steps:

Clone the repository to your local machine:
```
git clone https://github.com/priyank-amagi/toolbox.git
```

Navigate to the Toolbox directory:
```
cd toolbox
```

Build the CLI tool:
```
go build -o bin/toolbox (with GOOS=linux GOARCH=amd64 based on your host OS and Arch)
```