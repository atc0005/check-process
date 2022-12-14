// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/check-process
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package process

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// getProcDirsCount is a small helper function to calculate the number of
// directories from a given collection with valid process (ID pattern) names.
func getProcDirsCount(dirEntries []fs.DirEntry, re *regexp.Regexp) int {
	var ctr int
	for _, path := range dirEntries {
		if path.IsDir() && re.MatchString(path.Name()) {
			ctr++
		}
	}
	return ctr
}

// GetProcDirs evaluates the given base path (usually "/proc") for known
// process directories within the proc filesystem and returns an unqualified
// list of all valid matches.
func GetProcDirs(path string) ([]string, error) {
	// Will panic if the regex fails to compile. Because the regex value is
	// fixed and unlikely to be modified this should be a safe decision.
	re := regexp.MustCompile(ProcDirRegex)

	dirEntries, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to evaluate processes using proc path %s: %w",
			path,
			err,
		)
	}

	procDirsCount := getProcDirsCount(dirEntries, re)
	procDirs := make([]string, 0, procDirsCount)
	for _, path := range dirEntries {
		if !path.IsDir() {
			continue
		}

		// If the subdirectory of /proc matches our regular expression
		if re.MatchString(path.Name()) {
			procDirs = append(procDirs, path.Name())
		}
	}

	return procDirs, nil
}

// ParseProcStatusFile parses a given /proc/[pid]/status file and returns a
// Process value representing the status of a process.
func ParseProcStatusFile(filename string) (Process, error) {
	data, err := os.ReadFile(filepath.Clean(filename))
	if err != nil {
		return Process{},
			fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	r := bytes.NewReader(data)
	p, err := getProcessProperties(r)
	if err != nil {
		return Process{}, fmt.Errorf(
			"failed to obtain properties for process from file %s: %w",
			filename,
			err,
		)
	}

	return p, nil
}

// getProcessProperties is responsible for processing a status file's content
// (exposed as an io.Reader) and generating a Process value that represents
// common values (e.g., Name, State, Pid) and a full index of all process
// properties found.
func getProcessProperties(r io.Reader) (Process, error) {

	// Setup a new scanner and process each line splitting on ':' character.
	//
	// Then, trim the whitespace from column 1 and column 2.
	//
	// Create a lowercased version of column one and use as a map key. Store
	// column two as value in the map.
	//
	// After we populate the map, we retrieve specific values *from* the map
	// and convert (where needed) to type safe values for later direct use.
	//
	// Of interest are these properties:
	//
	// Name
	// State
	// Pid
	// PPid
	// Threads (could be useful for determining if something is running away?)
	// VMSwap
	//
	// Other values are stored as-is for later use if needed.

	var process Process
	process.AllProperties = make(Properties)

	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {

		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}

		// Blank lines have been skipped, so anything remaining should be in
		// the expected key:value format.
		//
		// NOTE: We have to limit the number of splits as some process names
		// may contain colons (e.g., 'flush-253:0').
		pair := strings.SplitN(line, ":", 2)

		// We should not have *more* than 2 fields for a key:value entry, but
		// we might only have the key (first) field. We can assert that no
		// more than 2 fields are present, but having just one field is
		// acceptable.
		var key string
		var value string
		switch {
		case len(pair) > 2 || len(pair) == 0:
			return Process{}, fmt.Errorf(
				"error parsing property for process: %w",
				ErrInvalidProcStatusLineFormat,
			)

		case len(pair) == 2:
			key = strings.ToLower(strings.TrimSpace(pair[0]))
			value = strings.TrimSpace(pair[1])
		case len(pair) == 1:
			key = strings.ToLower(strings.TrimSpace(pair[0]))
		}

		process.AllProperties[key] = value

	}

	// Handle any errors from scanning the status file.
	if err := scanner.Err(); err != nil {
		return Process{}, err
	}

	// Finally, let's populate the type-specific fields of our Process entry.
	if err := process.setProcessProps(); err != nil {
		return Process{}, err
	}

	return process, nil

}

// FromProcDirs evaluates the status file within a given list of /proc/[pid]
// directories and returns either a collection of Process values or an error
// if one occurs.
func FromProcDirs(procDirs []string) (Processes, error) {
	processes := make(Processes, 0, len(procDirs))
	for _, procDir := range procDirs {

		qualifiedPath := filepath.Join(ProcRootDir, procDir, ProcStatusFilename)
		p, err := ParseProcStatusFile(qualifiedPath)
		if err != nil {
			switch {
			case errors.Is(err, os.ErrNotExist):
				// open /proc/22991/status: no such file or directory
				//
				// This is commonly encountered for short lived processes; we
				// see the process directory when initially listing process
				// directories within the proc filesystem, but then don't find
				// it again when we attempt to parse the status file within
				// each subdirectory. In order to prevent erroring out over
				// short-lived processes we skip to the next item to evaluate.
				continue

			default:
				return nil, fmt.Errorf(
					"fatal error encountered processing proc status file %s: %w",
					qualifiedPath,
					err,
				)
			}
		}

		processes = append(processes, p)

	}

	return processes, nil
}
