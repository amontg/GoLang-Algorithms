package insertbubblesort

import (
	"fmt"
	"math"

	ge "github.com/amontg/cs310/generatearray"
)

func BubbleSort(list []int, N int) int {
	numberOfPairs := N
	swappedElements := true
	comparisons := 0

	for swappedElements == true {
		numberOfPairs--
		swappedElements = false

		for i := 0; i < numberOfPairs; i++ {
			if list[i] > list[i+1] {
				list = ge.Swap(list, i, 1)
				swappedElements = true
			}
			comparisons++
		}
	}
	if len(list) == 8 {
		fmt.Printf("(Bubble) Example List of Size 8: %v\n", list)
	}
	return comparisons
}

func DoBubbleSort() {
	tempSlice := []int{6, 2, 4, 7, 1, 3, 8, 5}
	BubbleSort(tempSlice, len(tempSlice))

	for x := 1; x <= 16; x++ {
		tests := 10.0
		N := math.Pow(2, float64(x))
		total_count := 0.0
		for i := 0.0; i < tests; i++ {
			array := ge.RandomizedIntArray(int(N))
			total_count += float64(BubbleSort(array, int(N)))
		}
		fmt.Printf("Size: %g, Avg. Comparisons: %g\n", N, total_count/tests)
	}
}
