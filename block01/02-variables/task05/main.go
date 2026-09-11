package main

import (
	"fmt"
)

func main() {
	// Experiment 1: Try to change a constant
	// const limit = 10
	// limit = 20
	// Error: cannot assign to limit (untyped int constant 10)

	// Experiment 2: Constant from a function call
	// const start = time.Now()
	// Error: time.Now() (value of type time.Time) is not constant

	// Experiment 3: Typed constant overflow
	// const big int = 1 << 100
	// Error: int overflows 1 << 100

	// Experiment 4: Untyped huge constant
	const huge = 1 << 100
	fmt.Println("shifted=", huge>>98)
	// fmt.Println(huge)
	// Error: cannot use huge (untyped int constant 1267650600228229401496703205376) as int value in argument to fmt.Println (overflows)

	// Experiment 5: Take the address of a constant
	// const Pi = 3.14
	// p := &Pi
	// Error: cannot take the address of Pi (untyped float constant 3.14)

	// Working part
	const limit = 10
	fmt.Println("limit=", limit)
}
