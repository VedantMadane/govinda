
// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main


// Add imports
import(
	"fmt"
	
	)

// main is the entry point for the application.
func main() {


	// Declare variables that are set to their zero value.
		var a int
		var b bool
		var c int16 
	// Display the value of those variables.
	fmt.Printf(" %T [%v]\n %T [%v]\n %T [%v]", a, a, b, b, c, c)
	// Declare variables and initialize.

	// Using the short variable declaration operator.
	aa:= 10
	bb:= "hare kRSNa"
	cc:= false

	// Display the value of those variables.
	fmt.Printf("%T [%v]\n%T [%v]\n%T [%v]\n", aa,aa,bb,bb,cc,cc)

	// Perform a type conversion.
	ee := int(aa)
	ff := int64(aa)

	// Display the value of that variable.
	fmt.Printf("%T [%v]\n%T [%v]\n%T [%v]\n", aa,aa,ee,ee,ff,ff)

}
