package main

	import (
		"fmt"
	)
	
	func main() {
		var limit int
		sum := 0
	
		fmt.Print("Enter the upper limit: ")
		fmt.Scan(&limit)
	
		for i := 1; i <= limit; i++ {
			sum += i
		}
	
		fmt.Printf("Sum of all numbers from 1 to %d is: %d\n", limit, sum)
	}
	