# Warpdir

Warpdir is a simple command line tool to create and manage shortcuts to directories.

## Installation

```bash
go install github.com/nousefreak/warpdir@latest
```

Or download the latest release and add the binary to your PATH.

```shell
warpdir install
```

Once installed you can run `wd help` to see the help page.

## Usage

```bash
# Add a warp point to the current directory
wd add <name>

# Add a warp point to a specific directory
wd add <name> <path>

# Jump to a warp point
wd <name>

# Jump to previous warp point
wd -

# Fuzzy search warp points
wd

# List all warp points
wd list

# Delete a warp point
wd delete <name>
```
