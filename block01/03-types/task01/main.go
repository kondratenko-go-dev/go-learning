package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("int8 range: %d .. %d\n", math.MinInt8, math.MaxInt8)
	fmt.Printf("uint8 max: %d\n", math.MaxUint8)
	fmt.Printf("int64 max: %d\n", math.MaxInt64)

	var x int8 = 127
	fmt.Printf("int8 overflow: %d -> ", x)
	x++
	fmt.Printf("%d\n", x)

	var y uint8 = 0
	fmt.Printf("uint8 underflow: %d -> ", y)
	y--
	fmt.Printf("%d\n", y)
}

/*
127 + 1 = -128 (int8):
In Go, signed integers use Two's Complement representation.
The number 127 in binary is 01111111 (the first bit is the sign bit, 0 means positive).
When we add 1 (00000001) to 01111111, binary arithmetic results in 10000000.
In Two's Complement for an 8-bit signed integer, the most significant bit (MSB) has a negative weight (-2^7 = -128).
Therefore, 10000000 represents exactly -128, which causes an overflow from max to min value.

0 - 1 = 255 (uint8):
The number 0 in binary is 00000000.
When we subtract 1 (00000001) from 00000000, we need to borrow from a higher non-existent bit (essentially treating it as 100000000 - 00000001).
This binary subtraction wraps around within the 8-bit limit, resulting in 11111111.
Since uint8 is unsigned, all bits have positive weights (128 + 64 + 32 + 16 + 8 + 4 + 2 + 1).
Therefore, 11111111 in unsigned binary represents exactly 255.
*/
