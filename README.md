你好！
很冒昧用这样的方式来和你沟通，如有打扰请忽略我的提交哈。我是光年实验室（gnlab.com）的HR，在招Golang开发工程师，我们是一个技术型团队，技术氛围非常好。全职和兼职都可以，不过最好是全职，工作地点杭州。
我们公司是做流量增长的，Golang负责开发SAAS平台的应用，我们做的很多应用是全新的，工作非常有挑战也很有意思，是国内很多大厂的顾问。
如果有兴趣的话加我微信：13515810775  ，也可以访问 https://gnlab.com/，联系客服转发给HR。
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
