package main

import (
	"context"
	"log"
	"net"

	"github.com/disharjayanth/gRPC-golang/tree/main/03-sum-practise/sumpb"
	"google.golang.org/grpc"
)

type server struct {
	sumpb.UnimplementedCalculatorServer
}

func (s *server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	log.Println("Sum rpc called")
	num1 := req.Num_1
	num2 := req.Num_2
	sum := num1 + num2
	return &sumpb.SumResponse{
		Result: sum,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Could'nt create a tcp connection: %v", lis)
	}

	s := grpc.NewServer()

	sumpb.RegisterCalculatorServer(s, &server{})

	log.Println("Listening @port:50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
