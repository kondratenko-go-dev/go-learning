package main

import "fmt"

func main() {
	fmt.Printf("%d\n", "hello")   // 1. fmt.Printf format %d has arg "hello" of wrong type string
	fmt.Printf("%d %d\n", 1)      // 2. fmt.Printf format %d reads arg #2, but call has 1 arg
	fmt.Printf("%d\n", 1, 2)      // 3. fmt.Printf call needs 1 arg but has 2 args
	fmt.Printf("discount: 50%\n") // 4. fmt.Printf format % has unknown verb
	fmt.Printf("%s\n")            // 5. fmt.Printf format %s reads arg #1, but call has 0 args

	fmt.Printf("discount: 50%%\n")

	/*
		Why the compiler passes formatting errors, but `go vet` does not:

		1. Compiler's Scope:
		   The Go compiler only checks syntax, type safety, and whether functions
		   receive the correct type and number of parameters based on their signature.
		   Since `fmt.Printf` is a variadic function (signature: `Printf(format string, a ...any)`),
		   the compiler only ensures that the first argument is a string and any following
		   arguments are passed. It does not look inside the string literal to parse percentage
		   signs or map them to the types of subsequent arguments.

		2. Go Vet's Scope:
		   `go vet` is a static analysis tool specifically designed to inspect source code
		   for suspicious constructs that are syntactically valid but likely bugs. It knows
		   the internal logic of the standard `fmt` library packages. Therefore, `go vet` parses
		   the format string at build time, manually counts the verbs, checks their compatibility
		   with the provided data types, and flags logic mistakes before the code is even run.
	*/

}
