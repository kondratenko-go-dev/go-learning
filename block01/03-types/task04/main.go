package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	count := 7
	price := 19.99
	total := float64(count) * price
	fmt.Printf("total: %.2f\n", total)

	f1, f2 := 3.9, -3.9
	fmt.Printf("truncate: %v -> %d, %v -> %d\n", f1, int(f1), f2, int(f2))

	fmt.Printf("round: %v -> %.0f, %v -> %.0f\n", f1, math.Round(f1), f2, math.Round(f2))

	var big int = 300
	fmt.Printf("overflow: 300 as int8 -> %d\n", int8(big))

	fmt.Printf("string(rune(65))=%s strconv.Itoa(65)=%s\n", string(rune(65)), strconv.Itoa(65))

	str := "Go"
	fmt.Printf("bytes: %v runes: %v\n", []byte(str), []rune(str))
}

/*
When replacing `string(rune(65))` with `string(65)`, the `go vet` command issues the following warning:

./main.go:27:54: conversion from int to string yields a string of one rune, not a string of digits (did you mean strconv.Itoa?)

1. The `string(65)` construct interprets the number 65 not as the string "65", but as a character code (Unicode code point).
   Since 65 is the code for the uppercase Latin letter 'A', the result is the string "A".
   Go considers a direct cast of an integer to `string` to be a potential error—as developers
   often confuse it with converting the number to its text representation—and requires
   either an explicit cast to `rune` or the use of the `strconv.Itoa` function.

2. Overflowing an `int8` with 300 results in 44 because 300 in binary is 100101100.
   When casting to an 8-bit `int8`, the high-order bits are discarded, leaving only the low-order byte: 00101100,
   which equals 44 in decimal.
*/
