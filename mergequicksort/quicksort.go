package mergequicksort

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	csv "github.com/amontg/cs310/createdatacsv"
	ga "github.com/amontg/cs310/generatearray"
)

var qcomparisons int64 = 0

func Quicksort(list []int, first, last, pivot int) {
	if first < last {
		pivot := PivotList(list, first, last, pivot)
		Quicksort(list, first, pivot-1, pivot)
		Quicksort(list, pivot+1, last, pivot)
	}
	//fmt.Println(list)
}

func PivotList(list []int, first, last, pivot int) int {

	switch pivot { // change this for different pivot locations
	case 0:
		// do nothing, this just for pivot = first
	case 1:
		ga.SwapPlaces(list, first, last) // pivot = last
	case 2:
		rand.Seed(time.Now().Unix())
		random := rand.Intn(last-first+1) + first
		ga.SwapPlaces(list, first, random)
	}

	PivotValue := list[first]
	PivotPoint := first

	for i := first + 1; i <= last; i++ {

		qcomparisons++
		if list[i] < PivotValue {
			PivotPoint++
			list = ga.SwapPlaces(list, i, PivotPoint)
			//fmt.Println(list)
		}
	}

	list = ga.SwapPlaces(list, first, PivotPoint)
	//fmt.Println(list)
	return PivotPoint
}

func DoQuickSort() {
	resultfile := "quicksortResults.csv"
	list := []int{6, 2, 4, 7, 1, 3, 8, 5}
	Quicksort(list, 0, len(list)-1, 2)
	fmt.Printf("Comparisons for array size %d: %d \n", len(list), qcomparisons)
	csv.AddToCSVFile(resultfile, []string{"Test Results"})
	csv.AddToCSVFile(resultfile, []string{fmt.Sprint(len(list)), fmt.Sprint(qcomparisons)})
	csv.AddToCSVFile(resultfile, []string{" "})

	for pivot := 0; pivot <= 2; pivot++ {
		csv.AddToCSVFile(resultfile, []string{fmt.Sprintf("Pivot %d", pivot)})
		csv.AddToCSVFile(resultfile, []string{" "})
		csv.AddToCSVFile(resultfile, []string{"Randomized"})
		csv.AddToCSVFile(resultfile, []string{"Size", "Comparisons"})
		for x := 1; x <= 12; x++ {
			qcomparisons = 0
			N := int(math.Pow(2, float64(x)))
			array := ga.RandomizedIntArray(N)

			Quicksort(array, 0, len(array)-1, pivot)
			csv.AddToCSVFile(resultfile, []string{fmt.Sprint(len(array)), fmt.Sprint(qcomparisons)})
		}
		csv.AddToCSVFile(resultfile, []string{" "})
		csv.AddToCSVFile(resultfile, []string{"Ascending"})
		csv.AddToCSVFile(resultfile, []string{"Size", "Comparisons"})
		for x := 1; x <= 12; x++ {
			qcomparisons = 0
			N := int(math.Pow(2, float64(x)))
			array := ga.AscendingIntArray(N)

			Quicksort(array, 0, len(array)-1, pivot)
			csv.AddToCSVFile(resultfile, []string{fmt.Sprint(len(array)), fmt.Sprint(qcomparisons)})
		}
		csv.AddToCSVFile(resultfile, []string{" "})
		csv.AddToCSVFile(resultfile, []string{"Descending"})
		csv.AddToCSVFile(resultfile, []string{"Size", "Comparisons"})
		for x := 1; x <= 12; x++ {
			qcomparisons = 0
			N := int(math.Pow(2, float64(x)))
			array := ga.DescendingIntArray(N)

			Quicksort(array, 0, len(array)-1, pivot)
			csv.AddToCSVFile(resultfile, []string{fmt.Sprint(len(array)), fmt.Sprint(qcomparisons)})
		}
		csv.AddToCSVFile(resultfile, []string{" "})
		csv.AddToCSVFile(resultfile, []string{" "})
	}
}
