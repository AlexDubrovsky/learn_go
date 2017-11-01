package main

import (
	"fmt"
)

func main() {
	// var student []string
	// var students [][]string

	student := make([]string, 35)
	students := make([][]string, 35)

	// student := []string{}
	// students := [][]string{}

	student[0] = "Alex" // will cause runtime error for var and shorthandhave to use append for them
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil) // will be true for var creation

	a := make([]int, 1)
	fmt.Println(a[0])
	a[0] = 7
	fmt.Println(a[0])
	a[0]++
	fmt.Println(a[0])
}

// 'make' is the common way of defining a slice
