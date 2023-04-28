package main

func main() {
	print(maximalRectangle([][]int{{1, 1, 1}, {0, 1, 1}, {1, 0, 0}}))
}

/**
 * @input A : 2D integer array
 *
 * @Output Integer
 */
func maximalRectangle(A [][]int) int {
	var maxArea int = 0
	for y := 0; y < len(A); y++ {
		for x := 0; x < len(A[0]); x++ {
			maxWidth := findMaxWidth(A, x, y)
			maxHeight := findMaxHeight(A, x, y)
			if maxWidth*maxHeight <= maxArea {
				continue
			}

			for hh := maxHeight; hh > 0; hh-- {
				for ww := maxWidth; ww > 0; ww-- {
					area := ww * hh
					if area <= maxArea {
						break
					}
					if isRectangle(A, x, y, ww, hh) {
						area := ww * hh
						if area > maxArea {
							maxArea = area
						}
					}
				}
			}
		}
	}
	return maxArea
}

func findMaxWidth(A [][]int, x, y int) int {
	fullWidth := len(A[y])
	for i := x; i < fullWidth; i++ {
		if A[y][i] == 0 {
			return i - x
		}
	}
	return fullWidth - x
}

func findMaxHeight(A [][]int, x, y int) int {
	fullHeight := len(A)
	for i := y; i < fullHeight; i++ {
		if A[i][x] == 0 {
			return i - y
		}
	}
	return fullHeight - y
}

func isRectangle(A [][]int, x, y, w, h int) bool {
	if A[y][x] == 0 {
		return false
	}
	var width int = 0
	for i := x; i < x+w; i++ {
		if A[y][i] == 0 {
			break
		}
		width++
	}
	var height int = 0
	for i := y; i < y+h; i++ {
		if A[i][x] == 0 {
			break
		}
		height++
	}
	for i := y; i < y+height; i++ {
		for j := x; j < x+width; j++ {
			if A[i][j] == 0 {
				return false
			}
		}
	}
	return true
}
