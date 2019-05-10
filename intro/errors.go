package main

import (
	"fmt"
)

type user struct {
	username string
}

/* validateUsers returns nil error if all users are valid (have a non-empty username),
   or an error message with list of indices in case some aren't.
   Example error message:

   "invalid list of users: indices [0 3 5] are invalid"

   Use fmt.Errorf("error blah blah %v: %v, %v", arg1, arg2, arg3) (similar to fmt.Sprintf) to generate a new "error" value.
*/
func validateUsers(users []user) error {
	// replace function body
	return nil
}

func main() {
	// expect error: element 2 is invalid
	users := []user{{"zero"}, {"one"}, {}, {"three"}, {"four"}}

	if err := validateUsers(users); err == nil {
		fmt.Println("All ok!")
	} else {
		fmt.Println("Failed:", err)
	}
}
