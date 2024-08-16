package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// c := make(chan os.Signal, 1)
	fmt.Println("Before")
	wait()
	fmt.Println("After")
}

// wait stops the main goroutine until an interrupt or kill signal is sent
func wait() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, os.Kill)
	fmt.Println(<-sig)
}
