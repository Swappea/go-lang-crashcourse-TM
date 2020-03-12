package main

import (
	"fmt"
	"strconv"
)

// Person ... define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method(value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old"
}

// hasBirthday method(pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getmarried(pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	}
	p.lastName = spouseLastName
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Jane", lastName: "Doe", city: "NY", gender: "F", age: 23}

	// Alternative
	person2 := Person{"John", "Doe", "NY", "M", 23}

	// person1.age++
	// fmt.Println(person1)
	// fmt.Println(person1.firstName)
	person1.hasBirthday()
	person1.getMarried("Williams")
	fmt.Println(person1.greet())

	person2.hasBirthday()
	person2.getMarried("Thompson")
	fmt.Println(person2.greet())
}
