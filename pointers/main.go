package main

import (
	"fmt"
)

const (
	metersToyards float64 = 1.09361
)

func zero(x int) {
	x = 0
}

func zeroPtr(x *int) {
	fmt.Println("In zeroPtr ", &x)
	*x = 0
}

func main() {
	x := 4
	var y *int = &x

	fmt.Printf("%p %d\n", y, *y)

	z := *y + 1
	fmt.Println(z)
	fmt.Printf("%T\n", y)

	// var meters float64
	// fmt.Print("Enter meters swam: ")
	// fmt.Scan(&meters)
	// yards := meters * metersToyards
	// fmt.Println(meters, " meters is ", yards, " yards.")

	zero(z)
	fmt.Println(z)

	zeroPtr(&z)
	fmt.Println(z)
	fmt.Println("In main ", &z)

	a := 13 % 2
	fmt.Println(a)

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("Even ", i)
		} else {
			fmt.Println("Odd ", i)
		}
	}
}
