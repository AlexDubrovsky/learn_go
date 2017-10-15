package main

import (
	"fmt"
)

func main() {
	var num1 int32
	var num2 int32

	fmt.Println("Enter two numbers:")
	fmt.Scan(&num1, &num2)

	if num1 < num2 {
		num1, num2 = num2, num1
	}

	fmt.Println("The remainder of ", num1, "/", num2, " is ", num1%num2)
}
