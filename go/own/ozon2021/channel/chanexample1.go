package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 10)
	// ch <- 3
	a, ok := <-ch
	//ch <- 1
	if ok {
		fmt.Println(a)
	}
	time.Sleep(time.Second * 10)

	//close(ch)
	//b, ok := <-ch
	//fmt.Println(b)
	//
	//c, ok := <-ch
	//if ok {
	//	fmt.Println(c)
	//}

}
