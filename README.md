<!-- omit in toc -->
# check-process

Go-based tooling used to monitor processes.

[![Latest Release](https://img.shields.io/github/release/atc0005/check-process.svg?style=flat-square)](https://github.com/atc0005/check-process/releases/latest)
[![Go Reference](https://pkg.go.dev/badge/github.com/atc0005/check-process.svg)](https://pkg.go.dev/github.com/atc0005/check-process)
[![go.mod Go version](https://img.shields.io/github/go-mod/go-version/atc0005/check-process)](https://github.com/atc0005/check-process)
[![Lint and Build](https://github.com/atc0005/check-process/actions/workflows/lint-and-build.yml/badge.svg)](https://github.com/atc0005/check-process/actions/workflows/lint-and-build.yml)
[![Project Analysis](https://github.com/atc0005/check-process/actions/workflows/project-analysis.yml/badge.svg)](https://github.com/atc0005/check-process/actions/workflows/project-analysis.yml)

<!-- omit in toc -->
## Table of Contents

- [Project home](#project-home)
- [Overview](#overview)
- [Features](#features)
  - [`check_process` plugin](#check_process-plugin)
  - [`lsps` CLI tool](#lsps-cli-tool)
- [Changelog](#changelog)
- [Requirements](#requirements)
  - [Building source code](#building-source-code)
  - [Running](#running)
- [Installation](#installation)
  - [From source](#from-source)
  - [Using release binaries](#using-release-binaries)
  - [Deployment](#deployment)
- [Configuration](#configuration)
  - [Command-line arguments](#command-line-arguments)
    - [`check_process`](#check_process)
    - [`lsps`](#lsps)
- [Process states](#process-states)
  - [Known states](#known-states)
    - [v2.6.32 (RHEL 6)](#v2632-rhel-6)
    - [v3.10 (RHEL 7)](#v310-rhel-7)
    - [v4.18 (RHEL 8)](#v418-rhel-8)
    - [v.5.14 (RHEL 9)](#v514-rhel-9)
    - [Summary](#summary)
  - [Process state to plugin state mappings](#process-state-to-plugin-state-mappings)
- [Examples](#examples)
  - [`OK` result](#ok-result)
  - [`WARNING` result](#warning-result)
  - [`CRITICAL` result](#critical-result)
- [License](#license)
- [References](#references)

## Project home

See [our GitHub repo][repo-url] for the latest code, to file an issue or
submit improvements for review and potential inclusion into the project.

## Overview

This repo is intended to provide various tools used to monitor processes.

| Tool Name       | Overall Status | Description                                                     |
| --------------- | -------------- | --------------------------------------------------------------- |
| `check_process` | Alpha          | Nagios plugin used to monitor processes for problematic states. |
| `lsps`          | Alpha          | Small CLI tool to list processes with known problematic states. |

## Features

### `check_process` plugin

Nagios plugin (`check_process`) used to monitor for problematic process states
on Linux distros.

NOTE: The intent is to support multiple operating systems, but as of this
writing Linux is the only supported OS

- Optional branding "signature"
  - used to indicate what Nagios plugin (and what version) is responsible for
    the service check result

- Optional, leveled logging using `rs/zerolog` package
  - JSON-format output (to `stderr`)
  - choice of `disabled`, `panic`, `fatal`, `error`, `warn`, `info` (the
    default), `debug` or `trace`

NOTE: This tool ignores its own process entry when reporting running processes.

### `lsps` CLI tool

Small CLI tool to list processes with known problematic processes.

- Optional expanded or "all" listing of processes grouped by process state
  - NOTE: This may produce a LOT of output

- Optional branding "signature"
  - used to indicate what Nagios plugin (and what version) is responsible for
    the service check result

- Optional, leveled logging using `rs/zerolog` package
  - JSON-format output (to `stderr`)
  - choice of `disabled`, `panic`, `fatal`, `error`, `warn`, `info` (the
    default), `debug` or `trace`

NOTE: This tool ignores its own process entry when reporting running processes.

## Changelog

See the [`CHANGELOG.md`](CHANGELOG.md) file for the changes associated with
each release of this application. Changes that have been merged to `master`,
but not yet an official release may also be noted in the file under the
`Unreleased` section. A helpful link to the Git commit history since the last
official release is also provided for further review.

## Requirements

The following is a loose guideline. Other combinations of Go and operating
systems for building and running tools from this repo may work, but have not
been tested.

### Building source code

- Go
  - see this project's `go.mod` file for *preferred* version
  - this project tests against [officially supported Go
    releases][go-supported-releases]
    - the most recent stable release (aka, "stable")
    - the prior, but still supported release (aka, "oldstable")
- GCC
  - if building with custom options (as the provided `Makefile` does)
- `make`
  - if using the provided `Makefile`

### Running

- Red Hat Enterprise Linux 6
- Red Hat Enterprise Linux 7
- Red Hat Enterprise Linux 8
- Ubuntu 20.04

## Installation

### From source

1. [Download][go-docs-download] Go
1. [Install][go-docs-install] Go
1. Clone the repo
   1. `cd /tmp`
   1. `git clone https://github.com/atc0005/check-process`
   1. `cd check-process`
1. Install dependencies (optional)
   - for Ubuntu Linux
     - `sudo apt-get install make gcc`
   - for CentOS Linux
     1. `sudo yum install make gcc`
1. Build
   - manually, explicitly specifying target OS and architecture
     - `GOOS=linux GOARCH=amd64 go build -mod=vendor ./cmd/check_process/`
     - `GOOS=linux GOARCH=amd64 go build -mod=vendor ./cmd/lsps/`
       - most likely this is what you want (if building manually)
       - substitute `amd64` with the appropriate architecture if using
         different hardware (e.g., `arm64` or `386`)
   - using Makefile `linux` recipe
     - `make linux`
       - generates x86 and x64 binaries
   - using Makefile `release-build` recipe
     - `make release-build`
       - generates the same release assets as provided by this project's
         releases
1. Locate generated binaries
   - if using `Makefile`
     - look in `/tmp/check-process/release_assets/check_process/`
     - look in `/tmp/check-process/release_assets/lsps/`
   - if using `go build`
     - look in `/tmp/check-process/`
1. Copy the applicable binaries to whatever systems needs to run them so that
   they can be deployed

**NOTE**: Depending on which `Makefile` recipe you use the generated binary
may be compressed and have an `xz` extension. If so, you should decompress the
binary first before deploying it (e.g., `xz -d check_process-linux-amd64.xz`).

### Using release binaries

1. Download the [latest release][repo-url] binaries
1. Decompress binaries
   - e.g., `xz -d check_process-linux-amd64.xz`
1. Copy the applicable binaries to whatever systems needs to run them so that
   they can be deployed

**NOTE**:

DEB and RPM packages are provided as an alternative to manually deploying
binaries.

### Deployment

1. Place `check_process` in a location where it can be executed by the
   monitoring agent
   - Usually the same place as other Nagios plugins
   - For example, on a default Red Hat Enterprise Linux system using
   `check_nrpe` the `check_process` plugin would be deployed to
   `/usr/lib64/nagios/plugins/check_process` or
   `/usr/local/nagios/libexec/check_process`
1. Place `lsps` in a location where it can be easily accessed
   - Usually the same place as other custom tools installed outside of your
     package manager's control
   - e.g., `/usr/local/bin/lsps`

**NOTE**:

DEB and RPM packages are provided as an alternative to manually deploying
binaries.

## Configuration

### Command-line arguments

- Use the `-h` or `--help` flag to display current usage information.
- Flags marked as **`required`** must be set via CLI flag.
- Flags *not* marked as required are for settings where a useful default is
  already defined, but may be overridden if desired.

#### `check_process`

| Flag              | Required | Default | Repeat | Possible                                                                | Description                                                                                          |
| ----------------- | -------- | ------- | ------ | ----------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `branding`        | No       | `false` | No     | `branding`                                                              | Toggles emission of branding details with plugin status details. This output is disabled by default. |
| `h`, `help`       | No       | `false` | No     | `h`, `help`                                                             | Show Help text along with the list of supported flags.                                               |
| `version`         | No       | `false` | No     | `version`                                                               | Whether to display application version and then immediately exit application.                        |
| `ll`, `log-level` | No       | `info`  | No     | `disabled`, `panic`, `fatal`, `error`, `warn`, `info`, `debug`, `trace` | Log message priority filter. Log messages with a lower level are ignored.                            |

#### `lsps`

| Flag              | Required | Default | Repeat | Possible                                                                | Description                                                                                          |
| ----------------- | -------- | ------- | ------ | ----------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `branding`        | No       | `false` | No     | `branding`                                                              | Toggles emission of branding details with plugin status details. This output is disabled by default. |
| `h`, `help`       | No       | `false` | No     | `h`, `help`                                                             | Show Help text along with the list of supported flags.                                               |
| `version`         | No       | `false` | No     | `version`                                                               | Whether to display application version and then immediately exit application.                        |
| `show-all`        | No       | `false` | No     | `show-all`                                                              | Toggles listing of all processes. WARNING: This may produce a LOT of output. Disabled by default.    |
| `ll`, `log-level` | No       | `info`  | No     | `disabled`, `panic`, `fatal`, `error`, `warn`, `info`, `debug`, `trace` | Log message priority filter. Log messages with a lower level are ignored.                            |

## Process states

Red Hat Enterprise Linux 6 running a 2.6.32 version kernel is the baseline
test environment for this project.

The valid process states for a 2.6.32 kernel differs from the process states
for a 3.10 kernel (RHEL 7) which in turn differs from a 4.18 (RHEL 8) and
newer kernel. This project attempts to evaluate processes in all supported
states. In an effort to simplify use, some assumptions are made regarding
which process states map to which monitoring plugin state.

### Known states

The state details in this section were pulled directly from the source code
for each of the upstream kernel versions for RHEL releases that this project
was tested against. See the [References](#references) section for additional
details.

#### v2.6.32 (RHEL 6)

```c
/*
 * The task state array is a strange "bitmap" of
 * reasons to sleep. Thus "running" is zero, and
 * you can test for combinations of others with
 * simple bit tests.
 */
static const char *task_state_array[] = {
  "R (running)",         /*  0 */
  "S (sleeping)",        /*  1 */
  "D (disk sleep)",      /*  2 */
  "T (stopped)",         /*  4 */
  "T (tracing stop)",    /*  8 */
  "Z (zombie)",          /* 16 */
  "X (dead)"             /* 32 */
};
```

#### v3.10 (RHEL 7)

```c
/*
 * The task state array is a strange "bitmap" of
 * reasons to sleep. Thus "running" is zero, and
 * you can test for combinations of others with
 * simple bit tests.
 */
static const char * const task_state_array[] = {
  "R (running)",         /*   0 */
  "S (sleeping)",        /*   1 */
  "D (disk sleep)",      /*   2 */
  "T (stopped)",         /*   4 */
  "t (tracing stop)",    /*   8 */
  "Z (zombie)",          /*  16 */
  "X (dead)",            /*  32 */
  "x (dead)",            /*  64 */
  "K (wakekill)",        /* 128 */
  "W (waking)",          /* 256 */
  "P (parked)",          /* 512 */
};

```

#### v4.18 (RHEL 8)

```c
/*
 * The task state array is a strange "bitmap" of
 * reasons to sleep. Thus "running" is zero, and
 * you can test for combinations of others with
 * simple bit tests.
 */
static const char * const task_state_array[] = {

  /* states in TASK_REPORT: */
  "R (running)",         /* 0x00 */
  "S (sleeping)",        /* 0x01 */
  "D (disk sleep)",      /* 0x02 */
  "T (stopped)",         /* 0x04 */
  "t (tracing stop)",    /* 0x08 */
  "X (dead)",            /* 0x10 */
  "Z (zombie)",          /* 0x20 */
  "P (parked)",          /* 0x40 */

  /* states beyond TASK_REPORT: */
  "I (idle)",            /* 0x80 */
};
```

#### v.5.14 (RHEL 9)

```c
/*
 * The task state array is a strange "bitmap" of
 * reasons to sleep. Thus "running" is zero, and
 * you can test for combinations of others with
 * simple bit tests.
 */
static const char * const task_state_array[] = {

  /* states in TASK_REPORT: */
  "R (running)",         /* 0x00 */
  "S (sleeping)",        /* 0x01 */
  "D (disk sleep)",      /* 0x02 */
  "T (stopped)",         /* 0x04 */
  "t (tracing stop)",    /* 0x08 */
  "X (dead)",            /* 0x10 */
  "Z (zombie)",          /* 0x20 */
  "P (parked)",          /* 0x40 */

  /* states beyond TASK_REPORT: */
  "I (idle)",            /* 0x80 */
};
```

#### Summary

```text
// kernel 2.6.32 (RHEL 6)
  "R (running)"
  "S (sleeping)"
  "D (disk sleep)"
  "T (stopped)"
  "T (tracing stop)"
  "Z (zombie)"
  "X (dead)"

// kernel 3.10 (RHEL 7)
  "R (running)"
  "S (sleeping)"
  "D (disk sleep)"
  "T (stopped)"
  "t (tracing stop)"
  "Z (zombie)"
  "X (dead)"
  "x (dead)"
  "K (wakekill)"
  "W (waking)"
  "P (parked)"

// kernel 4.18/5.14 (RHEL 8/9)
  "R (running)"
  "S (sleeping)"
  "D (disk sleep)"
  "T (stopped)"
  "t (tracing stop)"
  "X (dead)"
  "Z (zombie)"
  "P (parked)"
  "I (idle)"
```

### Process state to plugin state mappings

| Process State    | Monitoring State |
| ---------------- | ---------------- |
| `D (disk sleep)` | CRITICAL         |
| `Z (zombie)`     | WARNING          |

## Examples

### `OK` result

This output is emitted by the plugin when no problematic processes are found.

```console
$ ./check_process
OK: No problematic processes found (364 evaluated)

Process Summary:

  - R (running) [1]
  - S (sleeping) [363]


--------------------------------------------------


Problems:

  - None

 | 'dead'=0;;;; 'idle'=0;;;; 'parked'=0;;;; 'problem_processes'=0;;;; 'running'=1;;;; 'sleeping'=363;;;; 'stopped'=0;;;; 'time'=18ms;;;; 'tracing_stop'=0;;;; 'uninterruptible_disk_sleep'=0;;;; 'wakekill'=0;;;; 'waking'=0;;;; 'zombie'=0;;;;
```

Regarding the output:

- The last line beginning with a space and the `|` symbol are performance
  data metrics emitted by the plugin. Depending on your monitoring system, these
  metrics may be collected and exposed as graphs/charts.
- This output was captured on a Red Hat Enterprise Linux 6 system (baseline OS
  for testing). The output is comparable to other Linux distros.

### `WARNING` result

This output is emitted by the plugin when problematic processes of a WARNING
state are found.

TODO: Provide example output when this scenario is encountered.

### `CRITICAL` result

This output is emitted by the plugin when problematic processes of a CRITICAL
state are found.

In the case of the `rsync` entries below, the activity is fairly normal for
this system (daily, early AM backups). To work around this, you can either
modify the timeperiod used for notifications to exclude this scenario (until
`D` state processes are found outside of that window) or increase the number
of retries so that an alert is not raised until after all retry attempts have
been exceeded.

```console
$ ./check_process
CRITICAL: 2 problematic processes found (D (disk sleep) [2], R (running) [7], S (sleeping) [368], evaluated [377])

Process Summary:

- D (disk sleep) [2]
- R (running) [7]
- S (sleeping) [368]


--------------------------------------------------


Problems:
- Name: rsync [Parent: backup.sh (6761), State: D (disk sleep), Pid: 16431, PPid: 6761, Threads: 1]
- Name: rsync [Parent: backup.sh (6761), State: D (disk sleep), Pid: 18321, PPid: 6761, Threads: 1]

 | 'dead'=0;;;; 'idle'=0;;;; 'parked'=0;;;; 'problem_processes'=0;;;; 'running'=7;;;; 'sleeping'=368;;;; 'stopped'=0;;;; 'time'=18ms;;;; 'tracing_stop'=0;;;; 'uninterruptible_disk_sleep'=2;;;; 'wakekill'=0;;;; 'waking'=0;;;; 'zombie'=0;;;;
```

## License

See the [LICENSE](LICENSE) file for details.

## References

- proc filesystem (usually mounted at `/proc`)
  - <https://man7.org/linux/man-pages/man5/proc.5.html>
  - <https://linux.die.net/man/5/proc>
  - <https://stackoverflow.com/questions/39066998/what-are-the-meaning-of-values-at-proc-pid-stat>
- valid process states
  - <https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/fs/proc/array.c?h=v2.6.32#n136>
  - <https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/fs/proc/array.c?h=v3.10#n135>
  - <https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/fs/proc/array.c?h=v4.18#n130>
  - <https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/fs/proc/array.c?h=v5.14#n130>

<!-- Footnotes here  -->

[repo-url]: <https://github.com/atc0005/check-process>  "This project's GitHub repo"

[go-docs-download]: <https://golang.org/dl>  "Download Go"

[go-docs-install]: <https://golang.org/doc/install>  "Install Go"

[go-supported-releases]: <https://go.dev/doc/devel/release#policy> "Go Release Policy"
