package main

import(
	"fmt"
	"structs_example/person"
)

func personExample(){
	p1 := person.NewPerson(1, "Alice", 30)
	fmt.Println(p1) // Output: {1 Alice 30}
	
	// Accessing exported fields
	fmt.Println(p1.Name) // Output: Alice
	p1.Name = "Bob"
	fmt.Println(p1.Name) // Output: Bob
	fmt.Println(p1.Age)  // Output: 30
	p1.Age = 35
	fmt.Println(p1.Age)  // Output: 35
	
	// Accessing unexported field (will cause a compile error if uncommented)
	// fmt.Println(p1.id)
	// p1.id = 2

	person.UpdatePerson(&p1, "Charlie", 40)
	fmt.Println(p1) // Output: {2 Charlie 40}
	
	p2 := &p1 // p2 is a pointer to p1, they point to the same underlying data

	p2.Name = "Dave"
	fmt.Println(p1.Name) // Output: Dave


}

func main(){
	personExample()
}