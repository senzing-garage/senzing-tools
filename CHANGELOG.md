# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
[markdownlint](https://dlaa.me/markdownlint/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

-

## [0.2.5] - 2023-05-17

### Added in 0.2.5

- `observe` command
- Added `gosec` GitHub Action
- Update dependencies
  - github.com/senzing/init-database v0.2.4
  - github.com/senzing/observe v0.1.0
  - github.com/senzing/serve-grpc v0.4.5

### Changed in 0.2.5

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
