package profiler

import (
	"fmt"
	"time"
	"bufio"
)

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

func (pr *Profiler) printMyself(writer *bufio.Writer) {
	if !pr.isEnabled {
		return
	}
	for idx := range pr.startTime {
		fmt.Fprintf(writer, "Iteration: %d, Elapsed time (sec): %d\n", idx, pr.elapsedTime(idx))
	}
	writer.Flush()
}