package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"

	"grpc-calc/internal/pb/calculator"
)

func main() {
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := calculator.NewCalculatorServiceClient(conn)

	rand.Seed(time.Now().UnixNano())
	var num1 int32 = int32(rand.Intn(101))
	var num2 int32 = int32(rand.Intn(101))

	req := &calculator.AddRequest{Num1: num1, Num2: num2}
	res, err := client.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling Add: %v", err)
	}

	fmt.Printf("%d + %d = %d\n", num1, num2, res.Result)
}
