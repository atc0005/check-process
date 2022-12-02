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
	"path/filepath"

	"github.com/atc0005/check-process/internal/config"
	"github.com/atc0005/check-process/internal/process"
	"github.com/rs/zerolog"
)

func main() {

	// Setup configuration by parsing user-provided flags
	cfg, cfgErr := config.New(config.AppType{Inspector: true})
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
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr}
		logger := zerolog.New(consoleWriter).With().Timestamp().Caller().Logger()

		logger.Err(cfgErr).Msg("Error initializing application")

		return
	}

	logger := cfg.Log.With().Logger()

	logger.Debug().
		Str("base_path", process.ProcRootDir).
		Msg("Collecting process paths")
	procDirs, err := process.GetProcDirs(process.ProcRootDir)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to evaluate process directories")
		os.Exit(config.ExitCodeCatchall)
	}

	logger.Debug().
		Str("base_path", process.ProcRootDir).
		Int("process_paths", len(procDirs)).
		Msg("Successfully collected process paths")

	processes := make(process.Processes, 0, len(procDirs))
	for _, procDir := range procDirs {

		qualifiedPath := filepath.Join(process.ProcRootDir, procDir, process.ProcStatusFilename)
		p, err := process.ParseProcStatusFile(qualifiedPath)
		if err != nil {
			logger.Error().
				Err(err).
				Str("file", qualifiedPath).
				Msg("Failed to process file")
			os.Exit(config.ExitCodeCatchall)
		}

		processes = append(processes, p)

	}

	logger.Debug().
		Int("processes", len(processes)).
		Msg("Collected info on processes")

	fmt.Println("Problematic processes:")

	probProcs := processes.States(process.KnownProblemProcessStates())
	listProcesses(os.Stdout, probProcs)
	listOtherProcesses(os.Stdout, probProcs, processes, cfg.InspectorSettings.ShowAll)
}
