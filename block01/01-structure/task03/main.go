package main

import "fmt"

func main() {
	fmt.Println("Hello, Go!")

	// Broken experiments:
	// 1. unused import "os" -> "os" imported and not used
	// 2. brace on new line   -> syntax error: unexpected semicolon or newline before {
	// 3. package hello       -> s not a main package
	// 4. fmt.println         -> undefined: fmt.println (but have Println)
}
