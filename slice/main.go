package main

import(
	"fmt"
)

// A slice does not store any data, it just describes a section of an underlying array.

// The length of a slice is the number of elements it contains.
// The capacity of a slice is the number of elements in the underlying array, 
// counting from the first element in the slice.

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


func sliceExample(){
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it zero length.
	s = s[:0]
	printSlice(s)
	// Extend its length.
	s = s[:4]
	printSlice(s)
	// Drop its first two values.
	s = s[2:]
	printSlice(s)
	s = s[:4]
	printSlice(s)
	s[0] = 100
	printSlice(s)
}

func createSlice(){
	a:= make([]int, 5) // len(a)=5, cap(a)=5

	printSlice(a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice(b)

	c := b[:2] // len(c)=2, cap(c)=5
	printSlice(c)


	d := c[2:5] // len(d)=3, cap(d)=3
	printSlice(d)

}

func appendSlice(){
	var s []int
	printSlice(s) // len=0 cap=0 []

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s) // len=1 cap=1 [0]

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s) // len=2 cap=2 [0 1]

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s) // len=5 cap=6 [0 1 2 3 4]
}


func main(){
	// sliceExample()
	// createSlice()
	appendSlice()
}