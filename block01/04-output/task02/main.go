package main

import "fmt"

func main() {
	headerFmt := "%-12s%8s%5s\n"
	dataFmt := "%-12s%8.2f%5d\n"

	fmt.Printf(headerFmt, "Name", "Price", "Qty")

	p1, q1 := 19.99, 7
	p2, q2 := 5.50, 12
	p3, q3 := 129.00, 3

	fmt.Printf(dataFmt, "Coffee", p1, q1)
	fmt.Printf(dataFmt, "Tea", p2, q2)
	fmt.Printf(dataFmt, "Chocolate", p3, q3)

	totalPrice := (p1 * float64(q1)) + (p2 * float64(q2)) + (p3 * float64(q3))
	totalQty := q1 + q2 + q3

	fmt.Printf(dataFmt, "TOTAL", totalPrice, totalQty)
}
