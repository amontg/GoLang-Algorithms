package searchmethods

import (
	"fmt"
	"math"
)

func FibSearch(target int, N int, list []int) (int, int) {
	// degenerate case
	if N == 1 && list[N-1] == target {
		return N - 1, 0
	}

	// fibonacci numbers
	m2 := 0
	m1 := 1
	m := 1

	for m < N {
		m2 = m1
		m1 = m
		m = m2 + m1
	}

	start := -1

	count := 0
	for m > 1 {
		var index int
		if (start + m2) < (N - 1) {
			index = start + m2
		} else {
			index = N - 1
		}

		count++
		if target > list[index] {
			m = m1
			m1 = m2
			m2 = m - m1
			start = index
		} else {
			count++
			if target < list[index] {
				m = m2
				m1 = m1 - m2
				m2 = m - m1
			} else {
				return index, count
			}
		}
	}

	return -1, count
}

func DoFibSearch() {
	for x := 1; x <= 20; x++ {
		var fibTable []int
		N := int(math.Pow(2, float64(x)))
		for i := 0; i < N; i++ {
			fibTable = append(fibTable, i)
		}

		avg := 0.0
		for y := 1; y < len(fibTable); y++ {
			target := y

			//fmt.Println(" ")
			//fmt.Println("Target: ", target)
			_, count := FibSearch(target, N, fibTable)
			//fmt.Println("Final Key: ", key, " | Value: ", binaryTable[key])
			//fmt.Println(binaryTable)
			avg += float64(count)
		}
		//fmt.Println(" ")
		//fmt.Println("Size: ", N, " | Exponent: ", x)
		fmt.Printf("Size: %d, Avg. Comparisons: %v\n", N, avg/float64(N))
	}
}
