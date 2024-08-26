package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg         sync.WaitGroup
	sharedLock sync.Mutex
)

const runtime = time.Second

func greedyWorker() {
	defer wg.Done()
	var count int
	for begin := time.Now(); time.Since(begin) <= runtime; {
		sharedLock.Lock()
		time.Sleep(3 * time.Nanosecond)
		sharedLock.Unlock()
		count++
	}
	fmt.Printf("Greeedy worker was able to execute %d work loops\n", count)
}

func politeWorker() {
	defer wg.Done()
	var count int
	for begin := time.Now(); time.Since(begin) <= runtime; {
		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		sharedLock.Lock()
		time.Sleep(1 * time.Nanosecond)
		sharedLock.Unlock()

		count++
	}
	fmt.Printf("Polite worker was able to execute %d work loops\n", count)
}

func main() {
	wg.Add(2)
	go greedyWorker()
	go politeWorker()
	wg.Wait()
}
