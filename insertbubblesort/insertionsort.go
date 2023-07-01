package insertbubblesort

import (
	"fmt"
	"math"

	ge "github.com/amontg/cs310/generatearray"
)

func InsertionSort(list []int, N int) float64 {
	comparisons := 0.0
	for i := 0; i < N; i++ {
		newElement := list[i]
		location := i - 1

		comparisons++
		for (location >= 0) && (list[location] > newElement) {
			list[location+1] = list[location]
			location = location - 1
			comparisons++
		}
		list[location+1] = newElement
	}
	if len(list) == 8 {
		fmt.Printf("(Insertion) Example List of Size 8: %v\n", list)
	}
	return comparisons
}

func DoInsertionSort() {
	// 2^(1-16)
	tempSlice := []int{6, 2, 4, 7, 1, 3, 8, 5}
	InsertionSort(tempSlice, len(tempSlice))

	for x := 1; x <= 16; x++ {
		tests := 10.0
		N := math.Pow(2, float64(x))
		total_count := 0.0
		for i := 0.0; i < tests; i++ {
			array := ge.RandomizedIntArray(int(N))
			total_count += InsertionSort(array, int(N))
		}
		fmt.Printf("Size: %g, Avg. Comparisons: %g\n", N, total_count/tests)
	}
}
