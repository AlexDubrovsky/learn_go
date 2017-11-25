package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	m["k3"] = 45
	m["k4"] = 78

	fmt.Println(m)

	for key, val := range m {
		fmt.Println(key, " - ", val)
	}

	m2 := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(m2)

	// var myGreeting = make(map[string]string)
	var myGreeting = map[string]string{}
	// var myGreeting map[string]string // myGreeting = nil, can't append to it, since no memory is allocated to it
	myGreeting["Tim"] = "Good morning"
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)
	fmt.Println(len(myGreeting))

	m3 := map[int]int{
		1: 2,
		2: 4,
	}
	m3[3] = 6
	fmt.Println(m3)
	// delete(m3, 3)
	// fmt.Println(m3)

	if val, exists := m3[3]; exists {
		delete(m3, 3)
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}
}
