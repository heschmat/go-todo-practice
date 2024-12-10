# Methods in Go

Methods in Go can have **value receivers** or **pointre receivers**.

- **Value receivers**: The method operates on a copy of the struct.
- **Pointer receivers**: The method operates on the original struct, and can modify it.

Methods can be used to define behaviors associated with structs, which is a fundamental part of Go's **oject-oriented**-like capabilities.

## Methods with Pointer Receivers (`*Car`)

- The method `DisplayDetails()` has a **pointer receiver** `(c *Car)`. This means the method operates on a pointer to a `Car` object, and it can modify the original car object directly.

- Similarly, `UpdateYear(newYear int)` also has a pointer receiver and modifies the year of the car directly.

## Methods with Value Receivers (`Car`)

- The method `GetCarDescription()` has a value receiver `(c Car)`. This means it receives a copy of the `Car` object and does NOT modify the original object. If the method modifies the struct, the changes will NOT reflect on the original instance. It returns a sting describing the car.

## Method value promotion and receiver compatibility

This behavior in Go simplifies method calls by allowing flexibility with pointers and values.

### `car1.GetCarDescription()` works

In Go

- Methods with **value receivers** can be called on both **values** (e.g., `Car`) and **pointers to values** (e.g., `*Car`)
- When you call a value receiver method on a pointer (`*Car`), Go automatically **dereferences** the pointer for you and invokes the method.

Hence, `car1.GetCarDescription()` effectively is treated like `(*Car1).GetCarDescription()`

### `car2.DisplayDetails()` works

Although `DisplayDetails()` is defined with a **pointer receiver**, `car2.DisplayDetails()` works just fine.
In Go:

- Methods with pointer receivers can also be called on both **pointers** (e.g., `*Car`) and **values** (e.g., `Car`)
- When you call a pointer receiver method on a value (`Car`), Go automatically takes the address of the value and passes a pointer to it.

This implicit referencing is another feature of Go to simplify method calls.

REMEMBER: Go is designed to reduce verbosity and boilerplate code. This flexibility is part of Go's design to make method calls straightforward and less error-prone.

- It eliminates the need for explicit dereferencing (`*`) or taking addresses (`&`) when calling methods.
- It ensures that methods with both pointer and value receivers can be called in an intuitive way.
