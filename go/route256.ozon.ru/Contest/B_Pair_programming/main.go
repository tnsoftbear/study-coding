package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscanf(in, "%d\n", &testCount)

	for testIdx := 0; testIdx < testCount; testIdx++ {

		var devCount int
		fmt.Fscanf(in, "%d\n", &devCount)

		var levels = make([]int, devCount)

		for i := 0; i < devCount; i++ {
			fmt.Fscanf(in, "%d ", &levels[i])
		}
		fmt.Fscanf(in, "\n")

		var checkingLevel int
		var pairs = make([][2]int, devCount/2)
		var pairIdx int
		for i := 0; i < devCount; i++ {
			checkingLevel = levels[i]
			if checkingLevel == -1 {
				continue
			}

			levels[i] = -1

			var lowAbs int = -1
			var diff int
			var pair = [2]int{i + 1}
			for j := i + 1; j < devCount; j++ {
				if levels[j] == -1 {
					continue
				}
				diff = checkingLevel - levels[j]
				if diff < 0 {
					diff *= -1
				}
				if lowAbs == -1 || lowAbs > diff {
					lowAbs = diff
				}
			}

			for j := i + 1; j < devCount; j++ {
				if levels[j] == -1 {
					continue
				}
				diff = checkingLevel - levels[j]
				if diff < 0 {
					diff *= -1
				}
				if diff == lowAbs {
					pair[1] = j + 1
					levels[j] = -1
					break
				}
			}
			pairs[pairIdx] = pair
			pairIdx++
		}

		for _, pair := range pairs {
			fmt.Fprintf(out, "%d %d\n", pair[0], pair[1])
		}
		fmt.Fprintln(out)
	}
}
