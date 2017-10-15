package main

import (
	"fmt"
)

func main() {
	fmt.Println([]rune("Hello מ"))
	fmt.Println([]byte("Hello מ"))

	for i := 50; i <= 140; i++ {
		fmt.Println(i, "-", string(i), "-", []byte(string(i)), "-", []rune(string(i)))
		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
	}

	foo := "a" // string
	bar := 'a' // rune
	my_str := "my string"
	fmt.Println(foo)
	fmt.Println(bar)
	fmt.Println(my_str)
	fmt.Printf("%s \n", foo)
	fmt.Printf("%c \n", bar)
	fmt.Printf("%s \n", my_str)
	fmt.Printf("%T \n", foo)
	fmt.Printf("%T \n", bar)
	fmt.Printf("%T \n", my_str)
	// fmt.Printf("%T \n", rune(foo))
}
