package polynomialfit

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	csvme "github.com/amontg/cs310/createdatacsv"
	gj "github.com/amontg/cs310/matrixes"
	poly "github.com/amontg/cs310/polynomials"
)

func ReadNoisyData() (xvals, yvals []float64) {
	file, err := os.Open("NoisyPolynomialData.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)
	records, _ := reader.ReadAll()

	for i := 0; i < len(records); i++ {
		tempX, err := strconv.ParseFloat(strings.TrimSpace(records[i][0]), 64)
		if err != nil {
		}

		tempY, err := strconv.ParseFloat(strings.TrimSpace(records[i][1]), 64)
		if err != nil {
		}
		xvals = append(xvals, tempX)
		yvals = append(yvals, tempY)
	}

	return xvals, yvals
}

func SumXDegreeY(values []float64, y float64) float64 {
	var sum float64
	for i := 0; i < len(values); i++ {
		sum = sum + math.Pow(values[i], y)
	}
	return sum
}

func DoubleSummationDegreeY(xvals, yvals []float64, y float64) float64 {
	var sum float64
	for i := 0; i < len(xvals); i++ {
		sum += yvals[i] * (math.Pow(xvals[i], y))
	}

	return sum
}

func CreateMatrixQ(values []float64, degree int) [][]float64 {
	matrix := make([][]float64, degree+1)
	for i := 0; i < degree+1; i++ {
		matrix[i] = make([]float64, degree+1)
	}

	rows := len(matrix)
	cols := len(matrix[0])
	N := float64(len(values))

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if r == 0 && c == 0 {
				matrix[r][c] = N
			} else {
				matrix[r][c] = SumXDegreeY(values, float64(c+r))
			}
		}
	}

	// gj.PrintMatrix(matrix)
	// fmt.Println("")
	return matrix
}

func CreateMatrixU(xvals, yvals []float64, degree int) [][]float64 {
	matrix := make([][]float64, degree+1)
	for i := 0; i < degree+1; i++ {
		matrix[i] = make([]float64, 1)
	}

	rows := len(matrix)
	cols := len(matrix[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			matrix[r][c] = DoubleSummationDegreeY(xvals, yvals, float64(r))
			//matrix[r][c] = float64(r)
		}
	}

	// gj.PrintMatrix(matrix)
	// fmt.Println("")
	return matrix
}

func MultiplyMatrices(q [][]float64, u [][]float64) [][]float64 {
	// multiply q[row 0-X] by u[columns 0-X] and add them
	q = gj.GJInversion(q)
	matrix := make([][]float64, len(u))
	for i := 0; i < len(matrix); i++ {
		matrix[i] = make([]float64, len(u[0]))
	}
	/*
		for each row in q do
			for each column in q do
				answer = q[row][column] * u[row 0-2][0]
	*/

	for r := 0; r < len(q); r++ {
		for c := 0; c < len(q); c++ {
			matrix[r][0] += q[r][c] * u[c][0]
		}
	}

	//gj.PrintMatrix(matrix)
	return matrix
}

func PolynomialFit() {
	xvals, yvals := ReadNoisyData()

	resultfile := "NoisyPolynomialDataResults.csv"

	for d := 1; d <= 7; d++ {
		csvme.AddToCSVFile(resultfile, []string{" "})
		csvme.AddToCSVFile(resultfile, []string{fmt.Sprint("Degree ", d)})

		for x := -20.0; x <= 20.0; x += 0.2 {
			matrixQ := CreateMatrixQ(xvals, d)
			matrixU := CreateMatrixU(xvals, yvals, d)

			matrixC := MultiplyMatrices(matrixQ, matrixU)
			tempMatrix := make([]float64, len(matrixC))
			for i := 0; i < len(matrixC); i++ {
				tempMatrix[i] = matrixC[i][0]
			}

			answer := poly.HornersMethod(x, len(tempMatrix), tempMatrix)
			csvme.AddToCSVFile(resultfile, []string{fmt.Sprint(answer)})
		}
	}
}
