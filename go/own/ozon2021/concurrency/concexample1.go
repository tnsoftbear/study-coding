package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go func(counter int) {
			time.Sleep(time.Second)
			fmt.Println(counter)
		}(i)
	}
	time.Sleep(time.Second * 2)
}

// Здесь быстро проскочит инкрементация счетчика, а потом выведется 100 во всех зарегистрированных горутинах.
//func main() {
//	for i := 0; i < 100; i++ {
//		go func() {
//			time.Sleep(time.Second)
//			fmt.Println(i)
//		}()
//	}
//	time.Sleep(time.Second * 2)
//}

// Выставить кол-во проц. GOMAXPROCS и переключать горутину Gosched:
//func main() {
//	// runtime.GOMAXPROCS(1)
//	for i := 0; i < 10; i++ {
//		go func() {
//			fmt.Println(i)
//		}()
//		runtime.Gosched()
//	}
//	time.Sleep(time.Second * 2)
//}
