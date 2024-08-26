package main

import "fmt"

// Эквивалентные функции

func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

func main() {
	var i1 = newInt1()
	var i2 = newInt2()
	fmt.Printf("%d %d %p %p\n", *i1, *i2, i1, i2)
}
