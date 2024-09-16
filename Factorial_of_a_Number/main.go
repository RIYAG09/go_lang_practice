package main

import (
	"fmt"
)
	
func factorial (n int ) int{
	if n == 0 {
		return 1
	}
	return n * factorial(n - 1)
}

func main(){
	var num int
	fmt.Print("Enter the number: ")
	fmt.Scan(&num)

	if num < 0 {
		fmt.Println("Factorial does not exist for negative numbers")
		return
	}

	result := factorial(num)
	fmt.Println("Factorial of", num, "is", result)
}