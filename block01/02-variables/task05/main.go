package main

import "fmt"

func main() {
	// Experiment 1: Try to change a constant
	// const limit = 10
	// limit = 20
	// Error: cannot assign to limit (untyped int constant 10)

	// Experiment 2: Constant from a function call
	// const start = time.Now()
	// Error: time.Now() (value of type time.Time) is not constant

	// Experiment 3: Typed constant overflow
	//block01/02-variables/task05/main.go:18:18: cannot use 1 << 100 (
	//untyped int constant 1267650600228229401496703205376) as int value in constant declaration (overflows)
	//const big int = 1 << 100

	// Working part
	const limit = 10
	fmt.Printf("limit=%d\n", limit)

	// Experiment 4: Untyped huge constant
	const huge = 1 << 100
	fmt.Printf("shifted=%d\n", huge>>98)
	//In a nutshell: the boundary between "what the compiler thinks" and "what the program
	//executes" is precisely the point where a constant must acquire a concrete type.
	// fmt.Println(huge)
	// Error: cannot use huge (untyped int constant 1267650600228229401496703205376) as int value in argument to fmt.Println (overflows)

	// Experiment 5: Take the address of a constant
	// const Pi = 3.14
	// p := &Pi
	// Error: cannot take the address of Pi (untyped float constant 3.14)
}
