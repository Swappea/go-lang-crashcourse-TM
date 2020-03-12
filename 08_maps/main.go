package main

import "fmt"

func main() {
	// Define map
	// emails := make(map[string]string, 0)

	// assign kv
	// emails["Bob"] = "bob@gmail.com"
	// emails["John"] = "john@gmail.com"
	// emails["Jane"] = "jane@gmail.com"

	// declare map and add kv
	emails := map[string]string{"Bob": "bob@gmail.com", "John": "John@gmail.com"}

	emails["Jane"] = "jane@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	// delete from map
	delete(emails, "Bob")
	fmt.Println(emails)
}
