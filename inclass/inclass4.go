package inclass

import (
	"fmt"
)

func RadixSort(list []int) []int {
	buckets := make([][]int, len(list))
	for i := 0; i < 10; i++ {
		buckets[i] = make([]int, 10)
	}
	var bucketCounts []int
	keySize := 10

	shift := 1
	for i := 0; i < keySize; i++ {
		bucketCounts = make([]int, 10)
		for j := 0; j < len(list); j++ {
			bucketNumber := (list[j] / shift) % 10
			buckets[bucketNumber][bucketCounts[bucketNumber]] = list[j]
			bucketCounts[bucketNumber]++
		}

		list = CombineBuckets(buckets, bucketCounts)
		shift = shift * 10
	}

	fmt.Println(list)
	return list
}

func CombineBuckets(buckets [][]int, bucketCounts []int) []int {
	N := 0
	for i := 0; i < len(bucketCounts); i++ {
		N += bucketCounts[i]
	}

	combinedBuckets := make([]int, N)
	j := 0

	for bn := 0; bn < len(buckets); bn++ {
		for b := 0; b < bucketCounts[bn]; b++ {
			combinedBuckets[j] = buckets[bn][b]
			j++
		}
	}

	return combinedBuckets
}

func DoRadixSort() {
	RadixSort([]int{6, 2, 4, 7, 1, 3, 8, 5, 10, 9})
}
