package main

import (
	"github.com/tnsoftbear/study-coding/tree/master/go/own/distributed-go-book/internal/server"
	"log"
)

func main() {
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}
