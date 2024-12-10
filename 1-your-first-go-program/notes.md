# GO Language Basics

To run the program type `go run main.go` in the command line.

Let's break down the anatomy of a Go program:

## Packages

1. Every Go file starts with a `package` declaration. `main` is a special package for executable programs.

2. The `import` statement brings other packages into your program; for instance, `fmt` is a standard library package for formatted I/O operations like printing to the console.

## package main

This specifies that the file belongs to the `main` package. The `main` package is special in Go: it defines a standalone eecutable program, not a library. When you run a Go program (`go run` or `go build`), the `main` package is where execution starts.

## func main()

This is the entry point of a Go program.

When you execute your program, the Go runtime looks for the `main()` function in the `main` package and starts execution from there.

N.B. A program must have **exactly one** `main()` **function** in the `main package`, or the compiler will throw an error.

## Other Functions

You can define additional helper functions (besides `main()`), and they can be called from `main()` or other functions.
