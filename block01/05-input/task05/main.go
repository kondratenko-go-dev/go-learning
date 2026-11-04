package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: task05 <count>")
		os.Exit(1)
	}

	arg := os.Args[1]

	count, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid count: \\\"%s\\\"\\n", err)
		os.Exit(1)
	}

	if count < 1 {
		fmt.Fprintln(os.Stderr, "count must be positive")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		fmt.Fprintln(os.Stderr, "no input")
		os.Exit(1)
	}

	text := scanner.Text()
	text = strings.TrimSpace(text)

	if text == "" {
		fmt.Fprintln(os.Stderr, "no input")
		os.Exit(1)
	}

	result := strings.Repeat(text+" ", count)
	result = strings.TrimRight(result, " ")

	fmt.Println(result)
}
