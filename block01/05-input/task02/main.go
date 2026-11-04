package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	var age int

	n, err := fmt.Scan(&name, &age)

	fmt.Printf("read %d values\n", n)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("name: %s\n", name)
	fmt.Printf("age: %d\n", age)
	fmt.Printf("next year: %d\n", age+1)
}

/*
SCENARIO RESULTS & ANALYSIS:

1. printf 'Aleksandr 39\n' | go run ...
Output:
read 2 values
name: Aleksandr
age: 39
next year: 40
exit: 0
Explanation: Both values match the expected types (string and int). Scanning finishes successfully.

2. printf 'Aleksandr abc\n' | go run ...
Output:
read 1 values
error reading input: expected integer
exit: 1
Explanation: 'Aleksandr' is scanned successfully into the string variable (n=1). Then, 'abc'
fails to parse as an integer for the 'age' variable. fmt.Scan stops immediately and returns an error.

3. printf 'Aleksandr\n' | go run ...
Output:
read 1 values
error reading input: EOF
exit: 1
Explanation:
- Why it exits immediately with printf: The pipe closes the input stream after sending 'Aleksandr\n'.
  fmt.Scan reads 'Aleksandr', hits the End-Of-File (EOF), realizes there is no more data for 'age',
  and returns n=1 with an EOF error.
- What happens with interactive keyboard input: If you run this manually and type 'Aleksandr' and press Enter,
  the program WILL NOT exit. It will hang and wait. This happens because fmt.Scan treats newlines (\n)
  as spaces. It keeps waiting indefinitely for the second value ('age') until it either receives a number
  or the standard input stream is explicitly closed (e.g., via Ctrl+D).
*/
