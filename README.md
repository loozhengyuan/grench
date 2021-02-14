# grench

[![PkgGoDev](https://pkg.go.dev/badge/github.com/loozhengyuan/grench)](https://pkg.go.dev/github.com/loozhengyuan/grench)
[![Go Report Card](https://goreportcard.com/badge/github.com/loozhengyuan/grench)](https://goreportcard.com/report/github.com/loozhengyuan/grench)
![ci](https://github.com/loozhengyuan/grench/workflows/ci/badge.svg)

`grench` is an opinionated, zero-dependency utility library for developing Go applications.

The `grench` (_derived from the words 'go' and 'wrench'_) project aims to provide an importable, version-controlled package for commonly-used utility packages. These packages are usually too trivial to have its own dedicated package, yet pasting them into all your Go applications makes it hard to maintain whenever there are changes/fixes. The `grench` package provides all of these so you can focus on getting your Go application up and running.

## Install

To install `grench`, use `go get` to fetch the latest version:

```shell
go get -u github.com/loozhengyuan/grench
```

Once installed, you may import it directly into your Go application:

```shell
import "github.com/loozhengyuan/grench"
```

## License

[MIT](https://choosealicense.com/licenses/mit/)
