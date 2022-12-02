// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/check-process
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package process

import (
	"fmt"
	"sort"
	"strings"

	"github.com/atc0005/go-nagios"
)

// Processes is a collection of Process values.
type Processes []Process

// ParentProcess returns the parent Process for the specified Process value
// from the collection or an error if one occurs.
func (ps Processes) ParentProcess(p Process) (Process, error) {
	parentID := p.PPid
	for _, processEntry := range ps {
		if processEntry.Pid == parentID {
			return processEntry, nil
		}
	}

	return Process{}, fmt.Errorf(
		"failed to resolve parent process: %w",
		ErrMissingProcessEntry,
	)
}

// Exclude returns Process values from the collection which are not in the
// specified set. If the specified set is empty all Process values in the
// collection are returned. If the
func (ps Processes) Exclude(exclude Processes) Processes {

	// Build index of ID values the set we are to exclude.
	excludePids := make(map[int]struct{}, len(exclude))
	for _, p := range exclude {
		excludePids[p.Pid] = struct{}{}
	}

	// Skip further processing if we've been asked to exclude the full
	// collection (however unlikely).
	if len(excludePids) == len(ps) {
		return Processes{}
	}

	remaining := make(Processes, 0)
	for _, p := range ps {
		_, excluded := excludePids[p.Pid]
		if !excluded {
			remaining = append(remaining, p)
		}
	}

	return remaining
}

// StateCount is a helper method used to indicate how many processes in the
// collection are in the specified state. The caller is encouraged to specify
// state values using applicable state value constants for best results.
func (ps Processes) StateCount(state string) int {
	var ctr int
	for _, p := range ps {
		if p.State == state {
			ctr++
		}
	}
	return ctr
}

// Count returns the number of Process entries in the collection.
func (ps Processes) Count() int {
	return len(ps)
}

// State returns each Process from the collection that are in the specified
// state. The caller is encouraged to specify state values using applicable
// state value constants for best results. The returned collection may be
// empty if no processes are in the requested state.
func (ps Processes) State(state string) Processes {
	num := ps.StateCount(state)

	// Early exit if there are not any processes to report.
	if num == 0 {
		return Processes{}
	}

	processes := make(Processes, 0, num)
	for _, p := range ps {
		if p.State == state {
			processes = append(processes, p)
		}
	}

	return processes
}

// States returns each Process from the collection that are in the specified
// states. The caller is encouraged to specify state values using applicable
// state value constants for best results. The returned collection is empty if
// no processes are in the requested states.
func (ps Processes) States(states []string) Processes {

	// If an empty set of states was specified, an empty collection is
	// provided.
	if len(states) == 0 {
		return Processes{}
	}

	var num int
	for _, state := range states {
		num += ps.StateCount(state)
	}

	// Early exit if there are not any processes to report.
	if num == 0 {
		return Processes{}
	}

	processes := make(Processes, 0, num)
	for _, state := range states {
		for _, p := range ps {
			if p.State == state {
				processes = append(processes, p)
			}
		}
	}

	return processes
}

// StateRunning returns each Process from the collection currently in a
// running state. The returned collection may be empty if no processes are in
// the requested state.
func (ps Processes) StateRunning() Processes {
	num := ps.StateCount(KernelAnyProcessStateRunning)

	// Early exit if there are not any processes to report.
	if num == 0 {
		return Processes{}
	}

	processes := make(Processes, 0, num)
	for _, p := range ps {
		if p.State == KernelAnyProcessStateRunning {
			processes = append(processes, p)
		}
	}

	return processes
}

// StateRunningCount returns the number of processes from the collection
// currently in a running state.
func (ps Processes) StateRunningCount() int {
	num := ps.StateCount(KernelAnyProcessStateRunning)

	return num
}

// StateSleeping returns each Process from the collection currently in a
// sleeping state. The returned collection may be empty if no processes are in
// the requested state.
func (ps Processes) StateSleeping() Processes {
	num := ps.StateCount(KernelAnyProcessStateSleeping)

	// Early exit if there are not any processes to report.
	if num == 0 {
		return Processes{}
	}

	processes := make(Processes, 0, num)
	for _, p := range ps {
		if p.State == KernelAnyProcessStateSleeping {
			processes = append(processes, p)
		}
	}

	return processes
}

// StateSleepingCount returns the number of processes from the collection
// currently in a sleeping state.
func (ps Processes) StateSleepingCount() int {
	num := ps.StateCount(KernelAnyProcessStateSleeping)

	return num
}

// StateDiskSleep returns each Process from the collection currently in a disk
// sleep state. The returned collection may be empty if no processes are in
// the requested state.
func (ps Processes) StateDiskSleep() Processes {
	num := ps.StateCount(KernelAnyProcessStateDiskSleep)

	// Early exit if there are not any processes to report.
	if num == 0 {
		return Processes{}
	}

	processes := make(Processes, 0, num)
	for _, p := range ps {
		if p.State == KernelAnyProcessStateDiskSleep {
			processes = append(processes, p)
		}
	}

	return processes
}

