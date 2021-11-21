package main

import (
	"bufio"
	"fmt"
	"os"
	// "time"
)

const AREA = 1000001
const IS_PRIME = 1
const NOT_PRIME = 2
const UNKNOWN = 0
const DROPPED = 0

var numberStatuses [AREA]uint8
var primeNumbers []uint32	

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func produceNextPrime() uint32 {
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

func suggestNextPossiblyPrime(current uint32) uint32 {
	possiblyPrime := current + 1
	for {
		if isUnknownStatus(possiblyPrime) {
			return possiblyPrime
		}
		possiblyPrime++
	}
}

func registerPrime(prime uint32) {
	primeNumbers = append(primeNumbers, prime)
	setIsPrime(prime)
	unsetCompositesOfPrimeInNumberStatuses(prime)
}

func unsetCompositesOfPrimeInNumberStatuses(prime uint32) {
	var len = uint32(len(numberStatuses))
	var i, k uint32
	k = prime
	for i = 1; k < len; i++ {
		if isUnknownStatus(k) {
			setNotPrime(k)
		}
		k = prime * i
	}
}

func setIsPrime(number uint32) {
	numberStatuses[number-1] = IS_PRIME
}

func setNotPrime(number uint32) {
	numberStatuses[number-1] = NOT_PRIME
}

func isUnknownStatus(number uint32) bool {
	return numberStatuses[number-1] == UNKNOWN
}

func findNextPrime(currentPrime uint32) uint32 {
	var prevPrime uint32
	for _, prime := range primeNumbers {
		if prevPrime == currentPrime {
			return prime
		}
		prevPrime = prime
	}
	return produceNextPrime()
}

func runCase(min uint32, max uint32, caseIdx int) {
	var area [AREA]uint32
	size := max - min + 1
	var i, st uint32
	for i = 0; i < size; i++ {
		area[i] = min + i
	}
	if (area[0] == 1) {
		area[0] = DROPPED
	}
	var prime uint32 = 2
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
		// printf("prime: %d max: %d\n", prime, max)
		if prime*prime > max {
			// printf("break on prime: %d max: %d\n", prime, max)
			break
		}
	}

	for i := 0; uint32(i) < size; i++ {
		v := area[i]
		if v != DROPPED {
			printf("%d\n", v)
		}
	}
	writer.Flush()
}

func main() {
	setNotPrime(1)
	registerPrime(2)
	registerPrime(3)
	var caseCount int
	var min, max uint32
	var mins, maxs [150]uint32

	scanf("%d\n", &caseCount)
	for i := 0; i < caseCount; i++ {
		scanf("%d %d\n", &min, &max)
		mins[i] = min
		maxs[i] = max
	}

	for i := 0; i < caseCount; i++ {
		// var startTime, endTime []uint32
		// startTime = append(startTime, time.Now().UnixMilli())
		runCase(mins[i], maxs[i], i)
		// endTime = append(endTime, time.Now().UnixMilli())
		// fmt.Printf("Iteration: %d, Elapsed time (sec): %d\n", i, (endTime[i] - startTime[i]))
	}
}

// 2146483647 2147483647
// 9999000000 10000000000
// 999900000 1000000000
// 999800000 999900000
// 999700000 999800000
