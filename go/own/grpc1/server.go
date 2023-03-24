package main

import (
	"context"
	"fmt"
	"net"
	
	"google.golang.org/grpc"

	"grpc1/gen/calculator"
)

type calculatorServer struct{
	calculator.UnimplementedCalculatorServiceServer
}

func (s *calculatorServer) Add(ctx context.Context, req *calculator.AddRequest) (*calculator.AddResponse, error) {
	result := req.Num1 + req.Num2
	return &calculator.AddResponse{Result: result}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	calculator.RegisterCalculatorServiceServer(srv, &calculatorServer{})

	fmt.Println("Server started at :8080")
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