// StateDiskSleepCount returns the number of processes from the collection
// currently in a disk sleep state.
func (ps Processes) StateDiskSleepCount() int {
	num := ps.StateCount(KernelAnyProcessStateDiskSleep)

	return num
}

// StateStopped returns each Process from the collection currently in a
// stopped state. Processes in a "tracing stop" state are not evaluated.
//
// The returned collection may be empty if no processes are in the requested
// state.
func (ps Processes) StateStopped() Processes {
	num := ps.StateCount(KernelAnyProcessStateStopped)

	// Early exit if there are not any processes to report.
	if num == 0 {
		return Processes{}
	}

	processes := make(Processes, 0, num)
	for _, p := range ps {
		if p.State == KernelAnyProcessStateStopped {
			processes = append(processes, p)
		}
	}

	return processes
}

// StateStoppedCount returns the number of processes from the collection
// currently in a stopped state.
func (ps Processes) StateStoppedCount() int {
	num := ps.StateCount(KernelAnyProcessStateStopped)

	return num
}

// StateTracingStop returns each Process from the collection currently in a
// tracing stop state. The returned collection may be empty if no processes
// are in the requested state.
func (ps Processes) StateTracingStop() Processes {
	num := ps.StateCount(KernelCurrentProcessStateTracingStop)
	num += ps.StateCount(KernelLegacyProcessStateTracingStop)

	// Early exit if there are not any processes to report.
	if num == 0 {
		return Processes{}
	}

	processes := make(Processes, 0, num)
	for _, p := range ps {
		if p.State == KernelCurrentProcessStateTracingStop ||
			p.State == KernelLegacyProcessStateTracingStop {
			processes = append(processes, p)
		}
	}

	return processes
}

// StateTracingStopCount returns the number of processes from the collection
// currently in a tracing stop state.
func (ps Processes) StateTracingStopCount() int {
	num := ps.StateCount(KernelCurrentProcessStateTracingStop)
	num += ps.StateCount(KernelLegacyProcessStateTracingStop)

	return num
}

// StateZombie returns each Process from the collection currently in a zombie
// state. The returned collection may be empty if no processes are in the
// requested state.
func (ps Processes) StateZombie() Processes {
	num := ps.StateCount(KernelAnyProcessStateZombie)

	// Early exit if there are not any processes to report.
	if num == 0 {
		return Processes{}
	}

	processes := make(Processes, 0, num)
	for _, p := range ps {
		if p.State == KernelAnyProcessStateZombie {
			processes = append(processes, p)
		}
	}

	return processes
}

// StateZombieCount returns the number of processes from the collection
// currently in a zombie state.
func (ps Processes) StateZombieCount() int {
	num := ps.StateCount(KernelAnyProcessStateZombie)

	return num
}

// StateDead returns each Process from the collection currently in a dead
// state. The returned collection may be empty if no processes are in the
// requested state.
func (ps Processes) StateDead() Processes {
	num := ps.StateCount(KernelAnyProcessStateDead)
	num += ps.StateCount(KernelLegacyProcessStateDead)

	// Early exit if there are not any processes to report.
	if num == 0 {
		return Processes{}
	}

	processes := make(Processes, 0, num)
	for _, p := range ps {
		if p.State == KernelAnyProcessStateDead ||
			p.State == KernelLegacyProcessStateDead {
			processes = append(processes, p)
		}
	}

	return processes
}

// StateDeadCount returns the number of processes from the collection
// currently in a dead state.
func (ps Processes) StateDeadCount() int {
	num := ps.StateCount(KernelAnyProcessStateDead)
	num += ps.StateCount(KernelLegacyProcessStateDead)

	return num
}

// StateWakeKill returns each Process from the collection currently in a
// wakekill state. The returned collection may be empty if no processes are in
// the requested state.
func (ps Processes) StateWakeKill() Processes {
	num := ps.StateCount(KernelLegacyProcessStateWakeKill)

	// Early exit if there are not any processes to report.
	if num == 0 {
		return Processes{}
	}

	processes := make(Processes, 0, num)
	for _, p := range ps {
		if p.State == KernelLegacyProcessStateWakeKill {
			processes = append(processes, p)
		}
	}

	return processes
}

// StateWakeKillCount returns the number of processes from the collection
// currently in a wakekill state.
func (ps Processes) StateWakeKillCount() int {
	num := ps.StateCount(KernelLegacyProcessStateWakeKill)

	return num
}

// StateWaking returns each Process from the collection currently in a
// waking state. The returned collection may be empty if no processes are in
// the requested state.
func (ps Processes) StateWaking() Processes {
	num := ps.StateCount(KernelLegacyProcessStateWaking)

	// Early exit if there are not any processes to report.
	if num == 0 {
		return Processes{}
	}

	processes := make(Processes, 0, num)
	for _, p := range ps {
		if p.State == KernelLegacyProcessStateWaking {
			processes = append(processes, p)
		}
	}

	return processes
}

// StateWakingCount returns the number of processes from the collection
// currently in a waking state.
func (ps Processes) StateWakingCount() int {
	num := ps.StateCount(KernelLegacyProcessStateWaking)

	return num
}

