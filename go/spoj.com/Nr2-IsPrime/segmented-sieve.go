package main

import (
	"fmt"
	"strconv"
)

const AREA = 100001

var isPrimeArray [AREA]int8
var primeNumbers = []int64{2, 3}
var allResults []string

func produceNextPrime() int64 {
	highestPrime := primeNumbers[len(primeNumbers)-1]
	possiblyPrime := optimizedNextPossiblyPrime(highestPrime)
mainloop:
	for {
		for _, prime := range primeNumbers {
			if possiblyPrime%prime == 0 {
				isPrimeArray[possiblyPrime-1] = 0
				possiblyPrime = optimizedNextPossiblyPrime(possiblyPrime)
				continue mainloop
			}
			if prime*prime > possiblyPrime {
				primeNumbers = append(primeNumbers, possiblyPrime)
				isPrimeArray[possiblyPrime-1] = 1
				return possiblyPrime
			}
		}
	}
}

func optimizedNextPossiblyPrime(current int64) int64 {
	possiblyPrime := current + 1
	for {
		if isPrimeArray[possiblyPrime-1] == -1 {
			if possiblyPrime%2 == 0 || possiblyPrime%3 == 0 {
				isPrimeArray[possiblyPrime-1] = 0
				possiblyPrime++
			} else {
				return possiblyPrime
			}
		} else {
			possiblyPrime++
		}
	}
}

func findNextPrime(currentPrime int64) int64 {
	var prevPrime int64
	for _, prime := range primeNumbers {
		if prevPrime == currentPrime {
			return prime
		}
		prevPrime = prime
	}
	return produceNextPrime()
}

func runCase(min int64, max int64, caseI int) {
	size := max - min + 1
	var area [AREA]int64
	var i int64
	for i = 0; i < size; i++ {
		area[i] = min + i
	}
	if area[0] == 1 {
		area[0] = 0
	}
	var prime int64 = 2
	for {
		for i, v := range area {
			if v == 0 {
				continue
			}
			if v != prime && v%prime == 0 {
				area[i] = 0
			}
		}
		prime = findNextPrime(prime)
		if prime*prime > max {
			break
		}
	}
	var concatenatedResult string
	for _, v := range area {
		if v > 0 {
			concatenatedResult += fmt.Sprintf("%d\n", v)
		}
	}
	allResults = append(allResults, concatenatedResult)
}

func main() {
	for i := 0; i < AREA; i++ {
		isPrimeArray[i] = -1
	}
	isPrimeArray[1] = 0
	isPrimeArray[2] = 1
	isPrimeArray[3] = 1
	var minStr, maxStr, caseCountStr string
	fmt.Scan(&caseCountStr)
	caseCount, _ := strconv.Atoi(caseCountStr)
	for i := 0; i < caseCount; i++ {
		fmt.Scan(&minStr)
		fmt.Scan(&maxStr)
		min, _ := strconv.ParseInt(minStr, 10, 64)
		max, _ := strconv.ParseInt(maxStr, 10, 64)

		runCase(min, max, i)
	}
	for i := 0; i < caseCount; i++ {
		fmt.Println(allResults[i])
	}
}

// 999900000 1000000000
