package profiler

import (
	"bufio"
	"fmt"
	"time"
)

type Profiler struct {
	isEnabled bool
	startTime []int64
	endTime   []int64
}

func New(isEnabled bool) *Profiler {
	return &Profiler{isEnabled: isEnabled}
}

func (pr *Profiler) Start() {
	if !pr.isEnabled {
		return
	}
	pr.startTime = append(pr.startTime, time.Now().UnixMilli())
}

func (pr *Profiler) End() {
	if !pr.isEnabled {
		return
	}
	pr.endTime = append(pr.endTime, time.Now().UnixMilli())
}

func (pr *Profiler) ElapsedTime(idx int) int64 {
	if !pr.isEnabled {
		return 0
	}
	return pr.endTime[idx] - pr.startTime[idx]
}

func (pr *Profiler) PrintMyself(writer *bufio.Writer) {
	if !pr.isEnabled {
		return
	}
	for idx := range pr.startTime {
		pr.println(writer, "Iteration: %d, Elapsed time (sec): %d", idx, pr.ElapsedTime(idx))
	}
	err := writer.Flush()
	if err != nil {
		pr.println(writer, "Flush() call failed")
		return
	}
}

func (pr *Profiler) println(writer *bufio.Writer, f string, args ...interface{}) {
	_, err := fmt.Fprintf(writer, f+"\n", args...)
	if err != nil {
		return
	}
}
