package main

import "fmt"

type AbstractPluser interface {
	plus(int) int
}

type OnePluser struct{}

func (d OnePluser) plus(x int) int {
	return x + 1
}

type TwoPluser struct{}

func (d TwoPluser) plus(x int) int {
	return x + 2
}

func runTwice(pluser1 AbstractPluser, pluser2 AbstractPluser) int {
	return pluser1.plus(1) + pluser2.plus(1)
}

func main() {
	fmt.Println(runTwice(OnePluser{}, TwoPluser{}))
}
