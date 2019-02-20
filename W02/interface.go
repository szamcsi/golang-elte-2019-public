package main

import "fmt"

// Define methods on myType that will satisfy HasNameAndSize interface.
type myType int

/*
 interfaces
*/

type HasName interface {
	Name() string
}

type HasSize interface {
	Size() int
}

type HasNameAndSize interface {
	Name() string
	Size() int
}

// This example shows how you can compose interfaces.
type HasNameSizeAndColor interface {
	HasName
	HasSize
	Color() string
}

/*
 functions
*/

func printName(named HasName) {
	fmt.Println("Name", named.Name())
}

func printNameAndSize(h HasNameAndSize) {
	fmt.Println("Name", h.Name(), "size", h.Size())
}

// testing...
func main() {
	var x myType
	printNameAndSize(x)
	printName(x) // note how anything that implements HasNameAndSize interface also implements HasName interface.

	/*
	  var hnsc HasNameSizeAndColor = x  // note how x does not satisfy HasNameSizeAndColor.
	  printName(hnsc)
	*/
}
