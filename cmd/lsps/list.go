// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/check-process
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/atc0005/check-process/internal/process"
)

// listProcesses generates a summary of the given Process values and writes it
// to the specified io.Writer.
func listProcesses(w io.Writer, processes process.Processes) {

	switch {
	case len(processes) > 0:
		for _, p := range processes {
			writeProcessInfoLine(w, p, processes, true)
		}

		_, _ = fmt.Fprintf(w, "\nSummary:\n\n")
		for _, item := range processes.SummaryList() {
			_, _ = fmt.Fprintf(w, "  - %s\n", item)
		}
		_, _ = fmt.Fprintln(w)

	default:
		_, _ = fmt.Fprintf(w, "\n  - None\n")
	}

}

// listOtherProcesses generates a summary of the given Process values from the
// complete collection that are not present in the specified evaluated set.
// The summary is written to the specified io.Writer.
func listOtherProcesses(w io.Writer, evaluated process.Processes, all process.Processes, includeDetails bool) {

	remaining := all.Exclude(evaluated...)

	if len(remaining) == 0 {
		return
	}

	_, _ = fmt.Fprintf(w, "\nSUMMARY:\n\n")
	for _, item := range remaining.SummaryList() {
		_, _ = fmt.Fprintf(w, "  - %s\n", item)
	}
	_, _ = fmt.Fprintln(w)

	if includeDetails {
		fmt.Printf("\n%s\n\n", strings.Repeat("-", 50))

		_, _ = fmt.Fprintf(w, "\nDETAILS:\n")

		// for _, p := range remaining {
		// 	writeProcessInfoLine(w, p, remaining)
		// }

		// Create index of process state type to process collection of that state
		// type.
		stateIndex := make(map[string]process.Processes)
		for _, p := range remaining {
			stateIndex[p.State] = append(stateIndex[p.State], p)
		}

		// Write out the state as a header, emit the processes for that state.
		for state, ps := range stateIndex {
			_, _ = fmt.Fprintf(w, "\n%s\n", state)
			for _, p := range ps {
				writeProcessInfoLine(w, p, remaining, false)
			}
		}
	}

}

// writeProcessInfoLine generates a summary of the given Process and writes it
// to the specified io.Writer in a one-line format. The collection of gathered
// Process values is provided in order to resolve dependencies between
// processes (e.g., ancestry).
func writeProcessInfoLine(w io.Writer, p process.Process, ps process.Processes, includeStateField bool) {
	parentProcess, err := p.ParentProcess(ps)
	ppName, ppID := parentProcess.Name, parentProcess.Pid
	if err != nil {
		ppName = "missing"
		ppID = -1
	}

	switch includeStateField {
	case true:
		lineTmpl := "  - Name: %10s [Parent: %v (%v), State: %v, Pid: %v, PPid: %v, VMSwap: %v, Threads: %v]\n"
		_, _ = fmt.Fprintf(
			w,
			lineTmpl,
			p.Name,
			ppName,
			ppID,
			p.State,
			p.Pid,
			p.PPid,
			p.VMSwap,
			p.Threads,
		)
	default:
		lineTmpl := "  - Name: %10s [Parent: %v (%v), Pid: %v, PPid: %v, VMSwap: %v, Threads: %v]\n"
		_, _ = fmt.Fprintf(
			w,
			lineTmpl,
			p.Name,
			ppName,
			ppID,
			p.Pid,
			p.PPid,
			p.VMSwap,
			p.Threads,
		)
	}
}
