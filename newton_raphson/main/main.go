package main

import (
	"fmt"
	"github.com/learn_go/newton_raphson/NR"
)

func main() {
	x0 := 3.5
	fmt.Println(NR.NR(NR.F, NR.DF, x0))
}
