
// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
// Declare and initialize an anonymous struct type with the same three fields. Display the value.
package main

// Add imports.
import (
	"fmt"
	)
// Add user type and provide comment.
type User struct{
	name string
	email string
	age int
}
func main() {

	// Declare variable of type user and init using a struct literal.
	var firstUser User = User{"Vedant Madane", "vnm@xyz.com", 22}

	// Display the field values.
	fmt.Print(firstUser)

	// Declare a variable using an anonymous struct.
	anon := struct{
	name string
	email string
	age int
	}{"Siddesh Chaugule", "sid@chau.org", 23}

	// Display the field values.
	fmt.Println(anon)
}
