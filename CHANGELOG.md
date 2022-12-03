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

[Unreleased]: https://github.com/atc0005/check-process/compare/v0.1.0...HEAD
[v0.1.0]: https://github.com/atc0005/check-process/releases/tag/v0.1.0
