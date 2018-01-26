package main

import (
	"fmt"
)

// func main() {
//  c := make(chan int)
//  go func(c chan int) {
//      fmt.Println("Taking off the channel", <-c)
//  }(c)
//  c <- 1
//  close(c)
// }

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}

// original with deadlock

// func main() {
//     c := make(chan int)
//     c <- 1
//     fmt.Println(<-c)
// }

// channel c blocks until something takes off it.
// the program will not go to the next line because there is nothing to take off the channel
