package probabilities

import (
	"fmt"
	"math/rand"
	"time"
)

func rollMe(die []int) int {
	rand.Seed(time.Now().UnixNano())
	return die[rand.Intn(len(die))]
}

func fourDice() {
	d1 := []int{1, 2, 3, 9, 10, 11}
	d2 := []int{0, 1, 7, 8, 8, 9}
	d3 := []int{5, 5, 6, 6, 7, 7}
	d4 := []int{3, 4, 4, 5, 11, 12}
	var d1vx [4]float64 // {vd1, vd2, vd3, vd4} = increment if d1 >
	var d2vx [4]float64
	var d3vx [4]float64
	var d4vx [4]float64
	allDice := [][]int{d1, d2, d3, d4}
	compDice := [][4]float64{d1vx, d2vx, d3vx, d4vx}
	n := 60000.0

	for i := 0.0; i < n; i++ {
		for j := 0; j < len(allDice); j++ {
			for x := 0; x < len(allDice); x++ {
				if x != j {
					rollOne := rollMe(allDice[j])
					rollTwo := rollMe(allDice[x])
					if rollOne > rollTwo {
						compDice[j][x]++
					}
				}
			}
		}
	}

	for i := 0; i < len(compDice); i++ {
		/*
			for j := 0; j < len(compDice[i]); j++ {
				fmt.Println(compDice[i][j] / rolls)
			}
			fmt.Println("---")
		*/
		fmt.Println(compDice[i])
	}

}

func matrix4d() {
	fourDice()
}
