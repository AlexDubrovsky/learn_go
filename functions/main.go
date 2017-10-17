package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("Jane", "Doe"))
	fmt.Println(named_greet("John", "Doe"))
	fmt.Println(mult_greet("John", "Doe"))
	fmt.Println(average(43.5, 56, 100, 89.89, 25.5))
	data := []float64{43.5, 56, 100, 89.89, 25.5}
	fmt.Println(average(data...)) // data... opens the slice and passes the numbers one by one
	// fmt.Println(average(data)) // wworks with average(sf []float64) declaration

	// function expression, anonymous function
	greeting := func() {
		fmt.Println("Hello, world!")
	}

	greeting()
	fmt.Printf("%T\n", greeting)

	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
}

func greet(n1, n2 string) string {
	return fmt.Sprint(n1, n2) // concatenate 2 strings into one
}

// named return
func named_greet(n1, n2 string) (s string) {
	s = fmt.Sprint(n1, n2) // concatenate 2 strings into one
	return
}

// multiple return
func mult_greet(n1, n2 string) (string, string) {
	return fmt.Sprint(n1, n2), fmt.Sprint(n2, n1)
}

// passing variadic parameters
// func average(sf []float64) float64 {
func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)
	var total float64
	// range returns (index, value)
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

// return function that returns string
func makeGreeter() func() string {
	return func() string {
		return "Hello, world, from function!"
	}
}
