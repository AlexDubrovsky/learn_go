package main

import (
	"fmt"
)

func streak(num int64) int64 {
	var k int64
	for (num+k)%(k+1) == 0 {
		k++
	}

	return k
}

func P(s int64, N int64) int64 {
	var n int64 = 2
	for n < N {
		if streak(n) == s {
			break
		}
		n++
	}

	return n
}

func main() {
	fmt.Println(streak(21))
	fmt.Println(P(3, 14))
}
