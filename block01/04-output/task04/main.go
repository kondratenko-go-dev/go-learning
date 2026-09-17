package main

import (
	"fmt"
	"os"
)

func main() {
	itemsCount := 3
	totalPrice := 592.93

	report := fmt.Sprintf("report: %d items, total %.2f", itemsCount, totalPrice)
	fmt.Println(report)

	fmt.Fprintln(os.Stderr, "warning: prices are estimates")

	fmt.Fprintf(os.Stderr, "checked %d positions\n", itemsCount)

	/*
		Terminal Commands Breakdown:

		1. `go run ./block01/04-output/task04`
		   - Output shown: All three lines (the report, the warning, and the checked positions).
		   - Reason: By default, both stdout (1) and stderr (2) are printed to the terminal screen.

		2. `go run ./block01/04-output/task04 2>/dev/null`
		   - Output shown: Only the report line ("report: 3 items, total 592.93").
		   - Reason: The '2>' operator redirects the stderr stream (file descriptor 2) to /dev/null,
		     effectively discarding all warnings and error logs.

		3. `go run ./block01/04-output/task04 1>/dev/null`
		   - Output shown: Only the warning and checked lines.
		   - Reason: The '1>' operator redirects the stdout stream (file descriptor 1) to /dev/null,
		     hiding the main program output but preserving error/diagnostic messages.

		Why Stream Separation Matters:
		Separating streams is crucial for automation, logging, and building CLI tools.
		It allows users and scripts to pipe clean data (stdout) into files or other programs
		while simultaneously routing diagnostic logs, warnings, or errors (stderr) to separate
		log files or monitoring systems.
	*/

}
