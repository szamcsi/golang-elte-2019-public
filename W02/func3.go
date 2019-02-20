package main

import "fmt"

// Define a divWithRem function that takes two ints and returns two ints (result and remainder).
// In Go, "/" on int is integer division, "%" is remainder from division.

func main() {
	var res, rem int
	res, rem = divWithRem(44, 3)
	fmt.Println("44 / 3 = 14 rem 2, got", res, "rem", rem)
	res, rem = divWithRem(99, 3)
	fmt.Println("99 / 3 = 33 rem 0, got", res, "rem", rem)
	res, rem = divWithRem(5, 2)
	fmt.Println("5 / 2 = 2 rem 1, got", res, "rem", rem)
}
