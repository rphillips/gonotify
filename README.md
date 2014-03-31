[![Build Status](https://travis-ci.org/rphillips/gonotify.svg?branch=master)](https://travis-ci.org/rphillips/gonotify)

# gonotify

Pluggable command-line notification tool. Currently supports the following
notification backends:

* Prowl

# Installation

```
go get https://github.com/rphillips/gonotify
```

# Configuration

~/.gonotify: Insert your API key

```
[gonotify]
backend = prowl

[prowl]
api_key =
```

# Usage

Notify yourself after a build:

```
make ; gonotify -event="build done $?"
```

# License

Apache 2

# Contributions

Pull requests are always welcome.

