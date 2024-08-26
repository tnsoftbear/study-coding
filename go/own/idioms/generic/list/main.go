package main

import "fmt"

// List represents a singly-linked list that holds values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
	lst1 := List[uint8]{
		next: nil,
		val:  10,
	}
	lst2 := List[uint8]{
		next: &lst1,
		val:  20,
	}
	fmt.Printf("lst1: %v, lst2: %v", lst1, lst2)
}
