package main

import (
	"fmt"
)

type User struct {
	Name     string
	Age      int
	Email    string
	Location string
}

func main() {
	users := []User{
		{Name: "Alice", Age: 30, Email: "alice@example.com", Location: "New York"},
		{Name: "Bob", Age: 25, Email: "bob@example.com", Location: "San Francisco"},
		{Name: "Charlie", Age: 35, Email: "charlie@example.com", Location: "Chicago"},
	}

	for i := 0; i < len(users); i++ {
		user := users[i]
		fmt.Printf("Name: %s, Age: %d, Email: %s, Location: %s\n", user.Name, user.Age, user.Email, user.Location)
	}
}
