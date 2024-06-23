# Unix-Based Commands Implemented in Go
This repository contains implementations of several Unix-based commands in Go. Each command is implemented to mimic its standard behavior without additional flags, with specific enhancements noted for `cat -n`, `head`, and `tail`.

## Commands Implemented

### 1. echo

The `echo` command in Go prints its arguments to standard output.

### 2. env

The `env` command in Go lists the current environment variables and their values.

### 3. cat

The `cat` command in Go concatenates files and prints them to standard output. Implemented with:
- `cat -n`: Prints each line with line numbers.

### 4. wc

The `wc` command in Go counts lines, words, and bytes in a file and prints the counts.

### 5. head

The `head` command in Go prints the first N lines of a file. Implemented with:
- `<number> <filepath>`: Prints the first N lines of the specified file.

### 6. tail

The `tail` command in Go prints the last N lines of a file. Implemented with:
- `<number> <filepath>`: Prints the last N lines of the specified file.

### 7. yes

The `yes` command in Go continuously prints a line of 'y' to standard output or a word if specified.

### 8. true

The `true` command in Go does nothing except return a successful exit status.

### 9. false

The `false` command in Go does nothing except return a non-successful exit status.

### 10. tree

The `tree` command in Go recursively lists contents of directories in a tree-like format.

## Notes
- The implementations are designed to mirror standard Unix behavior without extensive flag support, except where specified above.
