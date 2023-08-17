# Changelog

## Overview

All notable changes to this project will be documented in this file.

The format is based on [Keep a
Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres to
[Semantic Versioning](https://semver.org/spec/v2.0.0.html).

Please [open an issue](https://github.com/atc0005/check-process/issues) for any
deviations that you spot; I'm still learning!.

## Types of changes

The following types of changes will be recorded in this file:

- `Added` for new features.
- `Changed` for changes in existing functionality.
- `Deprecated` for soon-to-be removed features.
- `Removed` for now removed features.
- `Fixed` for any bug fixes.
- `Security` in case of vulnerabilities.

## [Unreleased]

- placeholder

## [v0.3.3] - 2023-08-17

### Added

- (GH-117) Add initial automated release notes config
- (GH-119) Add initial automated release build workflow

### Changed

- Dependencies
  - `Go`
    - `1.19.11` to `1.20.7`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.11.3` to `go-ci-oldstable-build-v0.13.4`
  - `rs/zerolog`
    - `v1.29.1` to `v1.30.0`
  - `golang.org/x/sys`
    - `v0.10.0` to `v0.11.0`
- (GH-121) Update Dependabot config to monitor both branches
- (GH-141) Update project to Go 1.20 series

## [v0.3.2] - 2023-07-13

### Overview

- RPM package improvements
- Bug fixes
- Dependency updates
- built using Go 1.19.11
  - Statically linked
  - Windows (x86, x64)
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.10` to `1.19.11`
  - `atc0005/go-nagios`
    - `v0.15.0` to `v0.16.0`
  - `atc0005/go-ci`
    - `go-ci-oldstable-build-v0.11.0` to `go-ci-oldstable-build-v0.11.3`
  - `golang.org/x/sys`
    - `v0.9.0` to `v0.10.0`
- (GH-111) Update RPM postinstall scripts to use restorecon

### Fixed

- (GH-108) README missing performance data metrics table
- (GH-109) Correct logging format listed in README

## [v0.3.1] - 2023-06-16

### Overview

- Bug fixes
- GitHub Actions workflow updates
- Dependency updates
- built using Go 1.19.10
  - Statically linked
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.9` to `1.19.10`
  - `atc0005/go-ci` build image
    - `go-ci-oldstable-build-v0.10.5` to `go-ci-oldstable-build-v0.11.0`
  - `atc0005/go-nagios`
    - `v0.14.0` to `v0.15.0`
  - `golang.org/x/sys`
    - `v0.8.0` to `v0.9.0`
  - `mattn/go-isatty`
    - `v0.0.18` to `v0.0.19`
- (GH-99) Update vuln analysis GHAW to remove on.push hook

### Fixed

- (GH-95) Disable depguard linter
- (GH-96) Add missing branding flag support
- (GH-101) Restore local CodeQL workflow

## [v0.3.0] - 2023-05-11

### Overview

- Build improvements
- Bug fixes
- Dependency updates
- built using Go 1.19.9
  - Statically linked
  - Linux (x86, x64)

### Added

- (GH-82) Add rootless container builds via Docker/Podman

### Changed

- Dependencies
  - `Go`
    - `1.19.7` to `1.19.9`
  - `atc0005/go-ci` build image
    - `go-ci-oldstable-build-v0.9.1` to `go-ci-oldstable-build-v0.10.5`
  - `rs/zerolog`
    - `v1.29.0` to `v1.29.1`
  - `golang.org/x/sys`
    - `v0.6.0` to `v0.8.0`
  - `mattn/go-isatty`
    - `v0.0.17` to `v0.0.18`

### Fixed

- (GH-74) Fix CHANGELOG entry indentation
- (GH-76) Update vuln analysis GHAW to use on.push hook
- (GH-90) Fix revive linter errors

## [v0.2.1] - 2023-03-09

### Overview

- Dependency updates
- built using Go 1.19.7
  - Statically linked
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.6` to `1.19.7`
  - `atc0005/go-ci` build image
    - `go-ci-oldstable-build-v0.9.0` to `go-ci-oldstable-build-v0.9.1`

## [v0.2.0] - 2023-03-07

### Overview

- Add support for generating packages
- Generated binary changes
  - filename patterns
  - compression
  - executable metadata
- Build improvements
- built using Go 1.19.6
  - Statically linked
  - Linux (x86, x64)

### Added

- (GH-53) Generate RPM/DEB packages using nFPM

### Changed

- (GH-52) Switch to semantic versioning (semver) compatible versioning
  pattern
- (GH-54) Add version metadata to Windows executables
- (GH-55) Makefile: Compress binaries and use fixed filenames
- (GH-56) Makefile: Refresh recipes to add "standard" set, new
  package-related options
- (GH-57) Build dev/stable releases using go-ci Docker image

## [v0.1.2] - 2023-03-07

### Overview

- Bug fixes
- Dependency updates
- GitHub Actions Workflow updates
- built using Go 1.19.6
  - Statically linked
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.4` to `1.19.6`
  - `atc0005/go-nagios`
    - `v0.10.2` to `v0.14.0`
  - `rs/zerolog`
    - `v1.28.0` to `v1.29.0`
  - `golang.org/x/sys`
    - `v0.3.0` to `v0.6.0`
  - `mattn/go-isatty`
    - `v0.0.16` to `v0.0.17`
- GitHub Actions
  - (GH-40) Add Go Module Validation, Dependency Updates jobs
  - (GH-47) Drop `Push Validation` workflow
  - (GH-48) Rework workflow scheduling
  - (GH-50) Remove `Push Validation` workflow status badge

### Fixed

- (GH-35) Drop plugin runtime tracking, update library usage
- (GH-41) Add missing Makefile usage entry for release build
- (GH-59) Use UNKNOWN state for perfdata add failures
- (GH-60) Use UNKNOWN state for invalid command-line args
- (GH-61) Remove duplicate perfdata add step
- (GH-62) Use UNKNOWN state for evaluation failures

## [v0.1.1] - 2022-12-07

### Overview

- Bug fixes
- Dependency updates
- built using Go 1.19.4
  - Statically linked
  - Linux (x86, x64)

### Changed

- Dependencies
  - `Go`
    - `1.19.3` to `1.19.4`
  - `golang.org/x/sys`
    - `v0.2.0` to `v0.3.0`
- (GH-26) Exclude process ID of running tool

### Fixed

- (GH-25) Minor refactoring to resolve gocognit warnings
- (GH-27) Fix overview statements

## [v0.1.0] - 2022-12-03

### Overview

- Initial release
- built using Go 1.19.3
  - Statically linked
  - Linux (x86, x64)

### Added

Initial release!

This release provides early release versions of two tools used to monitor
processes on Linux distros:

| Tool Name       | Overall Status | Description                                                        |
| --------------- | -------------- | ------------------------------------------------------------------ |
| `check_process` | Alpha          | Nagios plugin used to monitor for problematic process states.      |
| `lsps`          | Alpha          | Small CLI tool to list processes with known problematic processes. |

See the project README for additional details.

[Unreleased]: https://github.com/atc0005/check-process/compare/v0.3.3...HEAD
[v0.3.3]: https://github.com/atc0005/check-process/releases/tag/v0.3.3
[v0.3.2]: https://github.com/atc0005/check-process/releases/tag/v0.3.2
[v0.3.1]: https://github.com/atc0005/check-process/releases/tag/v0.3.1
[v0.3.0]: https://github.com/atc0005/check-process/releases/tag/v0.3.0
[v0.2.1]: https://github.com/atc0005/check-process/releases/tag/v0.2.1
[v0.2.0]: https://github.com/atc0005/check-process/releases/tag/v0.2.0
[v0.1.2]: https://github.com/atc0005/check-process/releases/tag/v0.1.2
[v0.1.1]: https://github.com/atc0005/check-process/releases/tag/v0.1.1
[v0.1.0]: https://github.com/atc0005/check-process/releases/tag/v0.1.0
