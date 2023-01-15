package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

const SUCCESS = "SUCCESS\n"
const FAIL = "FAIL\n"
const SuccessNumTpl = "SUCCESS %d-%d\n"

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

type Place struct {
	Nr     int
	IsFree bool
}

type Coupe struct {
	P1, P2 Place
}

func (c Coupe) IsFree() bool {
	return c.P1.IsFree && c.P2.IsFree
}

func (c *Coupe) Occupy() {
	c.P1.IsFree = false
	c.P2.IsFree = false
}

type CoupeHeap []Coupe

func (ch CoupeHeap) Len() int {
	return len(ch)
}
func (ch CoupeHeap) Less(i, j int) bool {
	if !ch[i].IsFree() && ch[j].IsFree() {
		return false
	}
	if ch[i].IsFree() && !ch[j].IsFree() {
		return true
	}
	return ch[i].P1.Nr < ch[j].P1.Nr
}

func (ch CoupeHeap) Swap(i, j int) {
	ch[i], ch[j] = ch[j], ch[i]
}

func (ch *CoupeHeap) Push(v interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*ch = append(*ch, v.(Coupe))
}

func (ch *CoupeHeap) OccupyPlace(place int, statuses *[]bool) bool {
	if !(*statuses)[place] {
		(*statuses)[place] = true
		return true
	}
	return false
}

func (ch *CoupeHeap) FreePlace(place int, statuses *[]bool) bool {
	if (*statuses)[place] {
		(*statuses)[place] = false
		if place%2 == 0 {
			if !(*statuses)[place+1] {
				heap.Push(ch, Coupe{P1: Place{Nr: place, IsFree: true}, P2: Place{Nr: place + 1, IsFree: true}})
			}
		} else {
			if !(*statuses)[place-1] {
				heap.Push(ch, Coupe{P1: Place{Nr: place - 1, IsFree: true}, P2: Place{Nr: place, IsFree: true}})
			}
		}
		return true
	}
	return false
}

func (ch *CoupeHeap) Pop() interface{} {
	old := *ch
	n := len(old)
	x := old[n-1]
	*ch = old[0 : n-1]
	return x
}

func (ch *CoupeHeap) Peek() Coupe {
	old := *ch
	n := len(old)
	x := old[n-1]
	return x
}

func main() {
	defer out.Flush()

	var testCount int
	fmt.Fscanf(in, "%d\n", &testCount)

	for testI := 0; testI < testCount; testI++ {
		fmt.Fscanf(in, "\n")

		var total, query int
		fmt.Fscanf(in, "%d %d\n", &total, &query)

		//var coupes = make(CoupeHeap, total)
		var coupes = &CoupeHeap{}
		for i := 0; i < total; i++ {
			var coupe = Coupe{
				P1: Place{Nr: i * 2, IsFree: true},
				P2: Place{Nr: i*2 + 1, IsFree: true}}
			heap.Push(coupes, coupe)
		}

		var statuses = make([]bool, total*2)

	mainloop:
		for i := 0; i < query; i++ {
			var action, target int
			fmt.Fscanf(in, "%d", &action)
			if action < 3 {
				fmt.Fscanf(in, "%d\n", &target)
			} else {
				fmt.Fscanf(in, "\n")
			}

			target--

			if action == 1 {
				if coupes.OccupyPlace(target, &statuses) {
					fmt.Fprintf(out, SUCCESS)
				} else {
					fmt.Fprintf(out, FAIL)
				}
			} else if action == 2 {
				if coupes.FreePlace(target, &statuses) {
					fmt.Fprintf(out, SUCCESS)
				} else {
					fmt.Fprintf(out, FAIL)
				}
			} else {

				var coupe Coupe
				for {
					if coupes.Len() == 0 {
						fmt.Fprintf(out, FAIL)
						continue mainloop
					}
					coupe = heap.Pop(coupes).(Coupe)
					if coupe.IsFree() && !statuses[coupe.P1.Nr] && !statuses[coupe.P2.Nr] {
						break
					}
				}
				fmt.Fprintf(out, SuccessNumTpl, coupe.P1.Nr+1, coupe.P2.Nr+1)
				statuses[coupe.P1.Nr] = true
				statuses[coupe.P2.Nr] = true
			}
		}
		fmt.Fprintln(out)
	}
}
