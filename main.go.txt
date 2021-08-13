// Declare a struct type and create a value of this type. Declare a function
// that can change the value of some field in this struct type. Display the
// value before and after the call to your function.
package main

// Add imports.
import "fmt"
// Declare a type named user.
type User struct{
	name string
	age int
}
// Create a function that changes the value of one of the user fields.
func funcName(nptr *User, newAge int) {

	// Use the pointer to change the value that the
	// pointer points to.
	nptr.age = newAge
}

func main() {

	// Create a variable of type user and initialize each field.
	var firstUser User = User{"Omkar", 21}
	
	// Display the value of the variable.
	fmt.Print(firstUser, "\n")

	// Share the variable with the function you declared above.
	funcName(&firstUser, 22)

	// Display the value of the variable.
	fmt.Print(firstUser)
}
