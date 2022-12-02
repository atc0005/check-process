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
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// GetProcDirs evaluates the given base path (usually "/proc") for known
// process directories within the proc filesystem and returns an unqualified
// list of all valid matches.
func GetProcDirs(path string) ([]string, error) {
	// Will panic if the regex fails to compile. Because the regex value is
	// fixed and unlikely to be modified this should be a safe decision.
	re := regexp.MustCompile(ProcDirRegex)

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to evaluate processes using proc path %s: %w",
			path,
			err,
		)
	}

	procDirsCount := func() int {
		var ctr int
		for _, path := range files {
			if path.IsDir() && re.MatchString(path.Name()) {
				ctr++
			}
		}
		return ctr
	}()

	procDirs := make([]string, 0, procDirsCount)
	for _, path := range files {
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
