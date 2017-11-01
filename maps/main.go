package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println(m)

	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m2)

	// var myGreeting = make(map[string]string)
	var myGreeting = map[string]string{}
	// var myGreeting map[string]string // myGreeting = nil, can't append to it, since no memory is allocated to it
	myGreeting["Tim"] = "Good morning"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)
}
