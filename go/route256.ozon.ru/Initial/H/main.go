package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const EMPTY byte = 0
const SOMETHING byte = 1
const EDGE byte = 2
const LINE byte = 3
const CORNER byte = 4
const SINGLE byte = 5
const PROCESSED byte = 6

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)
var height, width int

type NearStatuses struct {
	x           int
	y           int
	status      byte
	left        byte
	right       byte
	top         byte
	bottom      byte
	leftTop     byte
	rightTop    byte
	leftBottom  byte
	rightBottom byte
}

func NewNearStatuses(x, y int, statuses [][]byte) NearStatuses {
	var ns = NearStatuses{
		x:           x,
		y:           y,
		status:      EMPTY,
		left:        EMPTY,
		right:       EMPTY,
		top:         EMPTY,
		bottom:      EMPTY,
		leftTop:     EMPTY,
		rightTop:    EMPTY,
		leftBottom:  EMPTY,
		rightBottom: EMPTY,
	}
	ns.detect(x, y, statuses)
	return ns
}

func (ns *NearStatuses) detect(x, y int, statuses [][]byte) {
	ns.status = statuses[y][x]

	if x > 0 {
		ns.left = statuses[y][x-1]
	}

	if x < width-1 {
		ns.right = statuses[y][x+1]
	}

	if y > 0 {
		ns.top = statuses[y-1][x]
	}

	if y < height-1 {
		ns.bottom = statuses[y+1][x]
	}

	if x > 0 && y > 0 {
		ns.leftTop = statuses[y-1][x-1]
	}

	if x < width-1 && y > 0 {
		ns.rightTop = statuses[y-1][x+1]
	}

	if x > 0 && y < height-1 {
		ns.leftBottom = statuses[y+1][x-1]
	}

	if x < width-1 && y < height-1 {
		ns.rightBottom = statuses[y+1][x+1]
	}
}

func (ns NearStatuses) isAnyShipDetail(detail byte) bool {
	return detail != EMPTY
}

func (ns NearStatuses) isEmptyDiagonals() bool {
	return ns.leftTop == EMPTY && ns.rightTop == EMPTY && ns.leftBottom == EMPTY && ns.rightBottom == EMPTY
}

func (ns NearStatuses) isSingle() bool {
	if ns.status == SINGLE {
		return true
	}
	if ns.status != SOMETHING {
		return false
	}
	if !ns.isEmptyDiagonals() {
		return false
	}
	var isSingle = ns.left == EMPTY && ns.right == EMPTY && ns.top == EMPTY && ns.bottom == EMPTY
	return isSingle
}

func (ns NearStatuses) isEdge() bool {
	//ns.PrintMe()
	if ns.status == EDGE {
		return true
	}
	if ns.status != SOMETHING {
		return false
	}

	if ns.isAnyShipDetail(ns.left) && ns.right == EMPTY && ns.top == EMPTY && ns.bottom == EMPTY && ns.rightTop == EMPTY && ns.rightBottom == EMPTY {
		return true
	}

	if ns.left == EMPTY && ns.isAnyShipDetail(ns.right) && ns.top == EMPTY && ns.bottom == EMPTY && ns.leftTop == EMPTY && ns.leftBottom == EMPTY {
		return true
	}

	if ns.left == EMPTY && ns.right == EMPTY && ns.isAnyShipDetail(ns.top) && ns.bottom == EMPTY && ns.leftBottom == EMPTY && ns.rightBottom == EMPTY {
		return true
	}

	if ns.left == EMPTY && ns.right == EMPTY && ns.top == EMPTY && ns.isAnyShipDetail(ns.bottom) && ns.leftTop == EMPTY && ns.rightTop == EMPTY {
		return true
	}

	return false
}

func (ns NearStatuses) isLine() bool {
	if ns.status == LINE {
		return true
	}
	if ns.status != SOMETHING {
		return false
	}

	if ns.isAnyShipDetail(ns.left) && ns.isAnyShipDetail(ns.right) && ns.top == EMPTY && ns.bottom == EMPTY {
		return true
	}

	if ns.left == EMPTY && ns.right == EMPTY && ns.isAnyShipDetail(ns.top) && ns.isAnyShipDetail(ns.bottom) {
		return true
	}

	return false
}

