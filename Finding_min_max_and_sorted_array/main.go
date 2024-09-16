package main

import (
	"fmt"
	"sort"
)

func findMinMax(arr []int) (int, int) {
	min := arr[0]
	max := arr[0]
	for i := 0; i < len(arr); i++ {
		v := arr[i]
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return min, max
}

func main() {
	var size int
	fmt.Print("Enter the size of the array: ")
	fmt.Scan(&size)

	arr := make([]int, size)
	fmt.Println("Enter the elements:")
	for i := 0; i < size; i++ {
		fmt.Scan(&arr[i])
	}

	min, max := findMinMax(arr)
	sort.Ints(arr)

	fmt.Printf("Minimum = %d\n", min)
	fmt.Printf("Maximum = %d\n", max)
	fmt.Print("Sorted array = ")
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}
