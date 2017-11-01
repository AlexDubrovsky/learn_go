package main

import (
	"fmt"
)

func main() {
	var x [58]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	for i := 65; i <= 122; i++ {
		x[i-65] = string(i)
	}

	for i, v := range x {
		fmt.Printf("%v - %T - %v\n", v, v, []byte(v))
		if i > 50 {
			break
		}
	}

	var y [256]int
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(y[42])

	for i := 0; i < 256; i++ {
		y[i] = i
	}

	for i, v := range y {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}

	var z [256]byte
	fmt.Println(len(z))
	fmt.Println(z[0])
	for i := 0; i < 256; i++ {
		z[i] = byte(i)
	}

	for i, v := range z {
		fmt.Printf("%v - %T - %b\n", v, v, v)
		if i > 50 {
			break
		}
	}
}
