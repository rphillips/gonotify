# gonotify

Pluggable command-line notification tool. Currently supports the following
notification backends:

* Prowl

# Installation

```
go get https://github.com/rphillips/gonotify
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
