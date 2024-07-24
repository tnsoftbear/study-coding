package main

import (
	"os"
	"time"
)

func main() {
	// Exit after 200 milliseconds - simply to not block forever in example.
	go func() {
		time.Sleep(200 * time.Millisecond)
		os.Exit(0)
	}()

	subCh := make(chan string, 128)
	for i := 0; i < 128; i++ {
		subCh <- "channel"
	}

	maxBatchSize := 50

	for {
		select {
		case channel := <-subCh:
			batch := []string{channel}
		loop:
			for len(batch) < maxBatchSize {
				select {
				case channel := <-subCh:
					batch = append(batch, channel)
					//time.Sleep(2 * time.Millisecond)
				default:
					break loop
				}
			}
			// Do sth with collected batch.
			println(len(batch))
		}
	}
}
