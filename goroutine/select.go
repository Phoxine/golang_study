
package main

import "fmt"

// The select statement lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

func fibonacci_select(c, quit chan int) {
	x, y := 0, 1
	for {
		// select is used to wait on multiple channel operations.
		select {
		// send x to channel c if there is a receiver ready to receive the value from channle c.
		case c <- x:
			x, y = y, x+y
		// receive quit signal from channel quit and terminate the function
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func SelectExample() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			// keep blocking until we receive a value from the channel c
			fmt.Println(<-c)
		}
		// notify the fibonacci function to quit
		quit <- 0
	}()
	fibonacci_select(c, quit)
}
