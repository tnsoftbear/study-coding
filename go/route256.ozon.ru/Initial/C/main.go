package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

type Matrix struct {
	height      int
	width       int
	Data        [][]int
	ClickColumn int
}

func NewMatrix(height int, width int) *Matrix {
	var m = &Matrix{height: height, width: width}
	m.Data = make([][]int, height)
	for i := 0; i < height; i++ {
		m.Data[i] = make([]int, width)
	}
	return m
}
func (m Matrix) Len() int { return m.height }
func (m Matrix) Less(i, j int) bool {
	return m.Data[i][m.ClickColumn] < m.Data[j][m.ClickColumn]
}
func (m *Matrix) Swap(i, j int) {
	m.Data[i], m.Data[j] = m.Data[j], m.Data[i]
}
func (m *Matrix) PrintMe(out io.Writer) {
	for i := 0; i < m.height; i++ {
		for j := 0; j < m.width; j++ {
			fmt.Fprintf(out, "%d", m.Data[i][j])
			if j == m.width-1 {
				fmt.Fprint(out, "\n")
			} else {
				fmt.Fprint(out, " ")
			}
		}
	}
	fmt.Fprint(out, "\n")
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var testCount int
	fmt.Fscanf(in, "%d\n\n", &testCount)

	for testI := 0; testI < testCount; testI++ {
		var height, width int
		fmt.Fscanf(in, "%d %d\n", &height, &width)

		var matrix = NewMatrix(height, width)

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				fmt.Fscanf(in, "%d", &matrix.Data[i][j])
			}
			fmt.Fscanf(in, "\n")
		}

		var clickCount int
		fmt.Fscanf(in, "%d\n", &clickCount)
		clickedColumns := make([]int, clickCount)
		for k := 0; k < clickCount; k++ {
			fmt.Fscanf(in, "%d", &clickedColumns[k])
		}
		fmt.Fscanf(in, "\n")
		fmt.Fscanf(in, "\n")

		for k := 0; k < clickCount; k++ {
			matrix.ClickColumn = clickedColumns[k] - 1
			sort.Stable(matrix)
		}

		matrix.PrintMe(out)
	}
}
