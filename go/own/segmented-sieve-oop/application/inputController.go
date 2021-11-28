package application

import (
	"bufio"
	"fmt"
	"segmented-sieve-oop/types"
)

type InputController struct {
	caseCount int
	ranges    [150]types.Range
	reader    *bufio.Reader
}

func NewInputController(reader *bufio.Reader) *InputController {
	return &InputController{reader: reader}
}

func (ic *InputController) Read() {
	var min, max uint32
	ic.scanf("%d\n", &ic.caseCount)
	for i := 0; i < ic.caseCount; i++ {
		ic.scanf("%d %d\n", &min, &max)
		ic.ranges[i] = types.Range{Min: min, Max: max}
	}
}

func (ic *InputController) RangeByIndex(idx int) types.Range {
	return ic.ranges[idx]
}

func (ic *InputController) CaseCount() int {
	return ic.caseCount
}

func (ic *InputController) scanf(f string, a ...interface{}) { fmt.Fscanf(ic.reader, f, a...) }
