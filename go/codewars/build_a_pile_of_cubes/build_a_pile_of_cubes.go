package main

import (
	"fmt"
	"math"
)

/**
 * Your task is to construct a building which will be a pile of n cubes. The cube at the bottom will have a volume of n^3, the cube above will have volume of (n-1)^3 and so on until the top which will have a volume of 1^3.
 * You are given the total volume m of the building. Being given m can you find the number n of cubes you will have to build?
 * The parameter of the function findNb (find_nb, find-nb, findNb, ...) will be an integer m and you have to return the integer n such as n^3 + (n-1)^3 + ... + 1^3 = m if such a n exists or -1 if there is no such
 */

func FindNb(m int) int {
	var volumeN, acc int
	for i := 1; true; i++ {
		volumeN = int(math.Pow(float64(i), 3))
		acc += volumeN
		if acc == m {
			return i
		} else if acc > m {
			return -1
		}
	}
	return -1
}

func main() {
	ll(fmt.Sprintf("Result: %d", FindNb(4183059834010)))
}

func ll(str string) {
	println(str)
}

/**
 * Другое решение: sum of cubes from 1 to n equals n^2(n + 1)^2 / 4
 */
