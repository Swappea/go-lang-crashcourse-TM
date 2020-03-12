package main

import "fmt"

func main() {

	// Using var
	// var name = "Swapnesh" // var name string = "Swapnesh" no need to mention string keyword as it is inferred from RHS

	var age = 24 //var age int = 37 - default is int

	// using const
	const isCool = true // cannot change value

	// shorthand
	// name := "Swapnesh" // cannot be defined outside of a function

	// float
	size := 1.3 // default is float64

	name, email := "Swapnesh", "swapneshsangle1@gmail.com"

	fmt.Println(name, age, isCool, size, email) // print
	fmt.Printf("%T\n", size)                    // print the type of variable
}
