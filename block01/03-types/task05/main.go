package main

import (
	"fmt"
	"strconv"
)

func main() {
	inputs := []string{"42", "abc", "3.14", "-7"}

	n0, err0 := strconv.Atoi(inputs[0])
	if err0 != nil {
		fmt.Println("atoi error:", err0)
	} else {
		fmt.Println("atoi ok:", n0)
	}

	n1, err1 := strconv.Atoi(inputs[1])
	if err1 != nil {
		fmt.Println("atoi error:", err1)
	} else {
		fmt.Println("atoi ok:", n1)
	}

	n2, err2 := strconv.Atoi(inputs[2])
	if err2 != nil {
		fmt.Println("atoi error:", err2)
	} else {
		fmt.Println("atoi ok:", n2)
	}

	f2, errFloat2 := strconv.ParseFloat(inputs[2], 64)
	if errFloat2 != nil {
		fmt.Println("parsefloat error:", errFloat2)
	} else {
		fmt.Println("parsefloat ok:", f2)
	}

	n3, err3 := strconv.Atoi(inputs[3])
	if err3 != nil {
		fmt.Println("atoi error:", err3)
	} else {
		fmt.Println("atoi ok:", n3)
	}
}

/*
Why Atoi("3.14") returns an error instead of truncating or rounding to 3:
The name `Atoi` stands for "ASCII to Integer". By design, its sole responsibility
is to parse strings that exclusively represent whole numbers (integers), including
an optional leading sign ('+' or '-').

When `strconv.Atoi` encounters the dot character ('.') in "3.14", it treats it as an
invalid character for an integer format, rather than a decimal separator. Go values
predictability and explicit behavior over implicit assumptions. Implicitly rounding or
truncating floating-point strings during integer parsing could hide bugs in data validation
or communication protocols. If a string contains a fractional part, you must explicitly
parse it using `strconv.ParseFloat` first, and then perform any truncation or rounding
manually if needed.
*/
