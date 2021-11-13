package main

import (
	"fmt"
	"strconv"
	"time"
)

var isPrimeArray [1000000001]int8
var primeNumbers []int64
var highestCalculated int64 = 1

func checkIfPrime(number int64) bool {
	if isPrimeArray[number] == -1 {
		if number == 1 {
			return true
		} else if number <= 3 {
			isPrimeArray[number] = 1
			primeNumbers = append(primeNumbers, number)
		} else if number%2 == 0 || number%3 == 0 {
			isPrimeArray[number] = 0
		} else {
			for _, prime := range primeNumbers {
				if prime*prime > number {
					isPrimeArray[number] = 1
					primeNumbers = append(primeNumbers, number)
					break
				}
				if number%prime == 0 {
					isPrimeArray[number] = 0
					break
				}
			}
		}
	}
	if highestCalculated < number {
		highestCalculated = number
	}
	return isPrimeArray[number] == 1
}

func main() {
	for i := 1; i <= 1000000000; i++ {
		isPrimeArray[i] = -1
	}
	var minStr, maxStr, caseCountStr string
	fmt.Scan(&caseCountStr)
	caseCount, _ := strconv.Atoi(caseCountStr)
	var j, startFrom int64
	for i := 0; i < caseCount; i++ {
		fmt.Scan(&minStr)
		fmt.Scan(&maxStr)
		startTime := time.Now().UnixMilli()
		min, _ := strconv.ParseInt(minStr, 10, 64)
		max, _ := strconv.ParseInt(maxStr, 10, 64)
		if min >= highestCalculated {
			startFrom = highestCalculated
		} else {
			startFrom = min
			for _, p := range primeNumbers {
				if p < min {
					continue
				}
				if p > max {
					startFrom = max + 1
					break
				}
				fmt.Println(p)
				startFrom = p + 1
			}
		}
		for j = startFrom; j <= max; j++ {
			isPrime := checkIfPrime(j)
			if j >= min && isPrime {
				fmt.Println(j)
			}
		}
		endTime := time.Now().UnixMilli()
		fmt.Printf("Start: %d, End: %d, Elapsed time: %d\n", startTime, endTime, (endTime - startTime))
	}
}
