package main

// Time: 0.47
// Mem: 14M

import (
	"bufio"
	"fmt"
	"os"
)

const INIT_PRIME_MAX = 46340   // sqrt(2147483647). Real possible highest prime is 46337 (the next is 46349)
const INIT_PRIME_AREA_SIZE = 15446 // INIT_PRIME_HIGHEST / 3

const AREA = 1000001
const DROPPED = 0

var primeNumbers = make([]uint32, INIT_PRIME_AREA_SIZE)
var primesCount uint32

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func buildInitialPrimes() {
	var isCompositeArr = make([]uint32, INIT_PRIME_MAX+1)
	var checkingNumber uint32
	for checkingNumber = 2; checkingNumber <= INIT_PRIME_MAX; checkingNumber++ {
		if isCompositeArr[checkingNumber] == 0 {
			primeNumbers[primesCount] = checkingNumber
			primesCount++
			for j := checkingNumber + checkingNumber; j <= INIT_PRIME_MAX; j += checkingNumber {
				isCompositeArr[j] = 1
			}
		}
	}
}

var area [AREA]uint32

func runCase(min uint32, max uint32, caseIdx int) {
	size := max - min + 1
	var i, st uint32
	for i = 0; i < size; i++ {
		area[i] = min + i
	}
	if area[0] == 1 {
		area[0] = DROPPED
	}
	var prime uint32 = 2
	for i = 0; i < primesCount; i++ {
		prime = primeNumbers[i]
		st = min % prime
		if st != 0 {
			st = prime - st
		}
		for i := st; i < size; i += prime {
			if area[i] != DROPPED && area[i] != prime {
				area[i] = DROPPED
			}
		}

		if prime*prime > max {
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
	buildInitialPrimes()
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
		runCase(mins[i], maxs[i], i)
	}
}

// 2146483647 2147483647
