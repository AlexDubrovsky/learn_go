package main

import (
	"fmt"
	"math"
)

const (
	A        = iota
	B        = iota
	C        = iota
	q        = 42
	p string = "death & taxes"
)

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
)

func main() {

	fmt.Println("A - ", A)
	fmt.Println("B - ", B)
	fmt.Println("C - ", C)
	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%T\n", q)
	fmt.Println(math.Sqrt(q))
}
