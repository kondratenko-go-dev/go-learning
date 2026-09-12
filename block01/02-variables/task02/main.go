package main

import "fmt"

func main() {
	var a int
	var b float64
	var c string
	var d bool
	var e rune
	var f byte

	fmt.Printf("int: %v\n", a)
	fmt.Printf("float64: %v\n", b)
	fmt.Printf("string: %q\n", c)
	fmt.Printf("bool: %v\n", d)
	fmt.Printf("rune: %c\n", e)
	fmt.Printf("byte: %c\n", f)

	// - 'byte' is an alias for 'uint8' (8-bit unsigned integer).
	// - 'rune' is an alias for 'int32' (32-bit signed integer) representing a Unicode code point.
}
