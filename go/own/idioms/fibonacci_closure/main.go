package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		// Capture the current value of `a` to return.
		next := a
		// Update `a` and `b` to the next Fibonacci numbers.
		a, b = b, a+b
		return next
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
