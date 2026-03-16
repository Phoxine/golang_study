package main

import "fmt"

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *Tree, ch chan int){
	if t == nil{
		return
	}
	// in-order traversal
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *Tree) bool{
	c1 := make(chan int)
	c2 := make(chan int)
	
	go Walk(t1, c1)
	go Walk(t2, c2)

	for i:=0; i<10; i++{
		if <-c1 != <-c2{
			return false
		}
	}
	return true

}

func EquivalentBinaryTrees() {
	fmt.Println(Same(NewTree(1), NewTree(1)))
	fmt.Println(Same(NewTree(1), NewTree(2)))
}