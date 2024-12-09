package main

import "fmt"

func formatValues(itemsCnt int, price float64) string {
	// Calculate total price
	// N.B. Go does NOT automatically convert between types.
	// If you try to multiply `int` & `float64` directly, Go will throw a **type error**.
	totalPrice := float64(itemsCnt) * price
	// Return the formatted string usin `fmt.Sprintf`, instead of printing directly.
	return fmt.Sprintf("You have %d items, for a total of %.2f", itemsCnt, totalPrice)
}

func main() {
	var num int = 5
	const pricePerItem = 27.19

	// short-hand declaration
	result := formatValues(num, pricePerItem)

	// `fmt.Printf` is used to format & print to the standard output (console).
	fmt.Printf("*** Your Invoice ***\n%s\n", result)
}
