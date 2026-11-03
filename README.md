# hexdump

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21%2B-00ADD8?logo=go" alt="Go version" />
  <img src="https://img.shields.io/badge/Platform-Windows%20%7C%20Linux%20%7C%20macOS-4B5563" alt="Platforms" />
  <img src="https://img.shields.io/badge/License-MIT-green" alt="License" />
</p>

A small Go utility for viewing binary data in a readable hex dump format directly from the terminal.

## Why use it?

`hexdump` is useful when you want to inspect raw bytes, debug binary files, analyze protocol payloads, or quickly check the contents of a file without opening it in a binary editor.

## Features

- dump binary data as `offset: hex bytes | ascii |`
- read input from a file with `-f`
- decode a hex string from command-line arguments
- hide the ASCII column with `-a`
- produce builds for multiple platforms

## Quick start

### Install

```bash
go mod tidy
```

### Build

```bash
./build.sh
```

### Run

Dump a file:

```bash
./hexdump -f somefile.bin
```

Dump a file without ASCII output:

```bash
./hexdump -f somefile.bin -a
```

Decode a hex string:

```bash
./hexdump 48656c6c6f20776f726c6421
```

Or with spaces:

```bash
./hexdump 48 65 6c 6c 6f 20 77 6f 72 6c 64 21
```

## Example output

```text
00000000: 48 65 6c 6c 6f 2c 20 57 6f 72 6c 64 21 0a | Hello, World!.|
```

## Supported platforms

The included build script targets:

- Windows AMD64
- Windows 386
- Linux AMD64
- Linux ARM64
- macOS AMD64
- macOS ARM64

## Testing

```bash
go test ./...
```

## Project status

This project is intentionally lightweight and focused on a single task: making raw binary content readable in a terminal with minimal friction.
