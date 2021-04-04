# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

<!-- START Unreleased -->

### Added

* New `auth/oidc` package ([#28])

<!-- END Unreleased -->

## [v0.3.0] - 2021-03-05

<!-- START v0.3.0 -->

### Added

* Health checks now run asynchronously ([#22])
* JSON struct tags for health.Info ([#24])

### Changed

* Disallow checks from returning errors ([#21])
* Rename functions in pointer package ([#25])

<!-- END v0.3.0 -->

## [v0.2.0] - 2021-02-24

<!-- START v0.2.0 -->

### Added

* Support for Go 1.16 ([#13])

### Changed

* Drop register pattern on the health.Checker interface ([#16])

### Fixed

* Set go.mod `go` directive to minimum supported Go version ([#12])

<!-- END v0.2.0 -->

## [v0.1.0] - 2021-02-17

<!-- START v0.1.0 -->

Initial release.

<!-- END v0.1.0 -->

[#12]: https://github.com/loozhengyuan/grench/pull/12
[#13]: https://github.com/loozhengyuan/grench/pull/13
[#16]: https://github.com/loozhengyuan/grench/pull/16
[#21]: https://github.com/loozhengyuan/grench/pull/21
[#22]: https://github.com/loozhengyuan/grench/pull/22
[#24]: https://github.com/loozhengyuan/grench/pull/24
[#25]: https://github.com/loozhengyuan/grench/pull/25
[#28]: https://github.com/loozhengyuan/grench/pull/28

[Unreleased]: https://github.com/loozhengyuan/grench/compare/v0.3.0...HEAD
[v0.3.0]: https://github.com/loozhengyuan/grench/releases/tag/v0.3.0
[v0.2.0]: https://github.com/loozhengyuan/grench/releases/tag/v0.2.0
[v0.1.0]: https://github.com/loozhengyuan/grench/releases/tag/v0.1.0
