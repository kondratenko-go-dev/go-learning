package main

import (
	"fmt"
	"math"
)

func main() {
	a, b, c := 0.1, 0.2, 0.3
	sum := a + b
	fmt.Printf("sum: %v\n", sum)
	fmt.Printf("naive equal: %v\n", sum == c)

	fmt.Printf("precise: %.20f\n", sum)

	epsilon := 1e-9
	epsilonEqual := math.Abs(sum-c) < epsilon
	fmt.Printf("epsilon equal: %v\n", epsilonEqual)

	infValue := math.Inf(1)
	fmt.Printf("inf: %v\n", infValue)

	zero := 0.0
	one := 1.0
	divByZero := one / zero
	fmt.Printf("div by zero: %v\n", divByZero)

	nanEqual := math.NaN() == math.NaN()
	fmt.Printf("nan equals nan: %v\n", nanEqual)
}

/*
Why math.NaN() == math.NaN() returns false:
According to the IEEE 754 standard for floating-point arithmetic, NaN (Not-a-Number)
represents an undefined or unrepresentable numerical value (such as the result of 0.0/0.0).
By definition, any comparison involving NaN returns false because one undefined state
cannot be assumed to be identical to another undefined state. Thus, even NaN == NaN is false.

How to correctly check for NaN:
To check if a floating-point variable `x` is NaN, you should use the built-in function
`math.IsNaN(x)` from the standard library. Alternatively, relying on the IEEE 754 specification,
you can check it using the expression `x != x`. If this expression evaluates to true,
then `x` is guaranteed to be NaN.
*/
