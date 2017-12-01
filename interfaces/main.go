package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (z Circle) area() float64 {
	return math.Pi * z.radius * z.radius
}

func (z Circle) circumference() float64 {
	return 2.0 * math.Pi * z.radius
}

func (z Square) area() float64 {
	return z.side * z.side
}

func (z Square) circumference() float64 {
	return 4.0 * z.side
}

type Shape interface {
	area() float64
	circumference() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println("Area ", z.area())
	fmt.Println("Circumference ", z.circumference())
}

func main() {
	s := Square{10}
	info(s)
	fmt.Println(s.area())

	c := Circle{10}
	info(c)
}
