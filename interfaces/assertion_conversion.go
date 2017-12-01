package main

import (
	"fmt"
)

func main() {
	var name interface{} = "Sydney"
	str, ok := name.(string)

	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Printf("Value is not a string")
	}

	var val interface{} = 7
	fmt.Printf("%T\n", val)
}
