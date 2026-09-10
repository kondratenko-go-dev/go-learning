package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "go is simple and fast"

	fmt.Println(str)
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ReplaceAll(str, "simple", "powerful"))
}
