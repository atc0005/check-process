// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/check-process
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/atc0005/check-process/internal/config"
	"github.com/atc0005/check-process/internal/process"
	"github.com/atc0005/check-process/internal/reports"
	"github.com/atc0005/go-nagios"
	"github.com/rs/zerolog"
)

func main() {
	// Start the timer. We'll use this to emit the plugin runtime as a
	// performance data metric.
	pluginStart := time.Now()

	// Set initial "state" as valid, adjust as we go.
	var nagiosExitState = nagios.ExitState{
		LastError:      nil,
		ExitStatusCode: nagios.StateOKExitCode,
	}

	// defer this from the start so it is the last deferred function to run
	defer nagiosExitState.ReturnCheckResults()

	// Setup configuration by parsing user-provided flags.
	cfg, cfgErr := config.New(config.AppType{Plugin: true})
	switch {
	case errors.Is(cfgErr, config.ErrVersionRequested):
		fmt.Println(config.Version())

		return

	case errors.Is(cfgErr, config.ErrHelpRequested):
		fmt.Println(cfg.Help())

		return

	case cfgErr != nil:

		// We make some assumptions when setting up our logger as we do not
		// have a working configuration based on sysadmin-specified choices.
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr, NoColor: true}
		logger := zerolog.New(consoleWriter).With().Timestamp().Caller().Logger()

		logger.Err(cfgErr).Msg("Error initializing application")

		nagiosExitState.ServiceOutput = fmt.Sprintf(
			"%s: Error initializing application",
			nagios.StateCRITICALLabel,
		)
		nagiosExitState.AddError(cfgErr)
		nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode

		return
	}

	// Collect last minute details just before ending plugin execution.
	defer appendPerfData(&nagiosExitState, pluginStart, cfg.Log)

	if cfg.EmitBranding {
		// If enabled, show application details at end of notification
		nagiosExitState.BrandingCallback = config.Branding("Notification generated by ")
	}

	logger := cfg.Log.With().Logger()

	logger.Debug().
		Str("base_path", process.ProcRootDir).
		Msg("Collecting process paths")
	procDirs, err := process.GetProcDirs(process.ProcRootDir)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to evaluate process directories")

		nagiosExitState.AddError(err)
		nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
		nagiosExitState.ServiceOutput = fmt.Sprintf(
			"%s: Failed to evaluate process directories",
			nagios.StateCRITICALLabel,
		)

		return
	}

	logger.Debug().
		Str("base_path", process.ProcRootDir).
		Int("process_paths", len(procDirs)).
		Msg("Successfully collected process paths")

	processes, err := process.FromProcDirs(procDirs)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to obtain list of process values")

		nagiosExitState.AddError(err)
		nagiosExitState.ExitStatusCode = nagios.StateCRITICALExitCode
		nagiosExitState.ServiceOutput = fmt.Sprintf(
			"%s: Failed to process proc status files",
			nagios.StateCRITICALLabel,
		)

		return
	}

	logger.Debug().
		Int("processes", len(processes)).
		Msg("Collected info on processes")

	logger.Debug().
		Str("my_name", os.Args[0]).
		Int("my_process_id", os.Getpid()).
		Msg("Excluding process of current tool")
	processes = processes.ExcludeMyPID()

	logger.Debug().
		Int("processes", len(processes)).
		Msg("Excluded process of current tool")

	switch {
	case !processes.IsOKState():

		logger.Debug().
			Int("critical_processes", processes.NumCriticalState()).
			Int("warning_processes", processes.NumWarningState()).
			Int("ok_processes", processes.NumOKState()).
			Msg("Problematic processes found")

		nagiosExitState.AddError(process.ErrProblemProcessesFound)

		nagiosExitState.ExitStatusCode = processes.ServiceState().ExitCode

		nagiosExitState.ServiceOutput = reports.CheckProcessOneLineSummary(processes)
		nagiosExitState.LongServiceOutput = reports.CheckProcessReport(processes)

		pd := getPerfData(processes)
		if err := nagiosExitState.AddPerfData(false, pd...); err != nil {
			logger.Error().
				Err(err).
				Msg("failed to add performance data")
		}

		return

	default:

		logger.Debug().Msg("No problematic processes detected")

		nagiosExitState.ServiceOutput = reports.CheckProcessOneLineSummary(processes)
		nagiosExitState.LongServiceOutput = reports.CheckProcessReport(processes)

		nagiosExitState.ExitStatusCode = processes.ServiceState().ExitCode

		pd := getPerfData(processes)
		if err := nagiosExitState.AddPerfData(false, pd...); err != nil {
			logger.Error().
				Err(err).
				Msg("failed to add performance data")
		}

		return

	}

}
