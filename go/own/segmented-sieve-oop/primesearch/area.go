package primesearch

import (
	"fmt"
	"bufio"
	. "segmented-sieve-oop/types"
)

type Area struct {
	Size     uint32
	Statuses []uint32
	MinMax   Range
}

const DROPPED = 0

func (a *Area) Construct(minMax Range, areaSize uint32) {
	a.MinMax = minMax
	a.Statuses = make([]uint32, areaSize)
	var i uint32
	size := a.calcSize()
	for i = 0; i < size; i++ {
		a.Statuses[i] = a.MinMax.Min + i
	}
	if a.Statuses[0] == 1 {
		a.Statuses[0] = DROPPED
	}
}

func (a *Area) calcSize() uint32 {
	return a.MinMax.Max - a.MinMax.Min + 1
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