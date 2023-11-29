# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
[markdownlint](https://dlaa.me/markdownlint/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

-

## [0.5.3] - 2023-11-29

### Changed in 0.5.3

- In `Dockerfile`
  - `senzingapi-runtime:3.8.0`
  - `golang:1.21.4-bullseye`
- In `package.Dockerfile`
  - `senzingapi-runtime:3.8.0`
  - `golang:1.21.4-bullseye`
- Update dependencies
  - github.com/spf13/cobra v1.8.0

## [0.5.2] - 2023-11-03

### Changed in 0.5.2

- Update dependencies
  - github.com/senzing/check-self v0.0.3
  - github.com/senzing/demo-entity-search v0.0.3
  - github.com/senzing/init-database v0.4.2
  - github.com/senzing/load v0.0.4
  - github.com/senzing/serve-grpc v0.5.4
  - github.com/senzing/serve-http v0.1.2

## [0.5.1] - 2023-10-25

### Changed in 0.5.1

- Refactor to [template-go](https://github.com/Senzing/template-go)
- Update dependencies
  - github.com/senzing/check-self v0.0.2
  - github.com/senzing/demo-entity-search v0.0.2
  - github.com/senzing/explain v0.1.5
  - github.com/senzing/go-cmdhelping v0.1.9
  - github.com/senzing/init-database v0.4.1
  - github.com/senzing/load v0.0.3
  - github.com/senzing/move v0.0.7
  - github.com/senzing/observe v0.1.4
  - github.com/senzing/serve-grpc v0.5.3
  - github.com/senzing/serve-http v0.1.1
  - github.com/senzing/validate v0.0.5

## [0.5.0] - 2023-10-03

### Changed in 0.5.0

- Supports SenzingAPI 3.8.0
- Deprecated functions have been removed
- Add `check-self` and `load` commands
- Update dependencies
  - github.com/senzing/init-database v0.4.0
  - github.com/senzing/move v0.0.6
  - github.com/senzing/serve-grpc v0.5.1
  - github.com/senzing/serve-http v0.1.0
  - github.com/senzing/validate v0.0.4
  - github.com/spf13/viper v1.17.0

## [0.4.4] - 2023-09-01

### Changed in 0.4.4

- Last version before SenzingAPI 3.8.0

## [0.4.3] - 2023-08-29

### Changed in 0.4.3

- Update dependencies
  - downloading github.com/senzing/go-cmdhelping v0.1.8
  - downloading github.com/senzing/move v0.0.5

## [0.4.2] - 2023-08-17

### Changed in 0.4.2

- Added move
- Update dependencies
  - github.com/senzing/go-cmdhelping v0.1.6
  - github.com/senzing/validate v0.0.3

## [0.4.1] - 2023-08-14

### Changed in 0.4.1

- Added validate
- Refactor to `template-go`
- Update dependencies
  - github.com/senzing/explain v0.1.4
  - github.com/senzing/go-cmdhelping v0.1.5
  - github.com/senzing/init-database v0.3.1
  - github.com/senzing/observe v0.1.3
  - github.com/senzing/serve-grpc v0.4.13
  - github.com/senzing/serve-http v0.0.4
  - github.com/senzing/validate v0.0.2

## [0.4.0] - 2023-07-28

### Changed in 0.4.0

- Consolidated `envar`, `help`, and `option` packages and moved them into `go-cmdhelping` repository.
- Moved `constant`, `option`, and `cmdhelper` packages to `go-cmdhelping` repository.
- Updated documentation.
- Updated dependencies
  - github.com/senzing/go-cmdhelping v0.1.3
  - github.com/senzing/observe v0.1.2
  - github.com/senzing/serve-grpc v0.4.11

## [0.3.0] - 2023-06-19

### Added in 0.3.0

- `serve-http` command, although not fully implemented.

### Changed in 0.3.0

- Added additional env vars and options
- In `Dockerfile`, updated to `senzing/senzingapi-runtime:3.5.3`
- Update dependencies
  - github.com/senzing/init-database v0.2.6
  - github.com/senzing/observe v0.1.1
  - github.com/senzing/serve-grpc v0.4.8
  - github.com/senzing/serve-http v0.0.1
  - github.com/spf13/viper v1.16.0
  - github.com/stretchr/testify v1.8.4

## [0.2.8] - 2023-05-26

### Changed in 0.2.8

- Added additional env vars and options
- Update dependencies
  - github.com/senzing/init-database v0.2.5
  - github.com/senzing/serve-grpc v0.4.7

## [0.2.7] - 2023-05-25

### Changed in 0.2.7

- Added additional env vars and options

## [0.2.6] - 2023-05-19

### Changed in 0.2.6

- Update dependencies
  - github.com/senzing/serve-grpc v0.4.6

## [0.2.5] - 2023-05-17

### Added in 0.2.5

- `observe` command
- Added `gosec` GitHub Action

### Changed in 0.2.5

- Update dependencies
  - github.com/senzing/init-database v0.2.4
  - github.com/senzing/observe v0.1.0
  - github.com/senzing/serve-grpc v0.4.5

## [0.2.4] - 2023-05-09

### Changed in 0.2.4

- In `Dockerfile`, updated FROM instruction:
  - `senzing/senzingapi-runtime:3.5.2`
  - `golang:1.20.4@sha256:31a8f92b17829b3ccddf0add184f18203acfd79ccc1bcb5c43803ab1c4836cca`

## [0.2.3] - 2023-04-21

### Changed in 0.2.3

- Update dependencies
  - github.com/senzing/init-database v0.2.2
  - github.com/senzing/serve-grpc v0.4.2

## [0.2.2] - 2023-04-18

### Changed in 0.2.2

- Updated dependencies

## [0.2.1] - 2023-04-03

### Changed in 0.2.1

- In `Dockerfile`, updated FROM instruction:
  - `senzing/senzingapi-runtime:3.5.0`
  - `golang:1.20.2@sha256:f7099345b8e4a93c62dc5102e7eb19a9cdbad12e7e322644eeaba355d70e616d`

## [0.2.0] - 2023-03-29

### Changed in 0.2.0

- Migrated from `initdatabase` to `init-database` command
- Migrated from `servegrpc` to `serve-grpc` command
- Updated dependencies

## [0.1.5] - 2023-03-23

### Changed in 0.1.5

- Updated dependencies

## [0.1.4] - 2023-03-14

### Added to 0.1.4

- Added support for configuration file.

## [0.1.3] - 2023-03-10

### Added to 0.1.3

- Updated dependencies.

## [0.1.2] - 2023-03-10

### Added to 0.1.2

- Updated documentation.

## [0.1.1] - 2023-03-09

### Added to 0.1.1

- Worked on GitHub actions.  No change to `senzing-tools` functionality.

## [0.1.0] - 2023-03-09

### Added to 0.1.0

- Initial functionality
- Commands: `init-database`, `serve-grpc`
