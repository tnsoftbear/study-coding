package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscanf(in, "%d\n", &testCount)

	for i := 0; i < testCount; i++ {
		var numberCount int
		fmt.Fscanf(in, "%d\n", &numberCount)
		var numbers = make([]int, numberCount)
		for j := 0; j < numberCount; j++ {
			fmt.Fscan(in, &numbers[j])
		}
		fmt.Fscanf(in, "\n")

		var aggregate = make(map[int]int)
		for j := 0; j < numberCount; j++ {
			_, exists := aggregate[numbers[j]]
			if exists {
				aggregate[numbers[j]]++
			} else {
				aggregate[numbers[j]] = 1
			}
		}

		var sum = 0
		var total = 0
		for idx, cnt := range aggregate {
			var full = int(math.Floor(float64(cnt) / 3))
			sum = full*2*idx + (cnt-full*3)*idx
			total += sum
		}

		fmt.Fprintf(out, "%v\n", total)
	}
}
