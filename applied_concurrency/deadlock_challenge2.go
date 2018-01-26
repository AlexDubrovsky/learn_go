package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
}

// the original with deadlock
// func main() {
//     c := make(chan int)

//     go func() {
//         for i := 0; i < 10; i++ {
//             c <- i
//         } // must close the channel to continue running the program
//     }()

//     fmt.Println(i) // must be looped over all values on the channel
// }
