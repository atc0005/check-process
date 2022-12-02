// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/check-process
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package process

// SupportedProcessStates provides a list of the process states supported by
// this package. Not all process states are supported by all kernel versions.
func SupportedProcessStates() []string {
	return []string{
		KernelAnyProcessStateRunning,
		KernelAnyProcessStateSleeping,
		KernelAnyProcessStateDiskSleep,
		KernelAnyProcessStateStopped,
		KernelAnyProcessStateZombie,
		KernelAnyProcessStateDead,
		KernelLegacyProcessStateTracingStop,
		KernelLegacyProcessStateDead,
		KernelLegacyProcessStateWakeKill,
		KernelLegacyProcessStateWaking,
		KernelCurrentProcessStateTracingStop,
		KernelCurrentProcessStateIdle,
		KernelCurrentProcessStateParked,
	}
}

// KnownProblemProcessStates provides a list of the process states known to be
// problematic.
func KnownProblemProcessStates() []string {
	return append(
		KnownCriticalProcessStates(),
		KnownWarningProcessStates()...,
	)
}

// KnownWarningProcessStates provides a list of the process states known to be
// problematic and warrant a WARNING, but not so severe to be considered
// CRITICAL.
func KnownWarningProcessStates() []string {
	return []string{
		// KernelAnyProcessStateDiskSleep,
		KernelAnyProcessStateZombie,

		// TODO: Do any of these belong here?
		//
		// KernelAnyProcessStateDead,
		// KernelAnyProcessStateStopped,
		// KernelLegacyProcessStateTracingStop,
		// KernelLegacyProcessStateDead,
		// KernelLegacyProcessStateWakeKill,
		// KernelLegacyProcessStateWaking,
		// KernelCurrentProcessStateTracingStop,
		// KernelCurrentProcessStateParked,
	}
}

// KnownCriticalProcessStates returns a list of the process states known to be
// problematic. These states are considered CRITICAL.
func KnownCriticalProcessStates() []string {
	return []string{
		KernelAnyProcessStateDiskSleep,
		// KernelAnyProcessStateZombie,

		// TODO: Do any of these belong here?
		//
		// KernelAnyProcessStateDead,
		// KernelAnyProcessStateStopped,
		// KernelLegacyProcessStateTracingStop,
		// KernelLegacyProcessStateDead,
		// KernelLegacyProcessStateWakeKill,
		// KernelLegacyProcessStateWaking,
		// KernelCurrentProcessStateTracingStop,
		// KernelCurrentProcessStateParked,
	}
}