// StateParked returns each Process from the collection currently in a parked
// state. The returned collection may be empty if no processes are in the
// requested state.
func (ps Processes) StateParked() Processes {
	num := ps.StateCount(KernelCurrentProcessStateParked)

	// Early exit if there are not any processes to report.
	if num == 0 {
		return Processes{}
	}

	processes := make(Processes, 0, num)
	for _, p := range ps {
		if p.State == KernelCurrentProcessStateParked {
			processes = append(processes, p)
		}
	}

	return processes
}

// StateParkedCount returns the number of processes from the collection
// currently in a parked state.
func (ps Processes) StateParkedCount() int {
	num := ps.StateCount(KernelCurrentProcessStateParked)

	return num
}

// StateIdle returns each Process from the collection currently in an idle
// state. The returned collection may be empty if no processes are in the
// requested state.
func (ps Processes) StateIdle() Processes {
	num := ps.StateCount(KernelCurrentProcessStateIdle)

	// Early exit if there are not any processes to report.
	if num == 0 {
		return Processes{}
	}

	processes := make(Processes, 0, num)
	for _, p := range ps {
		if p.State == KernelCurrentProcessStateIdle {
			processes = append(processes, p)
		}
	}

	return processes
}

// StateIdleCount returns the number of processes from the collection
// currently in a idle state.
func (ps Processes) StateIdleCount() int {
	num := ps.StateCount(KernelCurrentProcessStateIdle)

	return num
}

// SummaryOneLine returns a one line summary of all processes.
func (ps Processes) SummaryOneLine() string {
	tally := make(map[string]int)
	for _, p := range ps {
		tally[p.State]++
	}

	keys := make([]string, 0, len(tally))
	for k := range tally {
		keys = append(keys, k)
	}
	sort.Strings(sort.StringSlice(keys))

	summaryItems := make([]string, 0, len(tally))

	for _, key := range keys {
		summaryItems = append(summaryItems, fmt.Sprintf(
			"%s [%d]",
			key, tally[key],
		))
	}

	summary := strings.Join(summaryItems, ", ")

	return summary
}

// SummaryList returns a slice of process summary items.
func (ps Processes) SummaryList() []string {
	tally := make(map[string]int)
	for _, p := range ps {
		tally[p.State]++
	}

	keys := make([]string, 0, len(tally))
	for k := range tally {
		keys = append(keys, k)
	}
	sort.Strings(sort.StringSlice(keys))

	summaryItems := make([]string, 0, len(tally))

	for _, key := range keys {
		summaryItems = append(summaryItems, fmt.Sprintf(
			"%s [%d]",
			key, tally[key],
		))
	}

	return summaryItems
}

// IsOKState indicates whether all items in the collection were evaluated to
// an OK state.
func (ps Processes) IsOKState() bool {
	for _, p := range ps {
		if !p.IsOKState() {
			return false
		}
	}

	return true
}

// NumOKState indicates how many items in the collection were evaluated to
// an OK state.
func (ps Processes) NumOKState() int {
	var ctr int
	for _, p := range ps {
		if p.IsOKState() {
			ctr++
		}
	}

	return ctr
}

// HasCriticalState indicates whether any items in the collection were
// evaluated to a CRITICAL state.
func (ps Processes) HasCriticalState() bool {
	for _, p := range ps {
		if p.IsCriticalState() {
			return true
		}
	}

	return false
}

// NumCriticalState indicates how many items in the collection were evaluated
// to a CRITICAL state.
func (ps Processes) NumCriticalState() int {
	var ctr int
	for _, p := range ps {
		if p.IsCriticalState() {
			ctr++
		}
	}

	return ctr
}

// HasWarningState indicates whether any items in the collection were
// evaluated to a WARNING state.
func (ps Processes) HasWarningState() bool {
	for _, p := range ps {
		if p.IsWarningState() {
			return true
		}
	}

	return false
}

// NumWarningState indicates how many items in the collection were evaluated
// to a WARNING state.
func (ps Processes) NumWarningState() int {
	var ctr int
	for _, p := range ps {
		if p.IsWarningState() {
			ctr++
		}
	}

	return ctr
}

// ServiceState returns the appropriate Service Check Status label and exit
// code for the collection's evaluation results.
func (ps Processes) ServiceState() nagios.ServiceState {
	var stateLabel string
	var stateExitCode int

	switch {
	case ps.HasCriticalState():
		stateLabel = nagios.StateCRITICALLabel
		stateExitCode = nagios.StateCRITICALExitCode
	case ps.HasWarningState():
		stateLabel = nagios.StateWARNINGLabel
		stateExitCode = nagios.StateWARNINGExitCode
	case ps.IsOKState():
		stateLabel = nagios.StateOKLabel
		stateExitCode = nagios.StateOKExitCode
	default:
		stateLabel = nagios.StateUNKNOWNLabel
		stateExitCode = nagios.StateUNKNOWNExitCode
	}

	return nagios.ServiceState{
		Label:    stateLabel,
		ExitCode: stateExitCode,
	}
}
