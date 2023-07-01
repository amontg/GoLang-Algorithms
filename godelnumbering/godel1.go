/*
	Author: Amir Montgomery
	Date: 1/24
	Class: CS310, Algorithms, Reinhart
	Objective: Calculate the godel number from a list of numbers
		Godel Number = Input: x1, x2, x3, ... xe -> (2^x1 * 3^x2 * 5^x3 ... p^xe) where p = prime numbers, x = exponent

*/

package godelnumbering

import (
	"fmt"
	"math"
)

/*
	Pseudo Code
	1. Get array of exponent numbers (ENs)
	2. Find length of array
	3. Use array len to find sequence of prime numbers with primenumbers.go (output array of prime numbers)
	4. Use for loop to multiply primenumbers^godelnumbers
		4a. Is for loop most effective?


*/

func getgodel(exponents []int64) {
	primes := primeSlice(len(exponents))
	var sol int64 = 1

	i := 0
	for i < len(exponents) {
		sol = sol * int64(math.Pow(float64(primes[i]), float64(exponents[i])))
		//fmt.Println(sol)
		i++
	}

	primeFactors(sol)
	fmt.Println("Solution: ", sol)
}

func dogodel() {
	exponents := []int64{7, 6, 5, 4, 3, 2, 1}
	fmt.Println("Exponents: ", exponents)
	getgodel(exponents)
}
