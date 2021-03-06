
// Declare and initialize a variable of type int with the value of 20. Display
// the _address of_ and _value of_ the variable.
//
// Declare and initialize a pointer variable of type int that points to the last
// variable you just created. Display the _address of_ , _value of_ and the
// _value that the pointer points to_.
package main

// Add imports.
import "fmt"
func main() {

	// Declare an integer variable with the value of 20.
	var num int = 20

	// Display the address of and value of the variable.
	fmt.Printf("Address: [%p] \n Value: [%v]\n", &num, num)

	// Declare a pointer variable of type int. Assign the
	// address of the integer variable above.
	var numptr *int = &num

	// Display the address of, value of and the value the pointer
	// points to.
	fmt.Printf("address: [%p] \n Val [%v] \n points to [%v]", &numptr, numptr, *numptr)
}
