package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, world!")

	var name string

	fmt.Println("Please enter your name:")
	fmt.Scan(&name)

	fmt.Println("Hello, ", name)
}
