package main

import (
	"fmt"
)

func check_even(x int) bool {
	return x%2 == 0
}

func main() {
	b := true

	if food := "Beer"; b {
		fmt.Println(food)
	}

	z := 10
	if res := check_even(z); res {
		fmt.Println("z = ", z, " is even")
	} else {
		fmt.Println("z = ", z, " is odd")
	}
}
