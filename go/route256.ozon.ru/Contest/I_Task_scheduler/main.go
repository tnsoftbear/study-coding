package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CpuStatRecord struct {
	Nr           int
	IsIdle       bool
	BusyUntilIdx int
	LoadUnit     int
	LoadUsed     int64
}

type CsrHeap []CpuStatRecord

func (ch CsrHeap) Len() int {
	return len(ch)
}
func (ch CsrHeap) Less(i, j int) bool {
	if !ch[i].IsIdle && ch[j].IsIdle {
		return false
	}
	if ch[i].IsIdle && !ch[j].IsIdle {
		return true
	}
	return ch[i].LoadUnit < ch[j].LoadUnit
}

func (ch CsrHeap) Swap(i, j int) {
	ch[i], ch[j] = ch[j], ch[i]
}

func (ch *CsrHeap) Push(v interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*ch = append(*ch, v.(CpuStatRecord))
}

func (ch *CsrHeap) Pop() interface{} {
	old := *ch
	n := len(old)
	x := old[n-1]
	*ch = old[0 : n-1]
	return x
}

func (ch *CsrHeap) Peek() CpuStatRecord {
	old := *ch
	n := len(old)
	x := old[n-1]
	return x
}

func (ch CsrHeap) findEffectiveCpu() CpuStatRecord {
	var effectiveCsr CpuStatRecord
	var effectiveIdx int = -1
	for i, csr := range ch {
		if csr.IsIdle {
			if effectiveIdx == -1 || effectiveCsr.LoadUnit > csr.LoadUnit {
				effectiveCsr = csr
				effectiveIdx = i
			}
		}
	}
	return effectiveCsr
}

func (ch *CsrHeap) freeFinishedTasks(taskTick int) {
	for i, csr := range *ch {
		if !csr.IsIdle {
			if csr.BusyUntilIdx <= taskTick {
				(*ch)[i].IsIdle = true
				(*ch)[i].BusyUntilIdx = 0
			}
		}
	}
}

func (ch *CsrHeap) busy(effectiveCsr CpuStatRecord, taskTick, taskLength int) {
	for i, csr := range *ch {
		if csr.Nr == effectiveCsr.Nr {
			(*ch)[i].IsIdle = false
			(*ch)[i].BusyUntilIdx = taskTick + taskLength
			(*ch)[i].LoadUsed += int64((*ch)[i].LoadUnit) * int64(taskLength)
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var cpuCount, taskCount, loadUnit int
	fmt.Fscanf(in, "%d %d\n", &cpuCount, &taskCount)

	var cpuStats = &CsrHeap{}

	var line, _ = in.ReadString('\n')
	var inputTasks = strings.Split(strings.TrimSpace(line), " ")

	for i := 0; i < len(inputTasks); i++ {
		loadUnit, _ = strconv.Atoi(inputTasks[i])
		var csr = CpuStatRecord{
			Nr:           i,
			IsIdle:       true,
			BusyUntilIdx: 0,
			LoadUnit:     loadUnit,
			LoadUsed:     0}
		heap.Push(cpuStats, csr)
	}

	var taskTick, taskLength int

	for taskIdx := 0; taskIdx < taskCount; taskIdx++ {
		var line, _ = in.ReadString('\n')
		var inputTasks = strings.Split(strings.TrimSpace(line), " ")
		taskTick, _ = strconv.Atoi(inputTasks[0])
		taskLength, _ = strconv.Atoi(inputTasks[1])
		cpuStats.freeFinishedTasks(taskTick)
		var effectiveCsr = cpuStats.findEffectiveCpu()
		if effectiveCsr.LoadUnit > 0 {
			cpuStats.busy(effectiveCsr, taskTick, taskLength)
		}
		//fmt.Fprintf(out, "taskIdx: %d, TaskTick: %d, TaskLen: %d, CpuStats: %v\n", taskIdx, taskTick, taskLength, cpuStats)
	}

	var loadSum int64
	for _, csr := range *cpuStats {
		loadSum += csr.LoadUsed
	}

	fmt.Fprintf(out, "%d\r\n", loadSum)
}
