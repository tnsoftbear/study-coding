package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
	"math/rand"
)

const TIME_FORMAT = "15:04:05.0000"
type Circuit func(context.Context) (string, error)

func DebounceFirst(circuit Circuit, d time.Duration) Circuit {
	var threshold time.Time
	var result string
	var err error
	var m sync.Mutex
	var debounced = func(ctx context.Context) (string, error) {
		m.Lock()
		defer func() { // т.к. defer, то threshold будет инициализирован будущей датой только после первого выполнения circuit(ctx)
			threshold = time.Now().Add(d)
			fmt.Printf("threshold defined by %v + %v = %v (deferred)\n", time.Now().Format(TIME_FORMAT), d, threshold.Format(TIME_FORMAT))
			m.Unlock()
		}()

		if time.Now().Before(threshold) {
			fmt.Printf("now %v is before threshold %v, thus return cached result\n", time.Now().Format(TIME_FORMAT), threshold.Format(TIME_FORMAT))
			return "Cached - " + result, err
		} else {
			fmt.Printf("now %v is after threshold %v, thus call circuit()\n", time.Now().Format(TIME_FORMAT), threshold.Format(TIME_FORMAT))
		}

		result, err = circuit(ctx)
		
		return result, err
	}
	return debounced
}

func myFunction(ctx context.Context) (string, error) {
	fmt.Println("myFunction: Running...")
	time.Sleep(time.Second * 1) // Эмулируем сложную процедуру, напр. подключения к БД
	rand.Seed(time.Now().UnixNano())
	var r = rand.Intn(100)
	if r > 80 {
		return fmt.Sprintf("myFunction success (%d) :)", r), nil
	}
	return "", errors.New(fmt.Sprintf("myFunction failed (%d) :(", r))
}

func main() {
	fmt.Println("Start main")
	ctx := context.Background()
	var decoratedFunction Circuit = DebounceFirst(myFunction, time.Millisecond * 500)
	for {
		response, err := decoratedFunction(ctx)
		if err != nil {
			fmt.Printf("Result: %v\n", err)
		} else {
			fmt.Printf("Result: %v\n", response)
		}
		time.Sleep(time.Millisecond * time.Duration(100 * rand.Intn(10)))
	}
}