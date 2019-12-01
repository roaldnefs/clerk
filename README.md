# clerk

[![Travis CI](https://img.shields.io/travis/roaldnefs/clerk.svg)](https://travis-ci.org/roaldnefs/clerk)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4.svg)](https://godoc.org/github.com/roaldnefs/clerk)
[![Github All Releases](https://img.shields.io/github/downloads/roaldnefs/clerk/total.svg)](https://github.com/roaldnefs/clerk/releases)
[![GitHub](https://img.shields.io/github/license/roaldnefs/clerk.svg)](https://github.com/roaldnefs/clerk/blob/master/LICENSE)

`clerk` is a command line utility that can create and verify signatures of a Docker image using GPG. 

**Note:** this utility is for testing purposes only, do not use this in production!

* [Installation](README.md#installation)
     * [Binaries](README.md#binaries)
     * [Via Go](README.md#via-go)
* [Usage](README.md#usage)

## Installation

### Binaries

For installation instructions from binaries please visit the [Releases Page](https://github.com/roaldnefs/clerk/releases).

### Via Go

```console
$ go get github.com/roaldnefs/clerk
```

## Usage

```
usage: clerk [<flags>] <command> [<args> ...]

Command line utility that can create and verify signatures of a Docker image using GPG.

Flags:
  --help     Show context-sensitive help (also try --help-long and --help-man).
  --version  Show application version.

Commands:
  help [<command>...]
    Show help.

  sign --output=OUTPUT <manifest> <docker-reference> <key-fingerprint>
    Sign a Docker image.

  verify <manifest> <docker-reference> <key-fingerprint> <signature>
    Verify a Docker image.
```

## Authors

This project is heavily based on [skopeo](https://github.com/containers/skopeo), with the modified work by [Roald Nefs](https://github.com/roaldnefs).