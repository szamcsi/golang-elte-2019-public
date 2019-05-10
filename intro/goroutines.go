package main

import (
	"fmt"
)

// write plusOne and multByTwo functions to make this code print 22.

func main() {
	input := make(chan int)
	pass := make(chan int)
	output := make(chan int)

	go plusOne(input, pass)
	go multByTwo(pass, output)

	input <- 10
	fmt.Println("Want to get 22, got", <-output)
}
