# Toolbox CLI
Toolbox is a command-line interface (CLI) tool designed to help you manage and organize your tasks efficiently. It provides a set of commands to handle various things effectively.

# Installation
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

Set an environment variable:
```
JSON_API_HOST=https://jsonplaceholder.typicode.com
```

Or 
Simply install using
```
go install github.com/priyank-amagi/toolbox
```

# Usage
Toolbox provides the following subcommands and options:

### todos list
Lists todos based on specified options.

Options
- -c <count>: Specifies the total number of todos to be listed.
- -t <todo-type>: Specifies the type of todos to be listed (complete, incomplete, or both).
- -f <filter-type>: Specifies the filter type for todo index (even, odd, or both).

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