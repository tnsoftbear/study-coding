package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscanf(in, "%d\n", &testCount)

	for testIdx := 0; testIdx < testCount; testIdx++ {
		var dayCount int
		fmt.Fscanf(in, "%d\n", &dayCount)
		var tasks = make(map[int]int, dayCount)
		var runningTaskNr int
		var taskNr int
		var success = true

		var line, _ = in.ReadString('\n')
		var inputTasks = strings.Split(strings.TrimSpace(line), " ")

	testloop:
		for i := 0; i < len(inputTasks); i++ {
			taskNr, _ = strconv.Atoi(inputTasks[i])
			if runningTaskNr == 0 {
				runningTaskNr = taskNr
				tasks[runningTaskNr] = 1
			} else {
				if runningTaskNr != taskNr {
					runningTaskNr = taskNr
					if tasks[runningTaskNr] > 0 {
						success = false
						break testloop
					}
					tasks[runningTaskNr] = 1
				}
			}
		}

		if !success {
			fmt.Fprintf(out, "NO\n")
		} else {
			fmt.Fprintf(out, "YES\n")
		}
	}
}
