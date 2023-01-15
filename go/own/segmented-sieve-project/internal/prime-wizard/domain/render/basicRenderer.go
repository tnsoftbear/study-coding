package render

import (
	"bufio"
	"fmt"
)

type BasicRenderer struct {
	writer *bufio.Writer
}

func NewBasicRenderer(writer *bufio.Writer) *BasicRenderer {
	return &BasicRenderer{
		writer: writer,
	}
}

func (br *BasicRenderer) PrintPrimes(primes []uint32) {
	for _, prime := range primes {
		fmt.Fprintln(br.writer, prime)
	}
}

//func (p *PrimeFinder) println(writer *bufio.Writer, f string, args ...interface{}) {
//	_, err := fmt.Fprintf(writer, f+"\n", args...)
//	if err != nil {
//		return
//	}
//}
