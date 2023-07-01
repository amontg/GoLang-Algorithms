/*
	Author: Amir Montgomery
	Date: 1/24
	Class: CS 310, Algorithms, Reinhart
	Objective: Return an array (slice in Golang) of n number of prime numbers in sequential order
*/

package godelnumbering

import (
	"fmt"
	"math"
)

/*
	Pseudo Code
	1. Get number of primes to find (2 is always first)
	2. To determine if x is prime, divide by every number <= root x
*/

// Golang: int64 is the same as long, []int64 is a SLICE (>array), not an array. no while loops.
// Slices: starts at index 0, last index is length - 1
func primeSlice(n int) []int64 {
	var primeSeq []int64

	// append 2, append() returns []T
	primeSeq = append(primeSeq, 2)
	n--

	i := 3.0
	for n > 0 {
		if checkPrime(i) {
			primeSeq = append(primeSeq, int64(i))
			i++
			n--
		} else {
			i++
		}
	}

	fmt.Println("Prime Numbers: ", primeSeq)
	return primeSeq
}

func primeFactors(n int64) []int64 {
	var factorSeq []int64

	x := 1 // ignore
	for i := 2.0; n > 0.0; {
		if checkPrime(i) {
			if n%int64(i) == 0 { // if i is prime and n is evenly divisible by prime i, add it as an integer
				factorSeq = append(factorSeq, int64(i))
				n = n / int64(i)
				x++ // ignore
			} else {
				i++
			}
		} else if int64(i) > n {
			break
		} else {
			i++
			//fmt.Println(i)
		}
	}

	fmt.Println("Prime Factors: ", factorSeq)
	return factorSeq
}

func checkPrime(n float64) bool {
	// check if a number is prime. 2 is a given
	check := false

	if n == 1 {
		check = false
	} else if n == 2 {
		check = true
	} else {
		nSquare := math.Sqrt(n)

		i2 := 2.0
		if nSquare < i2 {
			if n/2.0 != 0.0 {
				check = true
			}
		} else {
			for i2 <= nSquare {
				if i2 == math.Floor(nSquare) {
					if math.Mod(n, i2) != 0.0 {
						check = true
						break
					} else {
						check = false
						break
					}
				} else {
					if math.Mod(n, i2) <= 0 {
						check = false
						break
					} else {
						i2++
					}
				}
			}
		}
	}

	return check
}

// func main() {
// 	fmt.Println(primeSlice(25))
// }
