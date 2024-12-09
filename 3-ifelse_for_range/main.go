package main

import "fmt"

func main() {
	// Define an array of integers
	numbers := [5]int{19, 27, 34, 41, 48}

	// Loop through each number in the array
	for idx, num := range numbers {
		// Check if the number is even or odd
		if num%2 == 0 {
			fmt.Printf("%d: %d is even.\n", idx+1, num)
		} else {
			fmt.Printf("%d: %d is odd.\n", idx+1, num)
		}
	}
}
