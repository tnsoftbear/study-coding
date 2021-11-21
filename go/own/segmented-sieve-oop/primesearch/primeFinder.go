package primesearch

import (
	"fmt"
	"bufio"
	"segmented-sieve-oop/types"
)

const IS_PRIME = 1
const NOT_PRIME = 2
const UNKNOWN = 0

type PrimeFinder struct {
	Value          uint32
	numberStatuses []uint8
	primeNumbers   []uint32
	isResultOutput bool
	areaSize uint32
}

func (p *PrimeFinder) Construct(areaSize uint32, isResultOutput bool) {
	p.numberStatuses = make([]uint8, areaSize)
	p.areaSize = areaSize
	p.isResultOutput = isResultOutput
	p.setNotPrime(1)
	p.registerPrime(2)
	p.registerPrime(3)
}

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
				p.Value = possiblyPrime
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
	var len = uint32(len(p.numberStatuses))
	var i, k uint32
	k = prime
	for i = 1; k < len; i++ {
		if p.isUnknownStatus(k) {
			p.setNotPrime(k)
		}
		k = prime * i
	}
}

func (p *PrimeFinder) setIsPrime(number uint32) {
	p.numberStatuses[number-1] = IS_PRIME
}

func (p *PrimeFinder) setNotPrime(number uint32) {
	p.numberStatuses[number-1] = NOT_PRIME
}

func (p *PrimeFinder) isUnknownStatus(number uint32) bool {
	return p.numberStatuses[number-1] == UNKNOWN
}

func (p *PrimeFinder) findNextPrime() {
	var prevPrime uint32
	currentPrime := p.Value
	for _, prime := range p.primeNumbers {
		if prevPrime == currentPrime {
			p.Value = prime
			return
		}
		prevPrime = prime
	}
	p.produceNextPrime()
}

func (p *PrimeFinder) startWith(init uint32) {
	p.Value = init
}

func (p *PrimeFinder) PrintMyself(prefix string, writer *bufio.Writer) {
	fmt.Fprintf(writer, "%s Value: %d, primeNumbers: %v numberStatuses: %v\n", prefix, p.Value, p.primeNumbers, p.numberStatuses)
	writer.Flush()
}

func (p *PrimeFinder) RunCase(minMax types.Range, caseIdx int, writer *bufio.Writer) {
	var area *Area = &Area{}
	area.Construct(minMax, p.areaSize)
	p.startWith(2)
	for {
		area.Sieve(p.Value)
		p.findNextPrime()
		if area.IsPrimeOverPossible(p.Value) {
			break
		}
	}

	if p.isResultOutput {
		area.PrintMyself(writer)
	}
}
