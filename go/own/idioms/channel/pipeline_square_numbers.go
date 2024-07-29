package _main

import "fmt"

func main() {
	naturalsCh := make(chan int)
	squaresCh := make(chan int)
	go func() {
		for n := 1; n <= 100; n++ {
			naturalsCh <- n
		}
		close(naturalsCh)
	}()

	go func() {
		for {
			n, ok := <-naturalsCh
			if !ok {
				break
			}
			squaresCh <- n * n
		}
		close(squaresCh)
	}()

	// for v, ok := <- squaresCh; ok; v, ok = <- squaresCh {
	// 	fmt.Println(v)
	// }

	for v := range squaresCh {
		fmt.Println(v)
	}
}