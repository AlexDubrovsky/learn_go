package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(i, "*", j, "=", i*j)
		}
	}

	// "while" loop in go
	k := 0
	for k < 10 {
		fmt.Println(k)
		k++
	}

	// infinite loop
	// for {
	// 	fmt.Println(k)
	// 	k++
	// }

	// same as "for cond {}"
	for {
		fmt.Println(k)
		if k >= 10 {
			break
		}
		k++
	}

	// continue and break
	m := 0
	for {
		m++
		if m%2 == 0 {
			continue
		}
		fmt.Println(m)
		if m >= 50 {
			break
		}
	}
}
