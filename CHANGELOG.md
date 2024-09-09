# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

<!-- START Unreleased -->

### Added

* `utils`: New package with `Pointer`/`Coalesce` functions via generics ([#72])
* `http/middleware`: New `RequestID` middleware function ([#73])

### Others

* **BREAKING**: Set Go 1.18 as minimum supported version ([#68])

[#68]: https://github.com/loozhengyuan/grench/pull/68
[#72]: https://github.com/loozhengyuan/grench/pull/72
[#73]: https://github.com/loozhengyuan/grench/pull/73

<!-- END Unreleased -->

## [v0.7.0] - 2021-08-26

<!-- START v0.7.0 -->

## Others

* **BREAKING**: Drop superfluous `pkg` directory ([#52])
* **BREAKING**: Move packages into new `utils` package ([#53])

[#52]: https://github.com/loozhengyuan/grench/pull/52
[#53]: https://github.com/loozhengyuan/grench/pull/53

<!-- END v0.7.0 -->

## [v0.6.0] - 2021-04-27

<!-- START v0.6.0 -->

### Added

* `coalesce`: New package for coalescing operations ([#47])
* `config`: New package for managing app config ([#48])

[#47]: https://github.com/loozhengyuan/grench/pull/47
[#48]: https://github.com/loozhengyuan/grench/pull/48

<!-- END v0.6.0 -->

## [v0.5.0] - 2021-04-15

<!-- START v0.5.0 -->

### Added

* `auth/oauth2`: New `TokenResponse` type ([#37])
* `auth/oidc`: New `TokenResponse` type ([#38])
* `kv`: New key-value storage interface ([#41])
* `kv/mem`: In-memory implementation of `kv.Store` ([#41])
* `kv/fs`: Local filesystem implementation of `kv.Store` ([#41])

[#37]: https://github.com/loozhengyuan/grench/pull/37
[#38]: https://github.com/loozhengyuan/grench/pull/38
[#41]: https://github.com/loozhengyuan/grench/pull/41

<!-- END v0.5.0 -->

## [v0.4.0] - 2021-04-04

<!-- START v0.4.0 -->

### Added

* `auth/oidc`: New `Provider` struct ([#28])
* `auth/oauth2/pkce`: New `CodeVerifier` struct ([#32])

[#28]: https://github.com/loozhengyuan/grench/pull/28
[#32]: https://github.com/loozhengyuan/grench/pull/32

<!-- END v0.4.0 -->

## [v0.3.0] - 2021-03-05

<!-- START v0.3.0 -->

### Added

* `health`: Health checks now run asynchronously ([#22])
* `health`: JSON struct tags for `health.Info` ([#24])

[#22]: https://github.com/loozhengyuan/grench/pull/22
[#24]: https://github.com/loozhengyuan/grench/pull/24

### Changed

* **BREAKING**: `health`: Disallow checks from returning errors ([#21])
* **BREAKING**: `pointer`: Rename functions ([#25])

[#21]: https://github.com/loozhengyuan/grench/pull/21
[#25]: https://github.com/loozhengyuan/grench/pull/25

<!-- END v0.3.0 -->

## [v0.2.0] - 2021-02-24

<!-- START v0.2.0 -->

### Changed

* **BREAKING**: `health`: Drop register pattern on the `health.Checker` interface ([#16])

[#16]: https://github.com/loozhengyuan/grench/pull/16

### Others

* Set go.mod `go` directive to minimum supported Go version ([#12])
* Support for Go 1.16 ([#13])

[#12]: https://github.com/loozhengyuan/grench/pull/12
[#13]: https://github.com/loozhengyuan/grench/pull/13

<!-- END v0.2.0 -->

## [v0.1.0] - 2021-02-17

<!-- START v0.1.0 -->

Initial release.

<!-- END v0.1.0 -->

[Unreleased]: https://github.com/loozhengyuan/grench/compare/v0.7.0...HEAD
[v0.7.0]: https://github.com/loozhengyuan/grench/releases/tag/v0.7.0
[v0.6.0]: https://github.com/loozhengyuan/grench/releases/tag/v0.6.0
[v0.5.0]: https://github.com/loozhengyuan/grench/releases/tag/v0.5.0
[v0.4.0]: https://github.com/loozhengyuan/grench/releases/tag/v0.4.0
[v0.3.0]: https://github.com/loozhengyuan/grench/releases/tag/v0.3.0
[v0.2.0]: https://github.com/loozhengyuan/grench/releases/tag/v0.2.0
[v0.1.0]: https://github.com/loozhengyuan/grench/releases/tag/v0.1.0
