// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/check-process
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package process

const (

	// ProcRootDir is the default mount point for the proc virtual filesystem.
	ProcRootDir string = "/proc"

	// ProcStatusFilename is the name of the human-readable process status
	// information present for each process directory in the proc filesystem.
	ProcStatusFilename string = "status"

	// ProcDirRegex is the regex pattern used to match process directory names
	// within the proc virtual filesystem.
	ProcDirRegex string = "^[0-9]+$"
)

// Process status field names.
const (
	ProcessNameField    string = "name"
	ProcessStateField   string = "state"
	ProcessPidField     string = "pid"
	ProcessPPidField    string = "ppid"
	ProcessThreadsField string = "threads"
	ProcessVMSwapField  string = "vmswap"
)

// Process state values observed for 2.6.32, 3.10, 4.18 & 5.14 kernels and
// presumed to be stable enough to be used by future kernels.
const (
	KernelAnyProcessStateRunning   string = "R (running)"
	KernelAnyProcessStateSleeping  string = "S (sleeping)"
	KernelAnyProcessStateDiskSleep string = "D (disk sleep)"
	KernelAnyProcessStateStopped   string = "T (stopped)"
	KernelAnyProcessStateZombie    string = "Z (zombie)"
	KernelAnyProcessStateDead      string = "X (dead)"
)

// Process state values observed for legacy kernels that are not present on
// current/modern kernel versions.
const (
	KernelLegacyProcessStateTracingStop string = "T (tracing stop)" // kernel 2.6.32
	KernelLegacyProcessStateDead        string = "x (dead)"         // kernel 3.10
	KernelLegacyProcessStateWakeKill    string = "K (wakekill)"     // kernel 3.10
	KernelLegacyProcessStateWaking      string = "W (waking)"       // kernel 3.10
)

// Process state values observed for newer kernel versions and presumed to be
// stable enough to be used by future kernels.
const (
	KernelCurrentProcessStateTracingStop string = "t (tracing stop)" // kernel 3.10, 4.18, 5.14
	KernelCurrentProcessStateIdle        string = "I (idle)"         // kernel 4.18, 5.14
	KernelCurrentProcessStateParked      string = "P (parked)"       // kernel 3.10, 4.18, 5.14
)
