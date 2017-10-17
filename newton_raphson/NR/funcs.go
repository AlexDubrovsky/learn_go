package NR

import (
	"math"
)

func F(x float64) float64 {
	return math.Pow(x, 2) + 3*x - 4
}

func DF(x float64) float64 {
	return 2*x + 3
}
