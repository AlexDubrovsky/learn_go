package main

import (
	"fmt"
)

func greatest_num(lst ...float64) float64 {
	res := lst[0]
	for _, v := range lst {
		if v > res {
			res = v
		}
	}

	return res
}

func foo(input ...int) {
	fmt.Println(input)
}

func main() {
	lst := []float64{-1, -5, 10, -20, 8, 6}
	fmt.Println(greatest_num(lst...))

	aSlice := []int{1, 2, 3, 4}
	foo(1, 2)
	foo(1, 2, 3)
	foo(aSlice...)
	foo()
}
