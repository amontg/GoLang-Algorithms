package polynomials

import (
	"fmt"

	csv "github.com/amontg/cs310/createdatacsv"
)

func Evaluate(x float64, coeff []float64) {
	result := coeff[0] + coeff[1]*x
	xPower := x
	n := len(coeff)

	for i := 2; i < n; i++ {
		xPower = xPower * x
		result += coeff[i] * xPower
	}

	fmt.Printf("\nX: %f, Result: %f", x, result)
}

func HornersMethod(x float64, n int, coeff []float64) float64 {
	n = n - 1
	result := coeff[n]
	for i := n - 1; i >= 0; i-- {
		result = result * x
		result = result + coeff[i]
	}

	return result
}

func DoPolynomials() {
	// p(x) = x^7 + 4x^6 + 8x^4 + 6x^3 + 9x^2 + 2x + 3

	coefficients := []float64{3, 2, 9, 6, 8, 0, 4, 1}
	for i := -20.0; i <= 20.0; i += 0.2 {
		Evaluate(i, coefficients)
	}

	resultfile := "hornerspolynomialResults.csv"

	for i := -20.0; i <= 20.0; i += 0.2 {
		prepDivisor := []float64{5, 0, 0, 0, 1}
		prepQ1 := []float64{-1, 0, 1}
		prepQ2 := []float64{4, 1}
		prepQ3 := []float64{12, 1}
		prepR1 := []float64{1, 0, 1}
		prepR2 := []float64{-11, 1}
		prepR3 := []float64{-26, 1}

		div := HornersMethod(i, len(prepDivisor), prepDivisor)
		q1 := HornersMethod(i, len(prepQ1), prepQ1)
		q2 := HornersMethod(i, len(prepQ2), prepQ2)
		q3 := HornersMethod(i, len(prepQ3), prepQ3)
		r1 := HornersMethod(i, len(prepR1), prepR1)
		r2 := HornersMethod(i, len(prepR2), prepR2)
		r3 := HornersMethod(i, len(prepR3), prepR3)
		answer := div*(q1*q2+q3) + (r1*r2 + r3)

		fmt.Printf("\nX: %f, Result: %f", i, answer)

		csv.AddToCSVFile(resultfile, []string{fmt.Sprint(i), fmt.Sprint(answer)})
	}
}
