package main

import (
    "fmt"
)

func main() {
    for {
        var input1, input2 int
        var choice int

        fmt.Print("Enter first number: ")
        fmt.Scan(&input1)

        fmt.Print("Enter second number: ")
        fmt.Scan(&input2)

        fmt.Println("Choose operation:")
        fmt.Println("1: Add")
        fmt.Println("2: Subtract")
        fmt.Println("3: Multiply")
        fmt.Println("4: Divide")
        fmt.Println("5: Exit")
        fmt.Print("Enter your choice (1-5): ")
        fmt.Scan(&choice)

        if choice == 5 {
            fmt.Println("Exiting...")
            break
        }

        var result float64

        switch choice {
        case 1:
            result = float64(input1 + input2)
        case 2:
            result = float64(input1 - input2)
        case 3:
            result = float64(input1 * input2)
        case 4:
            if input2 != 0 {
                result = float64(input1) / float64(input2)
            } else {
                fmt.Println("Error: Division by zero")
                continue
            }
        default:
            fmt.Println("Invalid choice")
            continue
        }

        fmt.Printf("Result: %.2f\n", result)
    }
}
