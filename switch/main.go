package main

import (
	"fmt"
)

type person struct {
	name string
	age  int32
}

func SwitchOnType(x interface{}) {
	switch x.(type) { // this is an assert; asserting, "x is of this type"
	case int32:
		fmt.Println("int32")
	case string:
		fmt.Println("string")
	case person:
		fmt.Println("person")
	default:
		fmt.Println("unknown")

	}
}

func main() {
	num := 5
	// default behavior - switch breaks after matching case is discovered
	// unlike in C.
	switch num {
	case 1:
		fmt.Println("No 1")
		// fallthrough
	case 3:
		fmt.Println("No 3")
	case 5, 6:
		fmt.Println("Yes 5")
	default:
		fmt.Println("default")
	}

	switch {
	case num == 6:
		fmt.Println("num = 6")
	case num == 5:
		fmt.Println("num = 5")
	}

	var t = person{"Alex", 38}

	SwitchOnType(t.age)
	SwitchOnType(t.name)
	SwitchOnType(t)
}
