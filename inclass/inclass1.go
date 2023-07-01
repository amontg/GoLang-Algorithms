package inclass

import (
	"fmt"
	"math/rand"
)

func median(x int, y int, z int) int {
	var rMedian int

	/*
		XY - YZ > 0 = (x>y) > (y>z) = x > y > z
		XZ - -YZ > 0 = (x>z) > (z>y) = x > z > y
		YZ - -XZ > 0 = (y>z) > (z>x) = y > z > x
		XZ - -XY < 0 = (y>x) > (x>z) = y > x > z | (z>x) < (y>x) - (y>x) > (x>z)
		-XZ - XY > 0 = (z>x) > (x>y) = z > x > y
		-YZ - -XY > 0 = (z>y) > (y>x) = z > y > x
	*/

	if ((x - y) - (y - z)) > 0 { // x > y > z
		rMedian = y
	} else if ((x - z) - (z - y)) > 0 { // x > z > y
		rMedian = z
	} else if ((y - z) - (z - x)) > 0 { // y > z > x
		rMedian = z
	} else if ((y - x) - (x - z)) > 0 { // y > x > z
		rMedian = x
	} else if ((z - x) - (x - y)) > 0 { // z > x > y
		rMedian = x
	} else if ((z - y) - (y - x)) > 0 { // z > y > x
		rMedian = y
	}

	return rMedian
}

func inclass1() {
	var x int
	var y int
	var z int
	for i := 0; i < 5; i++ {
		x = rand.Int()
		y = rand.Int()
		z = rand.Int()
		fmt.Println(x, y, z)
		fmt.Println(median(x, y, z))
	}
}
