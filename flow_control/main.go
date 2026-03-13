package main

import (
	"fmt"
	"runtime"
	"time"
)

func forLoop(){
	for i := 0; i<5; i++{
		fmt.Println(i)
	}
}


func whileLoop(){
	counter := 0
	for true{
		if counter >= 5{
			break
		}
		fmt.Println(counter)
		counter++
	}
}

func shortStatementIf(){
	// num is only accessible within this if else statement
	if num:= 10; num > 5{
		fmt.Println(num) // you can access num here
	} else if num == 5{
		fmt.Println(num) // you can access num here
	} else{
		fmt.Println(num) // you can access num here
	}
	// fmt.Println(num) // you cannot access num here, it will cause a compile error
}	

func switchCase(){
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func switchWithoutExpression(){
	t := time.Now()
	//Switch without a condition is the same as switch true.
	//This construct can be a clean way to write long if-then-else chains.
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferExample(){
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
	// 4, 3, 2, 1, 0
}

func main() {
	forLoop()
	whileLoop()
	shortStatementIf()
	switchCase()
	switchWithoutExpression()
	deferExample()
}
