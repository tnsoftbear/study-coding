package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"rpc-doc-ex/api"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	arith := new(api.Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	go http.Serve(l, nil)

	wait()
}

// wait stops the main goroutine until an interrupt or kill signal is sent
func wait() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, os.Kill)
	log.Println(<-sig)
}
