package probabilities

import (
	"fmt"
	"math/rand"
	"time"
)

func roll(sides int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(sides) + 1
}

func TwoDice() {
	var rolls [11]int
	dice := []int{0, 0}
	n := 60000.0

	for i := 0.0; i < n; i++ {
		for j := 0; j < len(dice); j++ {
			dice[j] = roll(5)
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

func prob2d5() {
	TwoDice()
}
