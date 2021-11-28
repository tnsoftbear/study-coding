package primesearch

import (
	"fmt"
	"bufio"
	"segmented-sieve-oop/types"
)

type Area struct {
	Statuses []uint32
	MinMax   types.Range
}

const DROPPED = 0

func NewArea(minMax types.Range, areaSize uint32) *Area {
	a := Area{
		Statuses: make([]uint32, areaSize),
		MinMax: minMax,
	}
	var i uint32
	size := a.calcSize()
	for i = 0; i < size; i++ {
		a.Statuses[i] = a.MinMax.Min + i
	}
	if a.Statuses[0] == 1 {
		a.Statuses[0] = DROPPED
	}
	return &a
}

func (a *Area) IsPrimeOverPossible(prime uint32) bool {
	return prime*prime > a.MinMax.Max
}

func (a *Area) Sieve(prime uint32) {
	// a.printMyselfInLine(fmt.Sprintf("Area before sieve for prime %d", prime))
	var st uint32
	if a.MinMax.Min%prime == 0 {
		st = 0
	} else {
		st = prime - (a.MinMax.Min % prime)
	}
	size := a.calcSize()
	for i := st; i < size; i += prime {
		if a.Statuses[i] != DROPPED && a.Statuses[i] != prime {
			a.Statuses[i] = DROPPED
		}
	}
	// a.printMyselfInLine(fmt.Sprintf("Area after sieve for prime %d", prime))
}

func (a *Area) PrintMyselfInLine(prefix string, writer *bufio.Writer) {
	fmt.Fprintf(writer, "%s: min: %d, max: %d, %v\n", prefix, a.MinMax.Min, a.MinMax.Max, a.Statuses)
}

func (a *Area) PrintMyself(writer *bufio.Writer) {
	size := a.calcSize()
	for i := 0; uint32(i) < size; i++ {
		v := a.Statuses[i]
		if v != DROPPED {
			fmt.Fprintf(writer, "%d\n", v)
		}
	}
	writer.Flush()
}

func (a *Area) calcSize() uint32 {
	return a.MinMax.Max - a.MinMax.Min + 1
}