func (ns NearStatuses) isCorner() bool {
	if ns.status == CORNER {
		return true
	}
	if ns.status != SOMETHING {
		return false
	}
	if !ns.isEmptyDiagonals() {
		return false
	}

	if ns.isAnyShipDetail(ns.left) && ns.right == EMPTY && ns.isAnyShipDetail(ns.top) && ns.bottom == EMPTY {
		return true
	}

	if ns.isAnyShipDetail(ns.left) && ns.right == EMPTY && ns.top == EMPTY && ns.isAnyShipDetail(ns.bottom) {
		return true
	}

	if ns.left == EMPTY && ns.isAnyShipDetail(ns.right) && ns.isAnyShipDetail(ns.top) && ns.bottom == EMPTY {
		return true
	}

	if ns.left == EMPTY && ns.isAnyShipDetail(ns.right) && ns.top == EMPTY && ns.isAnyShipDetail(ns.bottom) {
		return true
	}

	return false
}

func (ns NearStatuses) PrintMe() {
	fmt.Fprintf(out, "x:%d,y:%d,s:%d - l:%d,r:%d,t:%d,b:%d,lt:%d,rt:%d,lb:%d,rb:%d\n", ns.x, ns.y, ns.status, ns.left, ns.right, ns.top, ns.bottom, ns.leftTop, ns.rightTop, ns.leftBottom, ns.rightBottom)
}

func main() {
	defer out.Flush()

	var testCount int
	fmt.Fscanf(in, "%d\n", &testCount)

	for testI := 0; testI < testCount; testI++ {
		fmt.Fscanf(in, "%d %d\n", &height, &width)

		var statuses = make([][]byte, height)
		for i := 0; i < height; i++ {
			statuses[i] = make([]byte, width)
		}

		var line string

		// fmt.Fprintf(out, "h: %d, w: %d\n", height, width)

		for i := 0; i < height; i++ {
			fmt.Fscanf(in, "%s\n", &line)
			//fmt.Fprintf(out, "%s\n", line)

			for j, c := range line {
				if string(c) == "*" {
					statuses[i][j] = SOMETHING
				} else if string(c) == "." {
					statuses[i][j] = EMPTY
				}
			}
		}

		var ns NearStatuses
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				ns = NewNearStatuses(x, y, statuses)
				if ns.isSingle() {
					statuses[y][x] = SINGLE
				} else if ns.isEdge() {
					statuses[y][x] = EDGE
				} else if ns.isLine() {
					statuses[y][x] = LINE
				} else if ns.isCorner() {
					statuses[y][x] = CORNER
				}
			}
		}

		//for y := 0; y < height; y++ {
		//	for x := 0; x < width; x++ {
		//		fmt.Fprintf(out, "%d ", statuses[y][x])
		//	}
		//	fmt.Fprintln(out)
		//}

		var success = true
	findsomething:
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if statuses[y][x] == SOMETHING {
					success = false
					break findsomething
				}
			}
		}

		if !success {
			fmt.Fprintln(out, "NO")
			continue
		}

		success = true
		var sizes []int
	mainloop:
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				if statuses[y][x] == SINGLE {
					sizes = append(sizes, 1)
				} else if statuses[y][x] == CORNER {
					var size = detectShipByCorner(x, y, &statuses)
					if size == 0 {
						success = false
						break mainloop
					} else {
						sizes = append(sizes, size)
					}
				} else if statuses[y][x] == EDGE {
					var size = detectShipByEdge(x, y, &statuses)
					if size == 0 {
						success = false
						break mainloop
					} else {
						sizes = append(sizes, size)
					}
				}
			}
		}

		if !success || len(sizes) == 0 {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
			sort.Ints(sizes)
			for _, v := range sizes {
				fmt.Fprintf(out, "%d ", v)
			}
			fmt.Fprintln(out)
		}
	}
}

