package main

import (
	"fmt"
)

func main() {
	mySlice := []int{1, 3, 5, 7, 9, 11}
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice)

	s := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(s)
	fmt.Println(s[2:4])
	fmt.Println(s[2])
	fmt.Printf("%v - %c\n", "myString"[2], "myString"[2])
	fmt.Println(cap(s))

	var s2 = make([]int, 10, 20)
	for i := 0; i < 41; i++ {
		s2 = append(s2, i)
		fmt.Printf("%v - %v - %v\n", s2[i], &s2[i], cap(s2))
	}
}
