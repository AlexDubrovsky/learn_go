package main

import (
	"fmt"
	//"time"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("i = ", i)
			c <- i // the execution is blocked until the value is taken off the channel
		}
		close(c) // close channel after everything is passed to it
	}()

	// loop over the contents of the channel
	for n := range c {
		fmt.Println(n)
	}
}
