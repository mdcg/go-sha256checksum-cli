# go-sha256checksum-cli

[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/mdcg/go-sha256checksum-cli/blob/main/LICENSE)

## Introduction

CLI written in Go with two main features:

- Generate the checksum for a specific file.
- Compare the checksum of two or more files in order to compare whether they are all the same.

## Usage

Compile the program:

```
go build
```

For usage details use:

```
./go-sha256checksum-cli --help
```

Generate the checksum for a specific file:

```
./go-sha256checksum-cli file <path>
```

Compare checksums of two or more files:

```
./go-sha256checksum-cli compare <path_1> <path_2> ... <path_n>
```

## Contributing

Feel free to do whatever you want with this project. :-)