func detectShipByCorner(x, y int, statuses *[][]byte) int {
	var ns = NewNearStatuses(x, y, *statuses)
	(*statuses)[y][x] = PROCESSED
	var rightLen = 0
	if ns.right != EMPTY {
		rightLen = 1
		for i := x + 1; i < width; i++ {
			if (*statuses)[y][i] == LINE {
				rightLen++
				(*statuses)[y][i] = PROCESSED
			} else if (*statuses)[y][i] == EDGE {
				rightLen++
				(*statuses)[y][i] = PROCESSED
			} else if (*statuses)[y][i] == CORNER {
				return 0
			} else {
				break
			}
		}
	}

	var bottomLen = 0
	if ns.bottom != EMPTY {
		bottomLen = 1
		for j := y + 1; j < height; j++ {
			if (*statuses)[j][x] == LINE {
				bottomLen++
				(*statuses)[j][x] = PROCESSED
			} else if (*statuses)[j][x] == EDGE {
				bottomLen++
				(*statuses)[j][x] = PROCESSED
			} else if (*statuses)[j][x] == CORNER {
				return 0
			} else {
				break
			}
		}
	}

	if rightLen != bottomLen {
		// fmt.Fprintf(out, "1] rightLen (%d) != bottomLen (%d) x:%d, y:%d\n", rightLen, bottomLen, x, y)
		return 0
	}

	if rightLen == 0 {
		// fmt.Fprintf(out, "rightLen == 0\n")
		return 0
	}

	return rightLen*2 - 1
}

func detectShipByEdge(x, y int, statuses *[][]byte) int {
	var ns = NewNearStatuses(x, y, *statuses)
	var rightLen = 0
	var bottomLen = 0
	(*statuses)[y][x] = PROCESSED
	if ns.right != EMPTY {
		rightLen = 1
		for i := x + 1; i < width; i++ {
			rightLen++
			if (*statuses)[y][i] == LINE {
				(*statuses)[y][i] = PROCESSED
			} else if (*statuses)[y][i] == CORNER {
				(*statuses)[y][i] = PROCESSED
				var nsCorner = NewNearStatuses(i, y, *statuses)
				var bottomLen = 0
				if nsCorner.bottom != EMPTY {
					bottomLen = 1
					for j := y + 1; j < height; j++ {
						if (*statuses)[j][i] == LINE {
							bottomLen++
							(*statuses)[j][i] = PROCESSED
						} else if (*statuses)[j][i] == EDGE {
							bottomLen++
							(*statuses)[j][i] = PROCESSED
							if rightLen != bottomLen {
								// fmt.Fprintf(out, "2] rightLen (%d) != bottomLen (%d)\n", rightLen, bottomLen)
								return 0
							}
							return rightLen*2 - 1
						} else if (*statuses)[j][i] == CORNER {
							return 0
						}
					}
				}
			} else if (*statuses)[y][i] == EMPTY {
				return 0
			}
		}
	} else if ns.bottom != EMPTY {
		bottomLen = 1
		for j := y + 1; j < height; j++ {
			if (*statuses)[j][x] == LINE {
				bottomLen++
				(*statuses)[j][x] = PROCESSED
			} else if (*statuses)[j][x] == CORNER {
				bottomLen++
				(*statuses)[j][x] = PROCESSED
				var nsCorner = NewNearStatuses(x, j, *statuses)
				if nsCorner.right != EMPTY {
					var rightLen = 1
					for i := x + 1; i < width; i++ {
						if (*statuses)[j][i] == LINE {
							rightLen++
							(*statuses)[j][i] = PROCESSED
						} else if (*statuses)[j][i] == EDGE {
							rightLen++
							(*statuses)[j][i] = PROCESSED
							if rightLen != bottomLen {
								// fmt.Fprintf(out, "3] rightLen (%d) != bottomLen (%d)\n", rightLen, bottomLen)
								return 0
							}
							return rightLen*2 - 1
						} else if (*statuses)[j][i] == CORNER {
							return 0
						}
					}
				} else if nsCorner.left != EMPTY {
					var leftLen = 1
					for i := x - 1; i >= 0; i-- {
						if (*statuses)[j][i] == LINE {
							leftLen++
							(*statuses)[j][i] = PROCESSED
						} else if (*statuses)[j][i] == EDGE {
							leftLen++
							(*statuses)[j][i] = PROCESSED
							if leftLen != bottomLen {
								// fmt.Fprintf(out, "leftLen (%d) != bottomLen (%d)\n", leftLen, bottomLen)
								return 0
							}
							return leftLen*2 - 1
						} else if (*statuses)[j][i] == CORNER {
							return 0
						}
					}
				}
			} else if (*statuses)[j][x] == EMPTY {
				return 0
			}
		}
	}
	return 0
}
