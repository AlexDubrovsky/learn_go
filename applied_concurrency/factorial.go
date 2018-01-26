package main

import (
	"fmt"
)

// factorial() calculates the product of numbers from start to end
// sends the result to channel out, indicates it's done by sending to channel done
func factorial(start uint64, end uint64, out chan uint64, done chan bool) {
	go func() {
		total := uint64(1)
		for i := start; i < end; i++ {
			total *= i
		}
		// fmt.Println(total)
		out <- total
		done <- true
	}()
}

func main() {
	n := uint64(20)
	g := uint64(2)
	res := uint64(1)
	var last uint64
	c := make(chan uint64)
	done := make(chan bool)

	// calculate start and end, send to goroutine
	for i := uint64(0); i < g; i++ {
		start := i*(n/g) + 1
		end := (i+1)*(n/g) + 1
		last = end
		factorial(start, end, c, done)
	}

	// if the division to parts is unequal, calculate the rest of the factorial
	if last <= n {
		g++ // increment goroutines counter in case another one was needed
		factorial(last, n+1, c, done)
	}

	// blocked until all goroutines are finished, then close channel
	go func() {
		for i := uint64(0); i < g; i++ {
			<-done
		}
		close(c)
	}()

	// calculate the final result
	for i := range c {
		res *= i
	}
	fmt.Printf("%v! = %v\n", n, res)
}
