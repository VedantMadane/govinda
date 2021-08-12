
// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

// Add imports.
import "fmt"
func main() {

	// Declare a nil slice of integers.
	slice := []int{}
	fmt.Print(slice)

	// Append numbers to the slice.
	slice = append(slice, 0, 1, 2,3 ,4 , 5)
	fmt.Print(slice)

	// Display each value in the slice.
	for _, value := range slice {  
  	fmt.Printf("%d\n", value)  
 	}  

	// Declare a slice of strings and populate the slice with names.
	//var names []string
	names := []string{"Kaustubh", "Vedant", "Omkar"}

	// Display each index position and slice value.
	for index, _ := range names{
	fmt.Printf("%d\n", index)
	}

	// Take a slice of index 1 and 2 of the slice of strings.
	var my_slice_1 = slice[1:3]
	var my_slice_2 = names[1:3]
	fmt.Print(my_slice_1, my_slice_2)
	// Display each index position and slice values for the new slice.
	for index, ele := range slice {
        fmt.Printf("Index = %d and element = %d\n", index, ele)

	for index, ele := range names{
        fmt.Printf("Index = %d and element = %s\n", index, ele)
    }
    }
}
