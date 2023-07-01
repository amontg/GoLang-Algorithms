package inclass

import "fmt"

func findMe(slice []int, target int) int {
	var loc int
	count := 0
	for i := 0; i < len(slice); i++ {
		count++
		if slice[i] == target {
			loc = i
			break
		}
	}

	fmt.Println("Comparisons: ", count, "| Average: ", float64(count)/float64(len(slice)))
	return loc
}

func inclass3() {
	var testy [100]int

	for i := 0; i < len(testy); i++ {
		testy[i] = i
	}

	key := findMe(testy[:], 50)
	fmt.Println("Key: ", key, "| Value:", testy[key])
}
