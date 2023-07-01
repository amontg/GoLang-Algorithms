package polynomials

import "fmt"

func matMulti(A [][]float64, B [][]float64) [][]float64 {
	if len(A[0]) != len(B[0]) {
		fmt.Println("Mismatched arrays.")
		return nil
	}

	result := make([][]float64, len(A))
	for i := 0; i < len(A); i++ {
		result[i] = make([]float64, len(B[0]))
	}

	for i := 0; i < len(A); i++ { // row
		for j := 0; j < len(B[0]); j++ { // column
			for k := 0; k < len(A[0]); k++ {
				result[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return result
}

func DoMatMulti() {
	inputMatrixA := [][]float64{{2, -4, 6}, {6, -6, 6}, {4, 2, 2}}
	inputMatrixB := [][]float64{{-0.2, (1.0 / 6.0), 0.1}, {0.1, (-1.0 / 6.0), 0.2}, {0.3, (-1.0 / 6.0), 0.1}}

	fmt.Println(matMulti(inputMatrixA, inputMatrixB))
}
