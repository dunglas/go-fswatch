# go-fswatch: Go bindings for `libfswatch`

[![Go Reference](https://pkg.go.dev/badge/github.com/dunglas/go-fswatch.svg)](https://pkg.go.dev/github.com/dunglas/go-fswatch)

[Go](https://go.dev) bindings for [`libfswatch`](https://github.com/emcrisostomo/fswatch/blob/master/README.libfswatch.md).

`libfswatch` provides comprehensive, cross-platform file change monitoring capabilities.

## Features

go-fswatch exposes all features provided by `fswatch` and `libfswatch` through an idiomatic Go API:

* cross-platform: inotify (Linux), FSEvents (macOS), Windows, kqueue (*BSD), File Events Notification (Solaris) and polling
* watch files and directories
* recursive watching
* file and directory filtering (using regular expressions)
* events filtering

## Install

First, install `libfswatch`:

1. [Download the latest release (`fswatch-<version>.tar.gz`)](https://github.com/emcrisostomo/fswatch/releases)
2. Compile and install `libfswatch`:

```console
tar xzf fswatch-*.tar.gz
cd fswatch-*
./configure
make
sudo make install
```

Then, you can use this Go module as usual:

```console
go get github.com/dunglas/go-fswatch
```

## Usage and Examples

See [the documentation](https://pkg.go.dev/github.com/dunglas/go-fswatch).

## Cgo

This package depends on [cgo](https://go.dev/blog/cgo).
If you are looking for non-cgo alternatives, see:

* [fsnotify](https://github.com/fsnotify/fsnotify) (doesn't support FSEvents nor polling)
* [notify](https://github.com/rjeczalik/notify) (doesn't support include/exclude filters)

## Credits

Created by [Kévin Dunglas](https://dunglas.dev) and sponsored by [Les-Tilleuls.coop](https://les-tilleuls.coop).
