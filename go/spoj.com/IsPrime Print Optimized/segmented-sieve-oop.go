package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

const IS_PROFILING = true
const IS_RESULT_OUTPUT = false
const AREA = 1000001

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

// --- Area ---

type Area struct {
	Size     uint32
	Statuses [AREA]uint32
	MinMax   Range
}

const DROPPED = 0

func (a *Area) construct(minMax Range) {
	a.MinMax = minMax
	a.Statuses = [AREA]uint32{}
	var i uint32
	size := a.calcSize()
	for i = 0; i < size; i++ {
		a.Statuses[i] = a.MinMax.min + i
	}
	if a.Statuses[0] == 1 {
		a.Statuses[0] = DROPPED
	}
}

func (a *Area) calcSize() uint32 {
	return a.MinMax.max - a.MinMax.min + 1
}

func (a *Area) isPrimeOverPossible(prime uint32) bool {
	return prime*prime > a.MinMax.max
}

func (a *Area) sieve(prime uint32) {
	// a.printMyselfInLine(fmt.Sprintf("Area before sieve for prime %d", prime))
	var st uint32
	if a.MinMax.min%prime == 0 {
		st = 0
	} else {
		st = prime - (a.MinMax.min % prime)
	}
	size := a.calcSize()
	for i := st; i < size; i += prime {
		if a.Statuses[i] != DROPPED && a.Statuses[i] != prime {
			a.Statuses[i] = DROPPED
		}
	}
	// a.printMyselfInLine(fmt.Sprintf("Area after sieve for prime %d", prime))
}

func (a *Area) printMyselfInLine(prefix string) {
	printf("%s: min: %d, max: %d, %v\n", prefix, a.MinMax.min, a.MinMax.max, a.Statuses)
}

func (a *Area) printMyself() {
	size := a.calcSize()
	for i := 0; uint32(i) < size; i++ {
		v := a.Statuses[i]
		if v != DROPPED {
			printf("%d\n", v)
		}
	}
	writer.Flush()
}

// --- PrimeFinder ---

type PrimeFinder struct {
	Value          uint32
	numberStatuses [AREA]uint8
	primeNumbers   []uint32
	isResultOutput bool
}

const IS_PRIME = 1
const NOT_PRIME = 2
const UNKNOWN = 0

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
	// printf(fmt.Sprintf("after registerPrime: %d\n", prime))
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

func (p *PrimeFinder) printMyself(prefix string) {
	printf("%s Value: %d, primeNumbers: %v numberStatuses: %v\n", prefix, p.Value, p.primeNumbers, p.numberStatuses)
	writer.Flush()
}

func (p *PrimeFinder) runCase(minMax Range, caseIdx int) {
	var area *Area = &Area{}
	area.construct(minMax)
	p.startWith(2)
	for {
		area.sieve(p.Value)
		p.findNextPrime()
		if area.isPrimeOverPossible(p.Value) {
			break
		}
	}

	if p.isResultOutput {
		area.printMyself()
	}
}

// --- InputController ---

type InputController struct {
	caseCount int
	ranges    [150]Range
}

func (ic *InputController) read() {
	var min, max uint32
	scanf("%d\n", &ic.caseCount)
	for i := 0; i < ic.caseCount; i++ {
		scanf("%d %d\n", &min, &max)
		ic.ranges[i] = Range{min: min, max: max}
	}
}

func (ic *InputController) rangeByIndex(idx int) Range {
	return ic.ranges[idx]
}

// --- MinMax OV ---

type Range struct {
	min uint32
	max uint32
}

// --- Profiler ---

type Profiler struct {
	isEnabled bool
	startTime []int64
	endTime   []int64
}

func (pr *Profiler) start() {
	if !pr.isEnabled {
		return
	}
	pr.startTime = append(pr.startTime, time.Now().UnixMilli())
}

func (pr *Profiler) end() {
	if !pr.isEnabled {
		return
	}
	pr.endTime = append(pr.endTime, time.Now().UnixMilli())
}

func (pr *Profiler) elapsedTime(idx int) int64 {
	if !pr.isEnabled {
		return 0
	}
	return pr.endTime[idx] - pr.startTime[idx]
}

func (pr *Profiler) printMyself() {
	if !pr.isEnabled {
		return
	}
	for idx := range pr.startTime {
		printf(fmt.Sprintf("Iteration: %d, Elapsed time (sec): %d\n", idx, pr.elapsedTime(idx)))
	}
	writer.Flush()
}

// --- App ---

type App struct{}

func (app *App) run() {
	var primeFinder *PrimeFinder = &PrimeFinder{isResultOutput: IS_RESULT_OUTPUT}
	primeFinder.setNotPrime(1)
	primeFinder.registerPrime(2)
	primeFinder.registerPrime(3)

	var inputController *InputController = &InputController{}
	inputController.read()

	var profiler *Profiler = &Profiler{isEnabled: IS_PROFILING}

	for i := 0; i < inputController.caseCount; i++ {
		profiler.start()
		primeFinder.runCase(inputController.rangeByIndex(i), i)
		profiler.end()
	}
	profiler.printMyself()
}

// --- main ---

func main() {
	var app *App = &App{}
	app.run()
}

// 2146483647 2147483647
// 9999000000 10000000000
// 999900000 1000000000
// 999800000 999900000
// 999700000 999800000
