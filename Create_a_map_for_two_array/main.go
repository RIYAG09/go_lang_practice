package main

import (
	"fmt"
)

func main() {
	environments := []string{"dev", "qc", "uat", "prod"}
	instances := []string{"t3.micro", "t3.small", "t3.medium", "t3.large"}

	if len(environments) != len(instances) {
		fmt.Println("Arrays must have the same length.")
		return
	}

	envMap := make(map[string]string)
	for i := 0; i < len(environments); i++ {
		envMap[environments[i]] = instances[i]
	}

	fmt.Println("Map:")
	for key, value := range envMap {
		fmt.Printf("%s -> %s\n", key, value)
	}
}
