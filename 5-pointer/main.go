package main

import "fmt"

func main() {
	age := 26 // declare a variable

	// Use `&`` to get the address of the variable:
	pointerToAge := &age

	// Print the value and the address
	fmt.Println("Value of age:", age)              // Outputs the value of the variable
	fmt.Println("Address of age:", pointerToAge)   // Outputs the memory address of the variable
	fmt.Println()

	// Change the value via the pointer:
	age += 1            // increment the value the normaly way
	fmt.Println("Increment value of age by direct assignment:", age)

	*pointerToAge += 1  // Dereference the pointer to update the value at that memory address
	fmt.Println("Increment value of age via pointer:", age)      // Value of age is update
}


// This flexibility is a result of Go's ability to work with pointers for managing variables in memory.
// Both approaches change the same underlying value, but using a pointer allows indirect modification,
// which is useful in scenarios like function arguments where you want to change the original variable
// without returning values explicitly.
