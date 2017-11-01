package main

import (
	"fmt"
)

func two_returns(num int64) (int64, bool) {

	output := func(num int64) (int64, bool) {
		return num / 2, num%2 == 0
	}

	return output(num)
}

func main() {
	fmt.Println(two_returns(6))
	fmt.Println(two_returns(7))
}
