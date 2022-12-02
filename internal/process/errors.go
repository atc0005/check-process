// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/check-process
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package process

import "errors"

var (
	// ErrProblemProcessesFound indicates that one or more "problematic"
	// processes were found (e.g., uninterruptible disk sleep) which warrant
	// surfacing the issue.
	ErrProblemProcessesFound = errors.New("problematic processes found")

	// ErrInvalidProcStatusLineFormat indicates that a status file within the
	// /proc filesystem has an invalid key/value format.
	ErrInvalidProcStatusLineFormat = errors.New("invalid format for proc status line")

	// ErrInvalidProcStatusLineValue indicates that a status file within the
	// /proc filesystem has an invalid value (wrong type) for a process
	// property.
	ErrInvalidProcStatusLineValue = errors.New("invalid value for proc status line")

	// ErrMissingProcessPropertiesIndexEntry indicates that a requested
	// process property entry is missing from the index.
	ErrMissingProcessPropertiesIndexEntry = errors.New("entry missing from process properties index")

	// ErrInvalidProcessPropertiesIndexCount indicates that a
	// ProcessProperties index (map) is missing the minimum required entries
	// used to populate specific fields of a Process.
	ErrInvalidProcessPropertiesIndexCount = errors.New("minimum entries not present in process properties index")

	// ErrParentProcessResolveFailure indicates that a failure occurred trying
	// to determine the parent process
	// ErrParentProcessResolveFailure = errors.New("unable to determine parent process")

	// ErrMissingProcessEntry indicates that an expected process was not found
	// in a given collection.
	ErrMissingProcessEntry = errors.New("process missing from collection")
)
