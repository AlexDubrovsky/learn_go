package main

import (
	"fmt"
	"sort"
)

type people []string
type numbers []int

func (n numbers) Len() int {
	return len(n)
}

func (n numbers) Less(i, j int) bool {
	return n[i] < n[j]
}

func (n numbers) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	fmt.Println("Sorting slice of strings (user-defined type) with built-in methods")
	studyGroup := people{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(studyGroup)
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	fmt.Println(studyGroup)

	fmt.Println("Sorting slice of strings (user-defined type) with implemented interface")
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	fmt.Println("Sorting slice of strings with built-in methods")
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	// sort.Strings(s)
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)

	fmt.Println("Sorting slice of ints (user-defined type) with implemented interface")
	m := numbers{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Sort(m)
	fmt.Println(m)

	sort.Sort(sort.Reverse(sort.IntSlice(m)))
	fmt.Println(m)

	fmt.Println("Sorting slice of ints with built-in methods")
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Ints(n)
	fmt.Println(n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println(n)
}
