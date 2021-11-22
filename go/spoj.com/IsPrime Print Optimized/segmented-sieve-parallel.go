package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

const DEBUG = false
const CASE_MAX = 150
const INIT_PRIME_MAX = 46340       // sqrt(2147483647). Real possible highest prime is 46337 (the next is 46349)
const INIT_PRIME_AREA_SIZE = 15446 // INIT_PRIME_HIGHEST / 3

const AREA = 1000001
const DROPPED = 0

const PROCESSING = 0
const READY = 1
const PRINTED = 2

var completedCases []byte
var canPrintIdx byte
var caseCount byte

var primeNumbers = make([]uint32, INIT_PRIME_AREA_SIZE)
var primesCount uint32

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }
func ll(f string, a ...interface{}) {
	if DEBUG {
		fmt.Printf(f+"\n", a...)
	}
}

func canPrint(searchCaseIdx byte) bool {
	var b byte
	ll("canPrint %d?, completedCases: %v", searchCaseIdx, completedCases)
	for b = 0; b < searchCaseIdx; b++ {
		if completedCases[b] != PRINTED {
			return false
		}
	}
	return completedCases[searchCaseIdx] == READY
}

func runCase(min uint32, max uint32, caseIdx byte, wg *sync.WaitGroup) {
	// defer wg.Done() // Так вызовится по выходу из ф-ции
	completedCases[caseIdx] = PROCESSING
	ll("Run case %d, completedCases: %v", caseIdx, completedCases)
	var area [AREA]uint32
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
	completedCases[caseIdx] = READY
	ll("Collected case %d, completedCases: %v waiting...", caseIdx, completedCases)

	for !canPrint(caseIdx) {
		time.Sleep(time.Nanosecond)
	}

	for i = 0; i < size; i++ {
		v := area[i]
		if v != DROPPED {
			printf("%d\n", v)
		}
	}
	writer.Flush()

	completedCases[caseIdx] = PRINTED
	ll("End case %d, completedCases: %v", caseIdx, completedCases)

	wg.Done()
}

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

func main() {
	buildInitialPrimes()
	var b byte
	var min, max uint32
	var mins, maxs [CASE_MAX]uint32

	scanf("%d\n", &caseCount)
	for b = 0; b < caseCount; b++ {
		scanf("%d %d\n", &min, &max)
		mins[b] = min
		maxs[b] = max
	}

	completedCases = make([]byte, caseCount)
	wg := new(sync.WaitGroup) // allocate memory for type and return pointer to it
	wg.Add(int(caseCount))
	for b = 0; b < caseCount; b++ {
		go runCase(mins[b], maxs[b], b, wg)
	}
	wg.Wait()
}

// 2146483647 2147483647
