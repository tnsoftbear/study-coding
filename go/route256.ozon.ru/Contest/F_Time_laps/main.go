package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

type ByStart []Range

func (r ByStart) Len() int           { return len(r) }
func (r ByStart) Less(i, j int) bool { return r[i].Start < r[j].Start }
func (r ByStart) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscanf(in, "%d\n", &testCount)

	var line string
	var sh, sm, ss, eh, em, es, start, end int

	for testIdx := 0; testIdx < testCount; testIdx++ {
		var rangeCount int
		fmt.Fscanf(in, "%d\n", &rangeCount)

		var ranges = make([]Range, rangeCount)
		var success = true

		for i := 0; i < rangeCount; i++ {
			fmt.Fscanf(in, "%s\n", &line)
			var timesIso = strings.Split(line, "-")
			var startTimes = strings.Split(timesIso[0], ":")
			var endTimes = strings.Split(timesIso[1], ":")
			sh, _ = strconv.Atoi(startTimes[0])
			sm, _ = strconv.Atoi(startTimes[1])
			ss, _ = strconv.Atoi(startTimes[2])
			eh, _ = strconv.Atoi(endTimes[0])
			em, _ = strconv.Atoi(endTimes[1])
			es, _ = strconv.Atoi(endTimes[2])
			if sh < 0 || sh > 23 || eh < 0 || eh > 23 || sm < 0 || sm > 59 || em < 0 || em > 59 || ss < 0 || ss > 59 || es < 0 || es > 59 {
				success = false
			}

			start = sh*60*60 + sm*60 + ss
			end = eh*60*60 + em*60 + es
			if start > end {
				success = false
			}

			if success {
				ranges[i].Start = start
				ranges[i].End = end
			}
		}

		if !success {
			fmt.Fprintf(out, "NO\n")
			continue
		}

		sort.Sort(ByStart(ranges))

		var lastEnd = 0
		for i := 0; i < len(ranges); i++ {
			if lastEnd != 0 && lastEnd >= ranges[i].Start {
				success = false
				break
			}
			lastEnd = ranges[i].End
		}

		if !success {
			fmt.Fprintf(out, "NO\n")
			continue
		}

		fmt.Fprintf(out, "YES\n")
	}
}
