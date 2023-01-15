package primesearch

import (
	"bufio"
	"fmt"
	"segmented-sieve-project/internal/prime-wizard/domain/types"
)

const (
	unknown byte = iota
	isPrime
	notPrime
)

type PrimeFinder struct {
	value          uint32
	numberStatuses []uint8
	primeNumbers   []uint32
	isResultOutput bool
	areaSize       uint32
}

func NewPrimeFinder(areaSize uint32, isResultOutput bool) *PrimeFinder {
	p := PrimeFinder{
		numberStatuses: make([]uint8, areaSize),
		areaSize:       areaSize,
		isResultOutput: isResultOutput,
	}
	p.setNotPrime(1)
	p.registerPrime(2)
	p.registerPrime(3)
	return &p
}

func (p *PrimeFinder) RunCase(minMax types.Range, writer *bufio.Writer) {
	primes := p.DetectPrimes(minMax)

	if p.isResultOutput {
		p.printPrimes(primes, writer)
	}
}

func (p *PrimeFinder) DetectPrimes(minMax types.Range) []uint32 {
	area := NewArea(minMax, p.areaSize)
	p.startWith(2)
	for {
		area.Sieve(p.value)
		p.findNextPrime()
		if area.IsPrimeOverPossible(p.value) {
			break
		}
	}
	return area.GetFoundPrimes()
}

func (p *PrimeFinder) PrintMyself(prefix string, writer *bufio.Writer) {
	println(writer, "%s value: %d, primeNumbers: %v numberStatuses: %v\n", prefix, p.value, p.primeNumbers, p.numberStatuses)
	err := writer.Flush()
	if err != nil {
		println(writer, "Flush() call failed")
		return
	}
}

// --- private ---

func (p *PrimeFinder) produceNextPrime() {
	highestPrime := p.primeNumbers[len(p.primeNumbers)-1]
	possiblyPrime := p.suggestNextPossiblyPrime(highestPrime)
mainloop:
	for {
		for _, prime := range p.primeNumbers {
			if possiblyPrime%prime == 0 {
				p.setNotPrime(possiblyPrime)
				possiblyPrime = p.suggestNextPossiblyPrime(possiblyPrime)
				continue mainloop
			}
			if prime*prime > possiblyPrime {
				p.registerPrime(possiblyPrime)
				p.value = possiblyPrime
				return
			}
		}
	}
}

func (p *PrimeFinder) suggestNextPossiblyPrime(current uint32) uint32 {
	possiblyPrime := current + 1
	for {
		if p.isUnknownStatus(possiblyPrime) {
			return possiblyPrime
		}
		possiblyPrime++
	}
}

func (p *PrimeFinder) registerPrime(prime uint32) {
	p.primeNumbers = append(p.primeNumbers, prime)
	p.setIsPrime(prime)
	p.unsetCompositesOfPrimeInNumberStatuses(prime)
}

func (p *PrimeFinder) unsetCompositesOfPrimeInNumberStatuses(prime uint32) {
	var l = uint32(len(p.numberStatuses))
	var i, k uint32
	k = prime
	for i = 1; k < l; i++ {
		if p.isUnknownStatus(k) {
			p.setNotPrime(k)
		}
		k = prime * i
	}
}

func (p *PrimeFinder) setIsPrime(number uint32) {
	p.numberStatuses[number-1] = isPrime
}

func (p *PrimeFinder) setNotPrime(number uint32) {
	p.numberStatuses[number-1] = notPrime
}

func (p *PrimeFinder) isUnknownStatus(number uint32) bool {
	return p.numberStatuses[number-1] == unknown
}

func (p *PrimeFinder) findNextPrime() {
	var prevPrime uint32
	currentPrime := p.value
	for _, prime := range p.primeNumbers {
		if prevPrime == currentPrime {
			p.value = prime
			return
		}
		prevPrime = prime
	}
	p.produceNextPrime()
}

func (p *PrimeFinder) startWith(init uint32) {
	p.value = init
}

func (p *PrimeFinder) println(writer *bufio.Writer, f string, args ...interface{}) {
	_, err := fmt.Fprintf(writer, f+"\n", args...)
	if err != nil {
		return
	}
}

func (p *PrimeFinder) printPrimes(primes []uint32, writer *bufio.Writer) {
	for _, prime := range primes {
		fmt.Fprintln(writer, prime)
	}
}
