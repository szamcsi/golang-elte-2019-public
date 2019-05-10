package main

import "fmt"

type user struct {
	username string
	id       int
}

func resetID(u user) user {
	// replace function body
	return u
}

func resetIDInPlace(u user) {
	// replace function body
}

func main() {
	u1 := user{"alice", 11}
	fmt.Println("u1: id should be 0, got", resetID(u1).id)

	u2 := user{"bob", 12}
	resetIDInPlace(u2)
	fmt.Println("u2: id should be 0, got", u2.id)
}
