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
		var numberCount int
		fmt.Fscanf(in, "%d\n", &numberCount)
		var numbers = make([]int, numberCount)
		for i := 0; i < numberCount; i++ {
			fmt.Fscan(in, &numbers[i])
		}
		fmt.Fscanf(in, "\n")

		fmt.Fprintf(out, "\n")
	}
}
