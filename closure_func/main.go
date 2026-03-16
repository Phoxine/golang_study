package main

import(
	"fmt"
)

// A closure is a function value that references variables from outside its body. 
// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func fib_recursive(n int) int {
	var fib func(int) int
	var visited = make(map[int]int)
    fib = func(n int) int {
        if n <= 1 {
            return n
        }
        if val, ok := visited[n]; ok {
            return val
        }
        result := fib(n-1) + fib(n-2)
        visited[n] = result
        return result
    }

    return fib(n)
}

func main() {
	// f := fibonacci()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(f())
	// }

	fmt.Println(fib_recursive(6))
}
