package main

import "fmt"

var primeNumbers []uint32

func registerPrime(prime uint32) {
	primeNumbers = append(primeNumbers, prime)
}

func findNextPrime(current uint32) uint32 {
	possiblyPrime := current + 1
mainloop:
	for {
		for _, prime := range primeNumbers {
			if possiblyPrime%prime == 0 {
				possiblyPrime++
				continue mainloop
			}
			if prime*prime > possiblyPrime {
				registerPrime(possiblyPrime)
				return possiblyPrime
			}
		}
	}
}

func checkIsPrime(number uint32) bool {
	if number == 1 {
		return false
	}
	for _, prime := range primeNumbers {
		if number == prime {
			return true
		}
		if number%prime == 0 {
			return false
		}
		if prime*prime > number {
			return true
		}
	}
	prime := primeNumbers[len(primeNumbers)-1]
	for {
		prime = findNextPrime(prime)
		if number%prime == 0 {
			return false
		}
		if prime*prime > number {
			return true
		}
	}
}

func main() {
	var caseCount, i int8
	var testNumbers [20]uint32
	registerPrime(2)
	registerPrime(3)
	fmt.Scanf("%d\n", &caseCount)
	for i = 0; i < caseCount; i++ {
		fmt.Scanf("%d\n", &testNumbers[i])
	}
	for i = 0; i < caseCount; i++ {
		if checkIsPrime(testNumbers[i]) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}
