package main

import (
	"fmt"
)

func changeMe(z *string) {
	fmt.Println(&z)
	fmt.Println(z)
	fmt.Println(*z)

	*z = "alex"
	fmt.Println(z)
	fmt.Println(*z)
}

func main() {
	age := "Alex"
	fmt.Println(&age)
	changeMe(&age)

	fmt.Println(&age)
	fmt.Println(age)
}
