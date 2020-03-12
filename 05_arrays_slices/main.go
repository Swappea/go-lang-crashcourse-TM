package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2]string

	// Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare and assign
	fruitsArr := [2]string{"Apple", "Watermelon"}

	// slices
	fruitsSlice := []string{"Apple", "Watermelon", "Grapes", "Banana", "Avocado"}

	fmt.Println(fruitArr)
	fmt.Println(fruitsArr)
	fmt.Println(fruitsSlice)
	fmt.Println(len(fruitsSlice))
	fmt.Println(fruitsSlice[1:4])
}
