package main

import "fmt"

/*

Define a recursive function "fibRec(n int)" that will return Nth Fibonacci number,
and an iterative function "fibIter(n int)" that will do the same.


Fibonacci numbers are defined as:
F(0): 0
F(1): 1
F(n): F(n-1) + F(n-2)

i.e. 0, 1, 1, 2, 3, 5, 8, ...

*/

func main() {
	fmt.Println("2nd Fibonacci number is 1, got", fibRec(2), "and", fibIter(2))
	fmt.Println("8th Fibonacci number is 21, got", fibRec(8), "and", fibIter(8))
	fmt.Println("9th Fibonacci number is 34, got", fibRec(9), "and", fibIter(9))
	fmt.Println("10th Fibonacci number is 21+34 = 55, got", fibRec(10), "and", fibIter(10))
}
