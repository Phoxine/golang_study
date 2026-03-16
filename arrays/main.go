package main

import (
	"fmt"
)

func main() {
	// An array's length is part of its type, so arrays cannot be resized
	var arr [2]string

	// Assign values to the array
	arr[0] = "Hello"
	arr[1] = "World"

	// Print the array
	fmt.Println(arr)
}