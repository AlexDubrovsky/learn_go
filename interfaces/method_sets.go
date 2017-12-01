package main

import (
	"fmt"
)

// Square - exported type
type Square struct {
	side float64
}

// if the receiver is value - info() will accept both value and pointer
// if the receiver is pointer - will accept only pointers
func (z Square) area() float64 {
	return z.side * z.side
}

// Shape - exported type
type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println("Area ", z.area())
}

func main() {
	s := Square{10}
	info(&s)

	s.side = 5
	info(&s)
}
