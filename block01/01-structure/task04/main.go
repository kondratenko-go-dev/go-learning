package main

import (
	"fmt"

	"github.com/kondratenko-go-dev/go-learning/block01/01-structure/task04/greet"
)

func main() {
	fmt.Println(greet.Hello("Aleksandr"))
	//fmt.Println(greet.whisper("test"))

	//name whisper not exported by package greet
}
