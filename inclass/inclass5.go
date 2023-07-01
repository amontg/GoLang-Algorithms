package inclass

import (
	"fmt"
	"math"
)

func EvalPoly(coefs []float64, x int) {
	var result float64
	input := float64(x)
	for i := 0; i <= len(coefs)-1; i++ {
		result += coefs[i] * (math.Pow(input, coefs[i]))
	}

	fmt.Println(result)
}
