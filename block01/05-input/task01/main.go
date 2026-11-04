package main

import (
	"fmt"
	"os"
)

func main() {
	argsCount := len(os.Args)
	fmt.Printf("args count: %d\n", argsCount)

	if argsCount > 0 {
		fmt.Printf("args[0]: %s\n", os.Args[0])
	}
	if argsCount > 1 {
		fmt.Printf("args[1]: %s\n", os.Args[1])
	}
	if argsCount > 2 {
		fmt.Printf("args[2]: %s\n", os.Args[2])
	}
	if argsCount > 3 {
		fmt.Printf("args[3]: %s\n", os.Args[3])
	}

	userArgsCount := argsCount - 1

	if userArgsCount < 0 {
		userArgsCount = 0
	}

	fmt.Printf("user args count: %d\n", userArgsCount)
}

/*
What changed when running without arguments (go run ./block01/05-input/task01):
The "args count" became 1 because os.Args still contains the path to the executed binary as the 0-th element.
Only "arg[0]" is printed, while arg[1], arg[2], and arg[3] are omitted because they do not exist.
The "user args" count became 0.
*/

/*
Difference between 'go run' and compiled binary execution:
When using 'go run', arg[0] points to a temporary executable file created by Go in a system temporary directory (e.g., /tmp/go-build...).
When running the compiled binary directly (./bin/app), arg[0] contains the actual path used to invoke the program, which is './bin/app'.
*/
