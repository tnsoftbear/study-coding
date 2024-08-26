package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch1closed := false
	ch2closed := false
	go func() {
		Walk(t1, ch1)
		close(ch1)
		ch1closed = true
	}()
	go func() {
		Walk(t2, ch2)
		close(ch2)
		ch2closed = true
	}()

	var buf1 []int
	var buf2 []int
	forloop: for {
		select {
		case v1, ok := <-ch1:
			if ok {
				buf1 = append(buf1, v1)
			} else {
				if ch2closed {
					break forloop
				}
			}
		case v2, ok := <-ch2:
			if ok {
				buf2 = append(buf2, v2)
			} else {
				if ch1closed {
					break forloop
				}
			}
		}
	}
	return slicesEqual(buf1, buf2)
}

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func main() {
	t1 := tree.New(1)
	t2 := tree.New(1)
	res1 := Same(t1, t2)
	fmt.Printf("Result: %v\n", res1)
}
