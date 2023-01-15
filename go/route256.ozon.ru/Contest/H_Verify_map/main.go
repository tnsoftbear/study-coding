package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

const EMPTY byte = 0
const RED byte = 1
const GREEN byte = 2
const VIOLET byte = 3
const YELLOW byte = 4
const BLUE byte = 5

type FieldStatus struct {
	Type byte
	Nr   byte
	X    int
	Y    int
}

var Colors = map[byte]byte{}

func NewColors() map[byte]byte {
	return map[byte]byte{
		EMPTY:  0,
		RED:    0,
		GREEN:  0,
		VIOLET: 0,
		YELLOW: 0,
		BLUE:   0}
}

func printMe(fields [][]FieldStatus) {
	for _, ff := range fields {
		for _, f := range ff {
			fmt.Fprintf(out, "%d:%d ", f.Type, f.Nr)
		}
		fmt.Fprintln(out)
	}
	fmt.Fprintln(out)
}

func findNrFromNearAndUpdateNear(current FieldStatus, width, height int, fields *[][]FieldStatus) bool {
	var left, right, leftTop, rightTop, leftBottom, rightBottom FieldStatus
	var nearFieldsToUpdate = make([]FieldStatus, 6)
	var nearIdx = 0
	var x = current.X
	var y = current.Y

	if x > 1 {
		left = (*fields)[y][x-2]
		if left.Type == current.Type {
			if left.Nr > 0 {
				current.Nr = left.Nr
			} else {
				nearFieldsToUpdate[nearIdx] = left
				nearIdx++
			}
		}
	}

	if x < width-2 {
		right = (*fields)[y][x+2]
		if right.Type == current.Type {
			if right.Nr > 0 {
				current.Nr = right.Nr
			} else {
				nearFieldsToUpdate[nearIdx] = right
				nearIdx++
			}
		}
	}

	if y > 0 && x > 0 {
		leftTop = (*fields)[y-1][x-1]
		if leftTop.Type == current.Type {
			if leftTop.Nr > 0 {
				current.Nr = leftTop.Nr
			} else {
				nearFieldsToUpdate[nearIdx] = leftTop
				nearIdx++
			}
		}
	}

	if y > 0 && x < width-1 {
		rightTop = (*fields)[y-1][x+1]
		if rightTop.Type == current.Type {
			if rightTop.Nr > 0 {
				current.Nr = rightTop.Nr
			} else {
				nearFieldsToUpdate[nearIdx] = rightTop
				nearIdx++
			}
		}
	}

	if y < height-1 && x > 0 {
		leftBottom = (*fields)[y+1][x-1]
		if leftBottom.Type == current.Type {
			if leftBottom.Nr > 0 {
				current.Nr = leftBottom.Nr
			} else {
				nearFieldsToUpdate[nearIdx] = leftBottom
				nearIdx++
			}
		}
	}

	if y < height-1 && x < width-1 {
		rightBottom = (*fields)[y+1][x+1]
		if rightBottom.Type == current.Type {
			if rightBottom.Nr > 0 {
				current.Nr = rightBottom.Nr
			} else {
				nearFieldsToUpdate[nearIdx] = rightBottom
				nearIdx++
			}
		}
	}

	if current.Nr == 0 {
		Colors[current.Type]++
		current.Nr = Colors[current.Type]
	}
	(*fields)[y][x] = current
	if current.Nr > 1 {
		return false
	}

	for i := 0; i < nearIdx; i++ {
		var success = findNrFromNearAndUpdateNear(nearFieldsToUpdate[i], width, height, fields)
		if !success {
			return false
		}
	}
	return true
}

func main() {
	defer out.Flush()

	var testCount int
	fmt.Fscanf(in, "%d\n", &testCount)

	for testIdx := 0; testIdx < testCount; testIdx++ {

		Colors = NewColors()

		var height, width int
		fmt.Fscanf(in, "%d %d\n", &height, &width)

		var fields = make([][]FieldStatus, height)
		for i := 0; i < height; i++ {
			fields[i] = make([]FieldStatus, width)
		}

		var line string
		for y := 0; y < height; y++ {
			fmt.Fscanf(in, "%s\n", &line)

			for x, c := range line {
				if c == 46 {
					fields[y][x] = FieldStatus{Type: EMPTY, X: x, Y: y}
				} else {
					fields[y][x] = FieldStatus{Type: byte(c), X: x, Y: y}
				}
			}
		}

		// printMe(fields)

		var success = true
	findsomething:
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if fields[y][x].Type == EMPTY {
					continue
				}
				if fields[y][x].Nr == 0 {
					success = findNrFromNearAndUpdateNear(fields[y][x], width, height, &fields)
					if !success {
						break findsomething
					}
				}
			}
		}

		// fmt.Fprintf(out, "%v\n", fields)

		if !success {
			fmt.Fprintln(out, "NO")
			continue
		}
		fmt.Fprintln(out, "YES")

		// fmt.Fprintf(out, "\n")
	}
}
