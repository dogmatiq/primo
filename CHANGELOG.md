# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog], and this project adheres to
[Semantic Versioning].

<!-- references -->

[Keep a Changelog]: https://keepachangelog.com/en/1.0.0/
[Semantic Versioning]: https://semver.org/spec/v2.0.0.html

## [0.3.1] - 2024-07-16

### Added

- Added `MustSwitch_XXX()` and `MustMap_XXX()` functions for one-of groups.
  These functions panic if the message or the one-of discriminator field is
  `nil`. This is the same behavior that `Switch_XXX()` and `Map_XXX()` had prior
  to [0.3.0].

## [0.3.0] - 2024-07-11

### Changed

- **[BC]** Added a mandatory `default_` parameter to the `Map_XXX()` and
  `Switch_XXX()` functions generated for one-of fields.

## [0.2.4] - 2024-07-11

### Fixed

- Fixed panic when calling `Map_XXX()` and `Switch_XXX()` with a `nil` message.

## [0.2.3] - 2024-05-22

### Added

- Generate a basic stub implementation of each gRPC client interface.

## [0.2.2] - 2024-05-15

### Fixed

- Generate code for nested message and enum types.

## [0.2.1] - 2024-04-23

### Added

- Generate `TryGetXXX()` methods for each option within one-of groups.

## [0.2.0] - 2023-10-19

### Added

- Generate an `XXXBuilder` type for each Protocol Buffers message type, which
  constructs messages from a configurable prototype message.
- Generate a `Map_XXX()` function for each Protocol Buffers enumeration and
  one-of type.

### Removed

- **[BC]** Removed the return value from generated `Switch_XXX()` functions, use `Map_XXX()` instead.
- **[BC]** Removed generated `NewXXX()` functions, use `NewXXXBuilder()` instead.
- **[BC]** Removed the return value from generated mutator methods.

## [0.1.5] - 2023-10-18

### Fixed

- Fix incorrect switch case generation when the name of a "one-of discriminator"
  type conflicts with another Go generated type.

## [0.1.4] - 2023-10-10

### Added

- Generate a `Switch_XXX()` function for each Protocol Buffers enumeration type.

## [0.1.3] - 2023-10-07

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
[0.1.3]: https://github.com/dogmatiq/primo/releases/tag/v0.1.3
[0.1.4]: https://github.com/dogmatiq/primo/releases/tag/v0.1.4
[0.1.5]: https://github.com/dogmatiq/primo/releases/tag/v0.1.5
[0.2.0]: https://github.com/dogmatiq/primo/releases/tag/v0.2.0
[0.2.1]: https://github.com/dogmatiq/primo/releases/tag/v0.2.1
[0.2.2]: https://github.com/dogmatiq/primo/releases/tag/v0.2.2
[0.2.3]: https://github.com/dogmatiq/primo/releases/tag/v0.2.3
[0.2.4]: https://github.com/dogmatiq/primo/releases/tag/v0.2.4
[0.3.0]: https://github.com/dogmatiq/primo/releases/tag/v0.3.0
[0.3.1]: https://github.com/dogmatiq/primo/releases/tag/v0.3.1

<!-- version template
## [0.0.1] - YYYY-MM-DD

### Added
### Changed
### Deprecated
### Removed
### Fixed
### Security
-->
