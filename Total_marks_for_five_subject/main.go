package main

import (
	"fmt"
)

func main() {
	var marks1, marks2, marks3, marks4, marks5 float64

	fmt.Println("Enter marks for five subjects:")

	fmt.Print("Subject 1: ")
	fmt.Scan(&marks1)

	fmt.Print("Subject 2: ")
	fmt.Scan(&marks2)

	fmt.Print("Subject 3: ")
	fmt.Scan(&marks3)

	fmt.Print("Subject 4: ")
	fmt.Scan(&marks4)

	fmt.Print("Subject 5: ")
	fmt.Scan(&marks5)

	totalMarks := marks1 + marks2 + marks3 + marks4 + marks5

	fmt.Printf("Total marks: %.2f\n", totalMarks)
}
