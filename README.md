# Code Extractor

A Go program that extracts code blocks from Markdown files.

## Usage

```shell
go run main.go [options] <markdown_file>
```

### Options

- `-l, --language`: Specify the language to extract (default: "shell")
  - Use "all" to extract all code blocks regardless of language
  - Flags can be placed before or after the filename

### Examples

1. Extract all shell code blocks:

   ```shell
   go run main.go example.md
   ```

2. Extract Python code blocks (flags can be placed before or after the filename):

   ```shell
   go run main.go -l python example.md
   go run main.go example.md -l python
   ```

3. Extract all code blocks regardless of language:

   ```shell
   go run main.go -l all example.md
   go run main.go example.md -l all
   ```

## Code Block Comments

You can use language-specific comments in your code blocks to provide context or instructions. For example:

```shell
# Run this command on the server
sudo systemctl status minio
```

```shell
# Run this command on the client
mc ls lab
```

The program will preserve all comments in the extracted code blocks.

## Installation

1. Clone the repository
2. Install dependencies:

   ```shell
   go mod download
   ```

## Building

To build the program:

```shell
go build -o code_extractor
```
