package generatearray

import (
	"math/rand"
	"time"
)

func AscendingIntArray(size int) []int {
	var ascendingArray []int
	for a := 0; a <= size; a++ {
		ascendingArray = append(ascendingArray, a)
	}

	return ascendingArray
}

func DescendingIntArray(size int) []int {
	var descendingArray []int
	for d := 0; d <= size; d++ {
		descendingArray = append(descendingArray, size-d)
	}

	return descendingArray
}

func RandomizedIntArray(size int) []int {
	randomizedArray := AscendingIntArray(size)
	for i := size - 1; i >= 1; i-- {
		rand.Seed(time.Now().Unix())

		tempInt := randomizedArray[i]
		randSlot := rand.Intn(i + 1)
		randomizedArray[i] = randomizedArray[randSlot]
		randomizedArray[randSlot] = tempInt
	}

	return randomizedArray
}

func Swap(list []int, key int, bf int) []int {
	updatedList := list
	switch bf {
	case -1: // swap backwards
		tempInt := updatedList[key]
		updatedList[key] = updatedList[key-1]
		updatedList[key-1] = tempInt
		return updatedList
	case 1: // swap forward
		tempInt := updatedList[key]
		updatedList[key] = updatedList[key+1]
		updatedList[key+1] = tempInt
		return updatedList
	}
	return updatedList
}

func SwapPlaces(list []int, indexone, indextwo int) []int {
	updatedList := list
	temp := updatedList[indexone]
	updatedList[indexone] = updatedList[indextwo]
	updatedList[indextwo] = temp
	return updatedList
}
