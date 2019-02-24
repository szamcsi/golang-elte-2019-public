package main

import "fmt"

type user struct {
	username string
	id       int
}

func swap(one *user, two *user) {
	// replace this function body. Swap should swap user values of one and two.
}

func main() {
	u1 := user{"foo", 123}
	u2 := user{"bar", 234}

	fmt.Println("Before swap: u1:", u1, "u2:", u2)
	// invoke swap() here with correct arguments
	fmt.Println("After swap : u1:", u1, "u2:", u2)
}
