package main

import (
	"fmt"
)

func main() {
	m := make([]string, 1, 25)
	fmt.Println(&m[0])
	fmt.Println(m)
	changeMe(m)
	fmt.Println(m)

	n := make(map[string]int)
	changeMeMap(n)
	fmt.Println(n["Alex"])

	// anonymous self-executing function
	func() {
		fmt.Println("What is this sorcery?")
	}()

}

func changeMe(z []string) {
	z[0] = "Alex"
	// z[1] = "Dubrovsky"
	fmt.Println(&z[0])
	fmt.Println(z)
}

func changeMeMap(z map[string]int) {
	z["Alex"] = 38
}
