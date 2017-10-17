package NR

import (
	"math"
)

var conv_eps float64 = 1e-8

func NR(F func(float64) float64,
	DF func(float64) float64,
	x0 float64) (root float64) {

	eps := 1.0

	for eps > conv_eps {
		root = x0 - F(x0)/DF(x0)
		eps = math.Abs(root-x0) / math.Abs(root)
		x0 = root
	}

	return
}
