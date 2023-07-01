package radixsort

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func RadixStringSort(list []string) []string {
	N := 27
	buckets := make([][]string, N)
	for i := 0; i < N; i++ {
		buckets[i] = make([]string, N)
	}
	var bucketCounts []int
	keySize := 9

	shift := 9
	for i := 0; i < keySize; i++ {
		bucketCounts = make([]int, N)
		for j := 0; j < len(list); j++ {
			list[j] = RemovePadding(list[j])
			if utf8.RuneCountInString(list[j]) < 9 {
				list[j] = AddPadding(list[j], shift)
			}
			bucketNumber := (int(list[j][shift-1])) - 97

			buckets[bucketNumber][bucketCounts[bucketNumber]] = list[j]
			bucketCounts[bucketNumber]++
		}

		list = CombineBuckets(buckets, bucketCounts)
		shift--
	}

	return list
}

func AddPadding(word string, i int) string {
	return fmt.Sprintf("%*s", i, word)
}

func RemovePadding(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

func CombineBuckets(buckets [][]string, bucketCounts []int) []string {
	N := 0
	for i := 0; i < len(bucketCounts); i++ {
		N += bucketCounts[i]
	}

	combinedBuckets := make([]string, N)
	j := 0

	for bn := 0; bn < len(buckets); bn++ {
		for b := 0; b < bucketCounts[bn]; b++ {
			combinedBuckets[j] = buckets[bn][b]
			j++
		}
	}

	return combinedBuckets
}

func DoRadixStringSort() {
	unsorted := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}
	test := RadixStringSort(unsorted)
	fmt.Println(test)
}
