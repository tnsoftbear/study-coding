package main

import (
	"fmt"
	"log"
	"net/rpc"
	"rpc-doc-ex/api"
)

func main() {
	serverAddress := "localhost"
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// Synchronous call
	args := &api.Args{7, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	// Asynchronous call
	args = &api.Args{71, 35}
	quotient := new(api.Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	_ = <-divCall.Done // will be equal to divCall
	fmt.Printf("%d/%d = Quo: %d Rem: %d", args.A, args.B, quotient.Quo, quotient.Rem)
	// check errors, print, etc.
}
