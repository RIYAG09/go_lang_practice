package main

import (
	"fmt"
)

func main() {
	var age int
	var citizenship string

	// Get user input
	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Print("Are you a citizen of India? (yes/no): ")
	fmt.Scan(&citizenship)

	// Check eligibility
	if age >= 18 && (citizenship == "yes" || citizenship == "Yes") {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}
}
