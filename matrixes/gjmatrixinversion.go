package matrixes

import (
	"fmt"
	"math"
)

/*
	Matrix Notes

	1) Multiply/divide a row by a scalar (== the PV)
		You want the pivot value to be 1. If PV == 1, move to the next step. If PV == 0, find a row without PV == 0 and swap.
	2) Exchange 2 rows
	3) Subtract multiples of one row from another replacing the first
	4) If at any time the matrix fails to be square, bail out.
	You have a "SINGULAR" matrix

	For each row
	- make pivot != 0 (exchange rows)
		if inputMatrix[0][0] == 0 then (they would be [0][0], [1][1], [2][2] for the diagonals)
			for each row do
				if inputMatrix[foundrow][0] != 0 then
					SwapPlaces(inputMatrix, iM[0][0], iM[foundrow][0])
					break
				end if
			end for
		end if

	- make pivot = 1 (divide by pivot)
		if inputMatrix[0][0] != 1 then ([0][0], [1][1], [2][2] diagonals)
			tempDivisor = inputMatrix[0][0]
			inputMatrix[0][0] = inputMatrix[0][0]/inputMatrix[0][0]
			for each column do
				inpuxMatrix[0][i] = inputMatrix[0][i]/inputMatrix[0][0]
			end for
		end if
	- subtract multiples of pivot row from all other rows to set pivot column = 0
		man what???
*/

func GJInversion(A [][]float64) [][]float64 {
	rows := len(A)
	cols := len(A[0])

	I := MakeIdentityMatrix(rows)

	for i := 0; i < rows; i++ { // i = row
		pivotloc := i
		pivot := A[pivotloc][i]
		for math.Abs(pivot) < 0.0000000001 {
			pivotloc++
			pivot = A[pivotloc][i]
		}

		SwapRows(A, i, pivotloc)
		SwapRows(I, i, pivotloc)

		for j := 0; j < cols; j++ {
			A[i][j] /= pivot
			I[i][j] /= pivot
		}

		for ii := 0; ii < rows; ii++ {
			if i == ii {
				continue
			}

			multiplier := A[ii][i]
			for j := 0; j < cols; j++ {
				A[ii][j] -= multiplier * A[i][j]
				I[ii][j] -= multiplier * I[i][j]
			}
		}
	}

	//PrintMatrix(A)
	//fmt.Println("")
	return I
}

func SwapRows(A [][]float64, r0, r1 int) {
	temp := A[r0]
	A[r0] = A[r1]
	A[r1] = temp
}

func MakeIdentityMatrix(size int) [][]float64 {
	result := make([][]float64, size)
	for i := 0; i < size; i++ {
		result[i] = make([]float64, size)
		result[i][i] = 1.0
	}
	return result
}

func PrintMatrix(A [][]float64) {
	for i := 0; i < len(A); i++ {
		fmt.Printf("%v\n", A[i])
	}
}

func DoGJInversion() {
	matrix := [][]float64{{6, -6, 6}, {0, -4, 6}, {4, 2, 2}}

	PrintMatrix(GJInversion(matrix))
}
