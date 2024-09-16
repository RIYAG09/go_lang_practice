package main

import (
    "fmt"
)

func main() {
    var limit int

    fmt.Print("Enter the upper limit: ")
    fmt.Scan(&limit)

    fmt.Print("Odd numbers: ")

    for i := 1; i <= limit; i += 2 {
        fmt.Print(i)
		fmt.Println()
       
    }
    
     
}
