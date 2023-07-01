package probabilities

import (
	"fmt"
	"math/rand"
	"time"
)

func funkyRoll() int {
	rand.Seed(time.Now().UnixNano())
	funkyDie := []int{1, 2, 3, 3, 4, 5, 5, 5}
	return funkyDie[rand.Intn(len(funkyDie))]
}

func funkyTwoDice() {
	var rolls [11]int
	dice := []int{0, 0}
	n := 60000.0

	for i := 0.0; i < n; i++ {
		for j := 0; j < len(dice); j++ {
			dice[j] = funkyRoll()
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

func weirdprob2d5() {
	funkyTwoDice()
}
