package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

const TIME_FORMAT = "15:04:05.0000"

type Circuit func(context.Context) (string, error)

func Breaker(circuit Circuit, failureThreshold uint) Circuit {
	var consecutiveFailures int = 0
	var lastAttempt = time.Now()
	fmt.Printf("Breaker() initial state is - consecutiveFailures: %v, failureThreshold: %v\n", consecutiveFailures, failureThreshold)
	var m sync.RWMutex

	var circuitBreaker = func(ctx context.Context) (string, error) {
		m.RLock()
		failureRate := consecutiveFailures - int(failureThreshold)
		fmt.Printf("A] Now: %v, lastAttempt: %v, consecutiveFailures: %v, Failure rate: %d\n",
			time.Now().Format(TIME_FORMAT), lastAttempt.Format(TIME_FORMAT), consecutiveFailures, failureRate)
		/**
		 * Когда допустимый рейт ошибок превышен, мы больше не хотим запускать целевую ф-цию до тех пор пока не пройдёт интервал времени shouldRetryAt.
		 * Этот интервал времени так же увеличивается пропорционально увеличению рейта ошибок.
		 */
		if failureRate >= 0 {
			var secondToAdd time.Duration = time.Second * 2 << failureRate
			shouldRetryAt := lastAttempt.Add(secondToAdd)
			var isAfter bool = time.Now().After(shouldRetryAt)
			fmt.Printf(
				"B] Now: %v, lastAttempt: %v, secondToAdd: %v, shouldRetryAt: %v, now is After shouldRetryAt: %v\n",
				time.Now().Format(TIME_FORMAT), lastAttempt.Format(TIME_FORMAT), secondToAdd, shouldRetryAt.Format(TIME_FORMAT), isAfter)
			if !isAfter {
				m.RUnlock()
				return "", errors.New("service unreachable")
			}
		}
		m.RUnlock()

		response, err := circuit(ctx) // Выполнить целевую ф-цию

		m.Lock()
		defer m.Unlock()
		lastAttempt = time.Now() // Зафиксировать время попытки
		if err != nil {          // Если ваша ф-ция (circuit) вернула ошибку
			consecutiveFailures++ // увеличить счетчик ошибок
			return response, err  // и вернуть ошибку
		}

		consecutiveFailures = 0 // Иначе сбросить счетчик ошибок
		return response, nil
	}

	return circuitBreaker
}

func myFunction(ctx context.Context) (string, error) {
	fmt.Println("myFunction: Running...")
	time.Sleep(time.Second * 1) // Эмулируем сложную процедуру, напр. подключения к БД
	var success bool = false
	if success {
		return "myFunction success :)", nil
	}
	return "", errors.New("myFunction failed :(")
}

func main() {
	fmt.Println("Start main")
	ctx := context.Background()
	var decoratedFunction Circuit = Breaker(myFunction, 2)
	for {
		response, err := decoratedFunction(ctx)
		if err != nil {
			fmt.Printf("Result: %v\n", err)
		} else {
			fmt.Printf("Result: %v\n", response)
		}
		if err == nil {
			break
		}
		time.Sleep(time.Second * 1)
	}
}
