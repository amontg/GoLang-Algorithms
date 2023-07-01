package searchmethods

import (
	"fmt"
	"math"
)

func BinarySearch(target int, N int, list []int) (int, int) {
	start := 1
	end := N
	count := 0
	for start <= end {
		middle := (start + end) / 2
		count++

		switch Compare(list[middle], target) {
		case -1:
			start = middle + 1
			break
		case 0:
			//fmt.Println("Comparison Count: ", count)
			return middle, count
		case 1:
			end = middle - 1
			break
		}
	}

	return 0, count
}

func Compare(x int, y int) int {
	var res int
	if x < y {
		res = -1
	} else if x > y {
		res = 1
	} else {
		res = 0
	}

	return res
}

func DoBinarySearch() {
	for x := 1; x <= 20; x++ {
		var binaryTable []int
		N := int(math.Pow(2, float64(x)))
		for i := 0; i < N; i++ {
			binaryTable = append(binaryTable, i)
		}

		avg := 0.0
		for y := 1; y < len(binaryTable); y++ {
			target := y

			//fmt.Println(" ")
			//fmt.Println("Target: ", target)
			_, count := BinarySearch(target, N, binaryTable)
			//fmt.Println("Final Key: ", key, " | Value: ", binaryTable[key])
			//fmt.Println(binaryTable)
			avg += float64(count)
		}
		//fmt.Println(" ")
		//fmt.Println("Size: ", N, " | Exponent: ", x)
		fmt.Println(N, avg/float64(N))
	}
}
