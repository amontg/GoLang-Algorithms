package insertbubblesort

import (
	"fmt"
	"math"

	ge "github.com/amontg/cs310/generatearray"
)

func AltBubbleSort(list []int, N int) int {
	numberOfPairs := N
	swappedElements := true
	comparisons := 0

	for swappedElements == true {
		numberOfPairs--
		swappedElements = false

		direction := true
		if direction == true {
			for i := 0; i < numberOfPairs; i++ {
				if list[i] > list[i+1] {
					list = ge.Swap(list, i, 1)
					swappedElements = true
				}
				comparisons++
			}
		} else {
			for i := len(list) - 1; i > numberOfPairs; i-- {
				if list[i-1] > list[i] {
					list = ge.Swap(list, i, -1)
					swappedElements = true
				}
				comparisons++
			}
		}
		direction = !direction
	}
	if len(list) == 8 {
		fmt.Printf("Example List of Size 8: %v\n", list)
	}
	return comparisons
}

func DoAltBubbleSort() {
	tempSlice := []int{6, 2, 4, 7, 1, 3, 8, 5}
	AltBubbleSort(tempSlice, len(tempSlice))

	for x := 1; x <= 16; x++ {
		tests := 10.0
		N := math.Pow(2, float64(x))
		total_count := 0.0
		for i := 0.0; i < tests; i++ {
			array := ge.RandomizedIntArray(int(N))
			total_count += float64(AltBubbleSort(array, int(N)))
		}
		fmt.Printf("Size: %g, Avg. Comparisons: %g\n", N, total_count/tests)
	}
}
