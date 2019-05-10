package main

import "fmt"

type User struct{ firstName, lastName string }

// Write a SwapByValue method that swaps first and last name, struct value is the receiver.
func (u User) SwapByValue() {
	// replace method body
}

// Write a SwapByPtr method that swaps first and last name, pointer to struct is the receiver.
func (u *User) SwapByPtr() {
	// replace method body. It will be identical to SwapByValue body
}

func main() {
	u := User{"John", "Doe"}
	fmt.Println("User initially:", u)
	u.SwapByValue()
	fmt.Println("User after swap on value:", u)
	u.SwapByPtr()
	fmt.Println("User after swap on ptr:", u)

	uPtr := &User{"Sherlock", "Holmes"}
	fmt.Println("User ptr initially:", uPtr)
	uPtr.SwapByValue()
	fmt.Println("User ptr after swap on value:", uPtr)
	uPtr.SwapByPtr()
	fmt.Println("User ptr after swap on ptr:", uPtr)
}
