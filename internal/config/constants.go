// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/check-process
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package config

const myAppName string = "check-process"
const myAppURL string = "https://github.com/atc0005/check-process"

// ExitCodeCatchall indicates a general or miscellaneous error has occurred.
// This exit code is not directly used by monitoring plugins in this project.
// See https://tldp.org/LDP/abs/html/exitcodes.html for additional details.
const ExitCodeCatchall int = 1

// Shared flag help text
const (
	versionFlagHelp      string = "Whether to display application version and then immediately exit application."
	logLevelFlagHelp     string = "Sets log level."
	brandingFlagHelp     string = "Toggles emission of branding details with plugin status details. This output is disabled by default."
	emitBrandingFlagHelp string = "Toggles emission of branding details with plugin status details. This output is disabled by default."
	helpFlagHelp         string = "Emit this help text"
)

const (
	showAllProcessesFlagHelp string = "Toggles listing of all processes. WARNING: This may produce a LOT of output. Disabled by default."
)

// Flag names for consistent references. Exported so that they're available
// from tests.
const (
	HelpFlagLong             string = "help"
	HelpFlagShort            string = "h"
	VersionFlagLong          string = "version"
	BrandingFlag             string = "branding"
	TimeoutFlagLong          string = "timeout"
	TimeoutFlagShort         string = "t"
	LogLevelFlagLong         string = "log-level"
	LogLevelFlagShort        string = "ll"
	ShowAllProcessesFlagLong string = "show-all"
)

// Default flag settings if not overridden by user input
const (
	defaultHelp                  bool   = false
	defaultLogLevel              string = "info"
	defaultEmitBranding          bool   = false
	defaultDisplayVersionAndExit bool   = false
	defaultShowAllProcesses      bool   = false
)

const (
	appTypePlugin    string = "plugin"
	appTypeInspector string = "Inspector"
)
