package inclass

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func roll(sides int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(sides) + 1
}

func TwoDice() {
	var rolls [13]int
	dice := []int{0, 0}
	n := 6000.0

	for i := 0.0; i < n; i++ {
		for j := 0; j < len(dice); j++ {
			dice[j] = roll(6)
		}
		rolls[dice[0]+dice[1]]++
	}

	sum := 0.0
	for i := 1; i < len(rolls); i++ {
		fmt.Println(i, float64(rolls[i])/n)
		sum += float64(rolls[i]) / n
	}
	fmt.Println(sum)
}

func inclass2() {
	TwoDice()

	var a float64 = 5
	var sum float64 = 0
	n := 10.0
	for i := 0.0; i <= n; i++ {
		sum += math.Pow(a, i)
	}

	fmt.Println(sum, "=", ((math.Pow(a, n+1) - 1) / (a - 1)))
}
