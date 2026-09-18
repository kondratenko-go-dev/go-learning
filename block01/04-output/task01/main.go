package main

import "fmt"

func main() {
	value := 65
	fmt.Printf("v=%v T=%T d=%d b=%b o=%o x=%x X=%X c=%c q=%q U=%U\n", value, value, value, value,
		value, value, value, value, value, value) //fmt.Printf format %q has arg value of wrong type int

	pi := 1234.5678
	fmt.Printf("v=%v f=%f .2f=%.2f e=%e g=%g T=%T\n", pi, pi, pi, pi, pi, pi)

	text := "Go\tlang"
	fmt.Printf("s=%s q=%q T=%T len=%d\n", text, text, text, len(text))
}
