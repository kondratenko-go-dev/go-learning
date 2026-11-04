package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var totalLines int
	var nonEmptyLines int
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()

		fmt.Printf("lineNum: %d: %s\n", lineNum, line)

		if strings.TrimSpace(line) != "" {
			nonEmptyLines++
		}
		totalLines++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
		os.Exit(1)
	}

	fmt.Printf("total lines: %d\n", totalLines)
	fmt.Printf("non empty lines: %d\n", nonEmptyLines)
}

/*
Explanation of Ctrl+D vs Ctrl+C (as requested):

In Unix-like systems and terminals:
- Ctrl+D sends an EOF (End Of File) signal to the current process. It tells the program
  (like our bufio.Scanner) that the standard input stream has ended successfully.
  The program finishes reading normally and executes the remaining code (printing totals).

- Ctrl+C sends a SIGINT (Signal Interrupt) to the running process. It immediately terminates
  the program from the outside. If you press Ctrl+C, the program stops instantly,
  and the final summary (total lines / non-empty lines) will NOT be printed.
*/
