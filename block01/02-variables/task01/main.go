package main

import "fmt"

const myName = "Aleksandr"

func main() {
	var name string
	var age int = 39
	var height = 1.81
	isLearning := true
	var (
		country = "Ukraine"
		city    = "Odessa"
	)
	name = myName

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("IsLearning:", isLearning)
	fmt.Println("City:", city)
	fmt.Println("Country:", country)
}
