package main

import "fmt"

func main() {
	value := 1

	if true {
		value := 2
		fmt.Println("inside if:", value)
	}
	fmt.Println("after if:", value)

	for i := 0; i < 1; i++ {
		value = 3
		fmt.Println("inside for:", value)
	}
	fmt.Println("after for:", value)

	// Prediction:
	// inside if:  2
	// after if:   1
	// inside for: 3
	// after for:  3
	//
	// Actual:
	// inside if:  2
	// after if:   1
	// inside for: 3
	// after for:  3
	//
	// Why the difference:
	// The difference between the 'if' and 'for' blocks comes down to a single character: the colon ':' in the assignment operator.
	//
	// In the 'if' block, 'value := 2' uses the short variable declaration operator (:=).
	// This creates a completely new local variable named 'value' that shadows the outer 'value' variable within the 'if' block scope.
	//
	// In the 'for' block, 'value = 3' uses the standard assignment operator (=).
	// Since there is no colon, it does not declare a new variable. Instead, it modifies the existing outer 'value' variable.

}
