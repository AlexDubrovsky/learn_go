package main

import (
	"fmt"
)

func hello() {
	fmt.Print("hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	// defer the execution of world() to right before main exits
	// defer world()
	world()
	hello()
}
