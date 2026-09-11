package main

import "fmt"

func main() {
	a := 5
	b := 12

	fmt.Printf("before a=%d b=%d\n", a, b)

	a, b = b, a

	fmt.Printf("after a=%d b=%d\n", a, b)

	x, y, z := 1, 2, 3

	fmt.Printf("before x=%d y=%d z=%d\n", x, y, z)

	x, y, z = z, x, y

	fmt.Printf("after x=%d y=%d z=%d\n", x, y, z)

	_, keep := 100, 200

	fmt.Printf("keep=%d\n", keep)
}
