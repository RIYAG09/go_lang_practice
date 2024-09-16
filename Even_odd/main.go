package main

import (
	"fmt"
)

func main() {
	var number int

	// Get user input
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)

	result := number % 2

	// Print result: 0 for even, 1 for odd
	fmt.Println(result)
}
