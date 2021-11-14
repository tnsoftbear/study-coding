package main

import (
	"bufio"
	"fmt"
	"os"
	// "time"
)

const AREA = 1000001
const IS_PRIME = 1
const NOT_PRIME = -1
const UNKNOWN = 0
const DROPPED = 0

var numberStatuses [AREA]int8
var primeNumbers []int64	
var allResults [][]int64

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func produceNextPrime() int64 {
	highestPrime := primeNumbers[len(primeNumbers)-1]
	possiblyPrime := suggestNextPossiblyPrime(highestPrime)
mainloop:
	for {
		for _, prime := range primeNumbers {
			if possiblyPrime%prime == 0 {
				setNotPrime(possiblyPrime)
				possiblyPrime = suggestNextPossiblyPrime(possiblyPrime)
				continue mainloop
			}
			if prime*prime > possiblyPrime {
				registerPrime(possiblyPrime)
				return possiblyPrime
			}
		}
	}
}

func suggestNextPossiblyPrime(current int64) int64 {
	possiblyPrime := current + 1
	for {
		if isUnknownStatus(possiblyPrime) {
			return possiblyPrime
		}
		possiblyPrime++
	}
}

func registerPrime(prime int64) {
	primeNumbers = append(primeNumbers, prime)
	setIsPrime(prime)
	unsetCompositesOfPrimeInNumberStatuses(prime)
}

func unsetCompositesOfPrimeInNumberStatuses(prime int64) {
	var len = int64(len(numberStatuses))
	var i, k int64
	k = prime
	for i = 1; k < len; i++ {
		if isUnknownStatus(k) {
			setNotPrime(k)
		}
		k = prime * i
	}
}

func setIsPrime(number int64) {
	numberStatuses[number-1] = IS_PRIME
}

func setNotPrime(number int64) {
	numberStatuses[number-1] = NOT_PRIME
}

func isUnknownStatus(number int64) bool {
	return numberStatuses[number-1] == UNKNOWN
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
	var i, st int64
	for i = 0; i < size; i++ {
		area[i] = min + i
	}
	if (area[0] == 1) {
		area[0] = DROPPED
	}
	var prime int64 = 2
	for {
		if min%prime == 0 {
			st = 0
		} else {
			st = prime - (min % prime)
		}
		for i := st; i < size; i += prime {
			if area[i] != DROPPED && area[i] != prime {
				area[i] = DROPPED
			}
		}

 		prime = findNextPrime(prime)
		if prime*prime > max {
			break
		}
	}

	allResults = append(allResults, []int64{})
	for _, v := range area {
		if v != DROPPED {
			allResults[caseI] = append(allResults[caseI], v)
		}
	}
}

func main() {
	setNotPrime(1)
	registerPrime(2)
	registerPrime(3)
	// var startTime, endTime []int64
	var caseCount int
	var min, max int64

	defer writer.Flush()
	scanf("%d\n", &caseCount)
	for i := 0; i < caseCount; i++ {
		scanf("%d %d\n", &min, &max)
		// startTime = append(startTime, time.Now().UnixMilli())
		runCase(min, max, i)
		// endTime = append(endTime, time.Now().UnixMilli())
	}

	for i := 0; i < caseCount; i++ {
		for _, v := range allResults[i] {
			printf("%d\n", v)
		}
	}
	// for i := 0; i < caseCount; i++ {
	// 	fmt.Printf("Iteration: %d, Elapsed time (sec): %d\n", i, (endTime[i] - startTime[i]))
	// }
}

// 2146483647 2147483647
// 9999000000 10000000000
// 999900000 1000000000
// 999800000 999900000
// 999700000 999800000
