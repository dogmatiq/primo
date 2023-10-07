# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to
[Semantic Versioning].

<!-- references -->

[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html

## Unreleased

### Added

- Generate a `NewXXX()` "constructor" function for each Protocol Buffers message
  type.

### Changed

- Mutator methods now construct and return a new message when called on a `nil`
  message.

## [0.1.2] - 2023-10-07

### Changed

- Generated mutator methods now return the mutated message, allowing "fluent"
  style method chaining.

## [0.1.1] - 2023-09-12

### Fixed

- Fixed compilation error.

## [0.1.0] - 2023-09-11

- Initial release

<!-- references -->

[unreleased]: https://github.com/dogmatiq/primo
[0.1.0]: https://github.com/dogmatiq/primo/releases/tag/v0.1.0
[0.1.1]: https://github.com/dogmatiq/primo/releases/tag/v0.1.1
[0.1.2]: https://github.com/dogmatiq/primo/releases/tag/v0.1.2

<!-- version template
## [0.0.1] - YYYY-MM-DD

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security
-->
