// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/check-process
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fmt"
	"time"

	"github.com/atc0005/check-process/internal/process"
	"github.com/atc0005/go-nagios"
	"github.com/rs/zerolog"
)

func appendPerfData(exitState *nagios.ExitState, start time.Time, logger zerolog.Logger) {
	// Record plugin runtime, emit this metric regardless of exit
	// point/cause.
	runtimeMetric := nagios.PerformanceData{
		Label: "time",
		Value: fmt.Sprintf("%dms", time.Since(start).Milliseconds()),
	}
	if err := exitState.AddPerfData(false, runtimeMetric); err != nil {
		logger.Error().
			Err(err).
			Msg("failed to add time (runtime) performance data metric")
	}

}

// getPerfData gathers performance data metrics that we wish to report.
func getPerfData(processes process.Processes) []nagios.PerformanceData {

	probProcs := processes.States(process.KnownProblemProcessStates())

	return []nagios.PerformanceData{
		// The `time` (runtime) metric is appended at plugin exit, so do not
		// duplicate it here.
		{
			Label: "problem_processes",
			Value: fmt.Sprintf("%d", len(probProcs)),
		},
		{
			Label: "running",
			Value: fmt.Sprintf("%d", processes.StateRunningCount()),
		},
		{
			Label: "sleeping",
			Value: fmt.Sprintf("%d", processes.StateSleepingCount()),
		},
		{
			Label: "uninterruptible_disk_sleep",
			Value: fmt.Sprintf("%d", processes.StateDiskSleepCount()),
		},
		{
			Label: "stopped",
			Value: fmt.Sprintf("%d", processes.StateStoppedCount()),
		},
		{
			Label: "zombie",
			Value: fmt.Sprintf("%d", processes.StateZombieCount()),
		},
		{
			Label: "dead",
			Value: fmt.Sprintf("%d", processes.StateDeadCount()),
		},
		{
			Label: "tracing_stop",
			Value: fmt.Sprintf("%d", processes.StateTracingStopCount()),
		},
		{
			Label: "wakekill",
			Value: fmt.Sprintf("%d", processes.StateWakeKillCount()),
		},
		{
			Label: "waking",
			Value: fmt.Sprintf("%d", processes.StateWakingCount()),
		},
		{
			Label: "idle",
			Value: fmt.Sprintf("%d", processes.StateIdleCount()),
		},
		{
			Label: "parked",
			Value: fmt.Sprintf("%d", processes.StateParkedCount()),
		},
	}
}
