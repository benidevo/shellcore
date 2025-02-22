# ShellCore

A Unix shell implementation in Go featuring command execution, built-in commands, and path resolution. The shell supports basic operations with robust error handling and command parsing.

## Core Features

- Command Execution:
  - System command execution with argument parsing
  - Single quote handling for preserving whitespaces
  - PATH resolution for finding executables

- Built-in Commands:
  - `cd <path>` - Navigate directories
  - `pwd` - Show current directory
  - `echo <text>` - Display text
  - `type <command>` - Show command type
  - `cat <file>` - Display file contents
  - `exit 0` - Exit shell

## Quick Start

Prerequisites: Go 1.20+

Run the shell:

```bash
./run.sh
# OR
make run
```

## Examples

Interactive shell usage:

```bash
$ echo 'Hello World with spaces'
Hello World with spaces

$ type ls
ls is /bin/ls

$ type echo
echo is a shell builtin

$ pwd
/home/user

$ cd /tmp
```
