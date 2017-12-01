// concurrency
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2) // add 2 threads to WaitGroup
	go foo()
	go bar()
	wg.Wait() // wait until no items in WaitGroup
}

func foo() {
	for i := 0; i < 20; i++ {
		fmt.Println("Foo: ", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done() // thread is done, one item is removed from WaitGrout
}

func bar() {
	for i := 0; i < 20; i++ {
		fmt.Println("Bar: ", i)
		time.Sleep(time.Duration(5 * time.Millisecond))
	}
	wg.Done() // thread is done, one item is removed from WaitGrout
}
