package main

import "fmt"

func main() {
	// --- Slices -------------------------------
	fmt.Println("# ==================== Slice ====================== #")
	// Create a slice of length 5 with default zero values:
	numbers := make([]int, 5)
	fmt.Println("Slice with default values:", numbers)
	fmt.Println("len(numbers) =", len(numbers))
	fmt.Println("cap(numbers) =", cap(numbers))

	// Creating a slice with predefined values
	numbers = []int{19, 27, 34, 41}
	fmt.Println("\nSlice from an array:", numbers)

	// Slices can be iterated using a `for` loop:
	for idx, num := range numbers {
		fmt.Printf("item %d ==> %d\n", idx, num)
	}

	// Adding elements to the slice:
	// if the slice exceeds its capacity, it allocates a new underlying array:
	numbers = append(numbers, 48, 26)
	fmt.Println("\nAfter appending elements:", numbers)

	// Slicing the slice to create a sub-slice
	subSlice := numbers[1:4]
	fmt.Println("Sub-slice (index 1 up to 4)", subSlice)

	// Appending another slice:
	numbers2 := []int{-1, -2}
	numbers = append(numbers, numbers2...)
	fmt.Println("Another slice:", numbers2)
	fmt.Println("After appending above slice to **numbers**:", numbers)


	// Go doesn't have a built-in `delete` operation for slices:
	// To remove, say, 2nd item in an slice:
	index := 2
	numbers = append(numbers[:index], numbers[index+1:]...)
	fmt.Println("After removing the 2nd item:", numbers)

	// Clear a slice:
	numbers = numbers[:0]
	fmt.Println("The slice is cleared now:", numbers)

	// --- Maps --------------------------------------
	fmt.Println("\n# ==================== Map ======================== #")
	// Creating a map to store employee names with their respective ID number
	employees := map[string]int{
		"Anna": 19,
		"Alice": 26,
		"Lewis": 44,
	}

	fmt.Println("Create a map with initial values:", employees)
	employees["Toto"] = 1 // Adding a new entry to the map
	employees["Alice"] = 27
	fmt.Println("After adding and updating item :", employees)

	// Accessing a value using a key
	empID := employees["Alice"]
	fmt.Println("ID of Alice:", empID)

	// Checking if a key exists in the map
	empID, exists := employees["Eve"]
	if exists {
		fmt.Println("ID of Eve: ", empID)
	} else {
		fmt.Println("Even not found in the map.")
	}

	// Deleting a key-value pair from the map.
	delete(employees, "Toto")
	fmt.Println("Map after deleting **Toto**", employees)

	// Get map length
	fmt.Println("Number of employees:", len(employees))

	// Iterate over a map:
	println("\nLet's iterate over the map:------------")
	for key, val := range employees {
		fmt.Printf("Key: %s, Value: %d\n", key, val)
	}

	// Clear the map
	employees = make(map[string]int)
	fmt.Println("\nMap is cleared:", employees)
}
