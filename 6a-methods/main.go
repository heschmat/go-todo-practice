package main

import "fmt"

// Define a **struct** called "Car"
type Car struct {
	Make  string
	Model string
	Year  int
}

// Method to display the car details.
// This method has a `*Car` receiver, which means it operates on the Car object:
func (c *Car) DisplayDetails() {
	fmt.Printf("Car Details: \nMake: %s \nModel: %s \nYear: %d\n", c.Make, c.Model, c.Year)
}

// Method to update the year of the car.
// This method also has a `*Car` receiver, so it modifies the car object directly:
func (c *Car) UpdateYear(newYear int) {
	c.Year = newYear
	fmt.Printf("\nThe year of the car has been updated to `%d`\n", c.Year)
}

// MEthod to create a new car & return it as a value (not modifying the original).
// This method has a value receiver (no * before Car)
// So it does NOT modify the original car object.
func (c Car) GetCarDescription() string {
	return fmt.Sprintf("%d %s %s", c.Year, c.Make, c.Model)
}


func main() {
	// Create a new instance of Car
	car1 := &Car{
		Make:  "Tpyota",
		Model: "Corolla",
		Year: 2020,
	}

	car1.DisplayDetails()
	car1.UpdateYear(2024)

	// Call `GetCarDescription` to get a textual description of the car
	description := car1.GetCarDescription()
	fmt.Println("Car1 description:", description)

	// Creating another car using value receiver method:
	car2 := Car{
		Make: "Honda",
		Model: "Civic",
		Year: 2018,
	}
	// car2 remains unchanged after the method call
	// because `GEtCarDescription()` uses a value receiver.
	car2Details := car2.GetCarDescription()
	fmt.Println("Car2 description:", car2Details)

	fmt.Println("--------------")
	car2.DisplayDetails()
}
