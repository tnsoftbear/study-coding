package main

import (
	_ "fmt"
	"math/rand"

	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var ySlice [][]uint8
	for y := 0; y < dy; y++ {
		xSlice := make([]uint8, dx)
		for i := range xSlice {
			xSlice[i] = uint8(rand.Intn(256))
		}
		ySlice = append(ySlice, xSlice)
	}
	return ySlice
}

func main() {
	pic.Show(Pic)
	// fmt.Printf("%v", Pic(3, 3))
}
