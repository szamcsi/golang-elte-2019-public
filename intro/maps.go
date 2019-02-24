package main

import (
	"fmt"
)

/* Write a function that will build a maps-based multiplication table for ints 0..n.
   Writing an element to a nil map results in a panic:
     var m map[int]int
     m[2] = 3  <--- this panics

     var m map[int]int
     m = make(map[int]int)  <--- m is no longer nil
     m[2] = 3               <---- this works now

   Remember that make(map[int]map[int]int) will create a map structure where every element has a type of map[int]int and a default (zero) value of nil.

   However, reading an element from a nil map is perfectly fine, it behaves the same way as reading a non-existent key: returns the zero value.
*/
func multiplicationTable(n int) map[int]map[int]int {
	// replace function body
	return nil
}

func main() {
	m := multiplicationTable(10)
	fmt.Println("3 * 5:", m[3][5])
	fmt.Println("7 * 8:", m[7][8])
	fmt.Println("9 * 9:", m[9][9])
}
