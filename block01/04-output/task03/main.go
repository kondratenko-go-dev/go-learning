package main

import "fmt"

func main() {
	fmt.Print("Print strings: ")
	fmt.Print("a", "b")
	fmt.Println()

	fmt.Print("Print numbers: ")
	fmt.Print(1, 2)
	fmt.Println()

	fmt.Print("Print mixed: ")
	fmt.Print("a", 1)
	fmt.Println()

	fmt.Print("Println strings: ")
	fmt.Println("a", "b")

	fmt.Print("Println numbers: ")
	fmt.Println(1, 2)

	fmt.Printf("Printf strings: %s%s\n", "a", "b")
	fmt.Printf("Printf numbers: %d%d\n", 1, 2)

	fmt.Print("Experiment output: ")
	fmt.Print(1, "b", 2)
	fmt.Println()
}

//Rule: fmt.Print automatically inserts a space between arguments if and only if neither of the adjacent arguments is a string. If at least one of the two neighboring arguments is a string, no space is added between them.
