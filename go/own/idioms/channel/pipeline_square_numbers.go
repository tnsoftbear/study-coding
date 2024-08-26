package main

import "fmt"

func counter(naturalsCh chan<- int) {
	for n := 1; n <= 100; n++ {
		naturalsCh <- n
	}
	close(naturalsCh)
}

func squarer(naturalsCh <-chan int, squaresCh chan<- int) {
	for {
		n, ok := <-naturalsCh
		if !ok {
			break
		}
		squaresCh <- n * n
	}
	close(squaresCh)
}

func printer(squaresCh <-chan int) {
	for v := range squaresCh {
		fmt.Println(v)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch1, ch2)
	printer(ch2)
}
