// Copyright 2022 Adam Chalkley
//
// https://github.com/atc0005/check-process
//
// Licensed under the MIT License. See LICENSE file in the project root for
// full license information.

package reports

import (
	"fmt"
	"io"
	"strings"

	"github.com/atc0005/check-process/internal/process"
	"github.com/atc0005/go-nagios"
)

// CheckProcessOneLineSummary returns a one-line summary of the evaluation
// results suitable for display and notification purposes.
func CheckProcessOneLineSummary(processes process.Processes) string {
	var summary string

	switch {

	case !processes.IsOKState():

		probProcs := processes.States(process.KnownProblemProcessStates())

		summaryList := processes.SummaryList()
		summaryList = append(summaryList, fmt.Sprintf(
			"evaluated [%d]",
			processes.Count(),
		))
		procsSummary := strings.Join(summaryList, ", ")

		summary = fmt.Sprintf(
			"%s: %d problematic processes found (%s)",
			processes.ServiceState().Label,
			probProcs.Count(),
			procsSummary,
		)

	case processes.IsOKState():
		summary = fmt.Sprintf(
			"%s: No problematic processes found (%d evaluated)",
			processes.ServiceState().Label,
			processes.Count(),
		)

	default:
		summary = "BUG: Expected processes collection state unexpected"

	}

	return summary

}

// writeReportHeader generates a "header" or lead-in summary for the final
// plugin report.
func writeReportHeader(w io.Writer, processes process.Processes) {

	_, _ = fmt.Fprintf(w, "Process Summary:%[1]s%[1]s", nagios.CheckOutputEOL)
	for _, item := range processes.SummaryList() {
		_, _ = fmt.Fprintf(w, "  - %s\n", item)
	}
	_, _ = fmt.Fprintf(w, "%[1]s%[1]s", nagios.CheckOutputEOL)

}

// writeReportProblemEntries generates a listing of problem process entries
// for the final plugin report.
func writeReportProblemEntries(w io.Writer, processes process.Processes) {

	probProcs := processes.States(process.KnownProblemProcessStates())

	_, _ = fmt.Fprintf(w, "%[1]sProblems:%[1]s", nagios.CheckOutputEOL)

	switch {
	case len(probProcs) > 0:

		for _, p := range probProcs {
			parentProcess, err := p.ParentProcess(processes)
			ppName, ppID := parentProcess.Name, parentProcess.Pid
			if err != nil {
				ppName = "missing"
				ppID = -1
			}

			_, _ = fmt.Fprintf(
				w,
				// "Name: %s\n\tParent: %v (%v)\n\tState: %v\n\tPid: %v\n\tPPid: %v\n\tThreads: %v\n\n",
				"  - Name: %10s [Parent: %v (%v), State: %v, Pid: %v, PPid: %v, Threads: %v]%s",
				p.Name,
				ppName,
				ppID,
				p.State,
				p.Pid,
				p.PPid,
				p.Threads,
				nagios.CheckOutputEOL,
			)
		}
	default:
		_, _ = fmt.Fprintf(w, "%[1]s  - None%[1]s", nagios.CheckOutputEOL)
	}

}

// CheckProcessReport returns a formatted report of the evaluation results
// suitable for display and notification purposes.
func CheckProcessReport(processes process.Processes) string {
	var report strings.Builder

	writeReportHeader(&report, processes)

	_, _ = fmt.Fprintf(
		&report,
		"%[2]s%[1]s%[1]s",
		nagios.CheckOutputEOL,
		strings.Repeat("-", 50),
	)

	writeReportProblemEntries(&report, processes)

	return report.String()
}
