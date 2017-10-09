package main

import (
	"fmt"
	"reflect"
)

func main() {

	n, err := fmt.Println("1. Hello world!")
	fmt.Println("2. Hello world!")
	fmt.Println("3. Hello world!")
	fmt.Println("4. Hello git!")

	fmt.Println("n = ", n)
	fmt.Println("err = ", err)
	fmt.Println("type of err is ", reflect.TypeOf(err))
}
