package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const TIME_FORMAT = "15:04:05.0000"
type Circuit func(context.Context) (string, error)

func DebounceLast(circuit Circuit, d time.Duration) Circuit {
	var threshold time.Time = time.Now()
	var ticker *time.Ticker
	var result string
	var err error
	var once sync.Once
	var m sync.Mutex
	var debounced = func(ctx context.Context) (string, error) {
		m.Lock()
		defer m.Unlock()
		threshold = time.Now().Add(d)
		fmt.Printf("Now: %v threshold calculated: %v\n", time.Now().Format(TIME_FORMAT), threshold.Format(TIME_FORMAT))
		var doOnce = func() {
			ticker = time.NewTicker(time.Millisecond * 100)
			go func() {
				defer func() {
					m.Lock()
					ticker.Stop()
					once = sync.Once{}
					m.Unlock()
				}()
				for {
					select {
					case <-ticker.C:
						m.Lock()
						fmt.Printf("<-Ticker.C Now: %v, threshold: %v, is after: %v\n", time.Now().Format(TIME_FORMAT), threshold.Format(TIME_FORMAT), time.Now().After(threshold))
						if time.Now().After(threshold) {
							result, err = circuit(ctx)
							m.Unlock()
							return
						}
						m.Unlock()
					case <-ctx.Done():
						fmt.Printf("<-ctx.Done\n")
						m.Lock()
						result, err = "", ctx.Err()
						m.Unlock()
						return
					}
				}
			}()
		}

		once.Do(doOnce)
		fmt.Printf("Return result]\n")
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
	var decoratedFunction Circuit = DebounceLast(myFunction, time.Millisecond*400)
	for {
		response, err := decoratedFunction(ctx)
		if err != nil {
			fmt.Printf("Result-: %v\n", err)
		} else {
			fmt.Printf("Result+: %v\n", response)
		}
		time.Sleep(time.Millisecond * 500)
	}
}
