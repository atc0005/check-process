// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/check-process
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package process

import (
	"fmt"
)

// Process provides select type-safe values related to a running process (per
// /proc) and a map of property/values in string format.
//
// https://man7.org/linux/man-pages/man5/proc.5.html
// https://linux.die.net/man/5/proc
type Process struct {
	Name string

	// State is the current state of the process.
	//
	// Valid values for a 2.6.32 (RHEL 6) kernel:
	//
	// "R (running)"
	// "S (sleeping)"
	// "D (disk sleep)"
	// "T (stopped)"
	// "T (tracing stop)"
	// "Z (zombie)"
	// "X (dead)"
	//
	// Valid values for a 3.10 (RHEL 7) kernel:
	//
	// "R (running)"
	// "S (sleeping)"
	// "D (disk sleep)"
	// "T (stopped)"
	// "t (tracing stop)"
	// "Z (zombie)"
	// "X (dead)"
	// "x (dead)"
	// "K (wakekill)"
	// "W (waking)"
	// "P (parked)"
	//
	//
	// Valid values for a 4.18 (RHEL 8) and 5.14 (RHEL 9) kernel:
	//
	// "R (running)"
	// "S (sleeping)"
	// "D (disk sleep)"
	// "T (stopped)"
	// "t (tracing stop)"
	// "X (dead)"
	// "Z (zombie)"
	// "P (parked)"
	// "I (idle)"
	//
	// https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/fs/proc/array.c?h=v2.6.32#n136
	// https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/fs/proc/array.c?h=v3.10#n135
	// https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/fs/proc/array.c?h=v4.18#n130
	// https://git.kernel.org/pub/scm/linux/kernel/git/torvalds/linux.git/tree/fs/proc/array.c?h=v5.14#n130
	State         string
	Pid           int
	PPid          int
	Threads       int
	VMSwap        string
	AllProperties Properties
}

// Properties is a collection of key/value string pairs representing
// the various properties for a running process. Each key is lowercased with
// the original value retained as-is.
type Properties map[string]string

// ParentProcess returns the parent Process for the current Process value from
// the specified collection or an error if one occurs.
func (p Process) ParentProcess(processes Processes) (Process, error) {
	parentID := p.PPid
	for _, process := range processes {
		if process.Pid == parentID {
			return process, nil
		}
	}

	return Process{}, fmt.Errorf(
		"failed to resolve parent process: %w",
		ErrMissingProcessEntry,
	)
}

// IsOKState indicates whether the process state is not in a list of known
// problematic process states.
func (p Process) IsOKState() bool {
	for _, probState := range KnownProblemProcessStates() {
		if p.State == probState {
			return false
		}
	}

	return true
}

// IsWarningState indicates whether the state matches known WARNING severity
// process states.
func (p Process) IsWarningState() bool {
	for _, probState := range KnownWarningProcessStates() {
		if p.State == probState {
			return true
		}
	}

	return false
}

// IsCriticalState indicates whether the state matches known CRITICAL severity
// process states.
func (p Process) IsCriticalState() bool {
	for _, probState := range KnownCriticalProcessStates() {
		if p.State == probState {
			return true
		}
	}

	return false
}
