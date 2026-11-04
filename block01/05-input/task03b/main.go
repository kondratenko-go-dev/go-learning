package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// bufio.Scanner reads the input line by line (until '\n') by default.

	if scanner.Scan() {
		word := scanner.Text()
		fmt.Printf("scan result: %q\n", word)
	}
}

/*
What happened to the rest of the string:
bufio.NewScanner reads the entire line up to the newline character ('\n') by default.
Therefore, it successfully captures the whole string "Ivan Petrov", leaving nothing extra in the buffer.

Why %q is more convenient than %s here:
Using %q ensures we can explicitly see that the entire string "Ivan Petrov" was captured inside
the quotes, confirming that no trailing spaces or newline characters were accidentally included
or left out.
*/
