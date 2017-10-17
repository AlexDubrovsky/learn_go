package main

import (
	"fmt"
)

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func double_number(n int) {
	fmt.Println(n * 2)
}

func larger_than_one(n int) bool {
	return n > 1
}

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}

	return xs
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})

	visit([]int{1, 2, 3, 4}, double_number)

	fmt.Println(filter([]int{1, 2, 3, 4}, larger_than_one))
}
