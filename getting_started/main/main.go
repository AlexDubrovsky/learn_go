package main

import (
	"fmt"
	"github.com/learn_go/getting_started/stringutil"
	"io/ioutil"
	"net/http"
)

var x int = 42

func main() {
	fmt.Println(42)
	fmt.Printf("%d - %b\n", 42, 42)
	fmt.Printf("%d - %X\n", 42, 42)

	// for i := 0; i < 200; i++ {
	// 	fmt.Printf("%d - %b - %x - %q\n", i, i, i, i)
	// }
	// s := "Alex"
	fmt.Println(stringutil.Name)
	fmt.Println(stringutil.Reverse(stringutil.Name))
	fmt.Printf("%T\n", stringutil.Name)

	fmt.Printf("From main - %d\n", x)
	foo()

	y := 0

	// anonymous function, assigned to inc
	inc := func() int {
		y++
		return y
	}

	fmt.Println(inc())
	fmt.Printf("y = %d\n", y)
	fmt.Println(inc())
	fmt.Printf("y = %d\n", y)

	fmt.Println("\nTest for wrapper")
	incr := wrapper()
	fmt.Printf("z = %d\n", incr())
	fmt.Printf("z = %d\n", incr())
	fmt.Printf("z = %d\n", incr())

	fmt.Println("HTTP example")
	http_example()
}

func wrapper() func() int {
	z := 5

	return func() int {
		z++

		return z
	}
}

func foo() {
	fmt.Printf("From foo - %d\n", x)
}

func http_example() {
	res, _ := http.Get("http://www.google.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}
