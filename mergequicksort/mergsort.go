package mergequicksort

import (
	"fmt"
)

var comparisons int64 = 0

func MergeSort(list []int, first int, last int) {
	if first < last {
		middle := int((first + last) / 2)
		MergeSort(list, first, middle)
		MergeSort(list, middle+1, last)
		MergeLists(list, first, middle, middle+1, last)
	}
	fmt.Println(list)
}

func MergeLists(list []int, start1 int, end1 int, start2 int, end2 int) []int {
	finalStart := start1
	finalEnd := end2
	indexC := 0
	result := make([]int, len(list))
	for (start1 <= end1) && (start2 <= end2) {
		if list[start1] < list[start2] {
			result[indexC] = list[start1]
			start1 += 1
		} else {
			result[indexC] = list[start2]
			start2 += 1
		}
		comparisons++
		indexC += 1
	}

	if start1 <= end1 {
		for i := start1; i < end2; i++ {
			if indexC <= end2 {
				result[indexC] = list[i]
				indexC += 1
			}
		}
	} else {
		for i := start2; i <= end2; i++ {
			result[indexC] = list[i]
			indexC += 1
		}
	}

	indexC = 0
	for i := finalStart; i <= finalEnd; i++ {
		list[i] = result[indexC]
		indexC += 1
	}

	return result
}

func DoMergeSort() {
	//resultfile := "mergesortResults.csv"
	list := []int{6, 2, 4, 7, 1, 3, 8, 5}
	MergeSort(list, 0, len(list)-1)
	fmt.Printf("Comparisons for array size %d: %d \n", len(list), comparisons)
	/*
		csv.AddToCSVFile(resultfile, []string{"Test Results"})
		csv.AddToCSVFile(resultfile, []string{fmt.Sprint(len(list)), fmt.Sprint(comparisons)})

		csv.AddToCSVFile(resultfile, []string{" "})
		csv.AddToCSVFile(resultfile, []string{"Randomized"})
		csv.AddToCSVFile(resultfile, []string{"Size", "Comparisons"})
		for x := 1; x <= 10; x++ {
			comparisons = 0
			N := int(math.Pow(2, float64(x)))
			array := ga.RandomizedIntArray(N)

			MergeSort(array, 0, len(array)-1)
			fmt.Printf("Comparisons for randomized array, size %d (2^%d): %d \n", len(array), x, comparisons)
			csv.AddToCSVFile(resultfile, []string{fmt.Sprint(len(array)), fmt.Sprint(comparisons)})
		}
		csv.AddToCSVFile(resultfile, []string{" "})
		csv.AddToCSVFile(resultfile, []string{"Ascending"})
		csv.AddToCSVFile(resultfile, []string{"Size", "Comparisons"})
		for x := 1; x <= 10; x++ {
			comparisons = 0
			N := int(math.Pow(2, float64(x)))
			array := ga.AscendingIntArray(N)

			MergeSort(array, 0, len(array)-1)
			csv.AddToCSVFile(resultfile, []string{fmt.Sprint(len(array)), fmt.Sprint(comparisons)})
		}
		csv.AddToCSVFile(resultfile, []string{" "})
		csv.AddToCSVFile(resultfile, []string{"Descending"})
		csv.AddToCSVFile(resultfile, []string{"Size", "Comparisons"})
		for x := 1; x <= 10; x++ {
			comparisons = 0
			N := int(math.Pow(2, float64(x)))
			array := ga.DescendingIntArray(N)

			MergeSort(array, 0, len(array)-1)
			csv.AddToCSVFile(resultfile, []string{fmt.Sprint(len(array)), fmt.Sprint(comparisons)})
		}
	*/
}
