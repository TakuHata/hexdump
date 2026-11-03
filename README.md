# hexdump

A small command-line utility for printing binary data in a readable hex dump format.

## Overview

`hexdump` displays file contents or hex-encoded input as a classic hex dump with optional ASCII output. It is useful for debugging binary files, protocol payloads, and raw byte sequences.

## Features

- dump a file as `offset: hex bytes | ascii |`
- parse and decode a hex string passed as arguments
- optionally disable the ASCII column with `-a`
- build binaries for multiple platforms with the included script

## Install

Install the binary directly from GitHub:

```bash
go install github.com/TakuHata/hexdump@latest
```

If the binary is not in your `PATH`, make sure Go's bin directory is added, typically:

```bash
export PATH="$(go env GOPATH)/bin:$PATH"
```

## Build from source

```bash
go build -o hexdump .
```

Or use the bundled build script:

```bash
./build.sh
```

The script generates binaries for:

- Windows AMD64
- Windows 386
- Linux AMD64
- Linux ARM64
- macOS AMD64
- macOS ARM64

## Usage

### Dump a file

```bash
hexdump -f somefile.bin
```

### Dump a file without ASCII output

```bash
hexdump -f somefile.bin -a
```

### Decode a hex string

```bash
hexdump 48656c6c6f20776f726c6421
```

You can also pass the hex value as separate arguments:

```bash
hexdump 48 65 6c 6c 6f 20 77 6f 72 6c 64 21
```

## Example output

```text
00000000: 48 65 6c 6c 6f 2c 20 57 6f 72 6c 64 21 0a | Hello, World!.|
```

## Testing

```bash
go test ./...
```

## Notes

This project is intentionally lightweight and focused on a single goal: making raw binary data easy to inspect in a terminal.
