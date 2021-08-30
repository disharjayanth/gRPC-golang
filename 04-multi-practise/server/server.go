package main

import (
	"context"
	"log"
	"net"

	"github.com/disharjayanth/gRPC-golang/tree/main/04-multi-practise/multipb"
	"google.golang.org/grpc"
)

type server struct {
	multipb.UnimplementedMultiplyServiceServer
}

func (s *server) Multiply(ctx context.Context, req *multipb.MultiplyRequest) (*multipb.MultiplyResponse, error) {
	num1 := req.GetNum_1()
	num2 := req.GetNum_2()
	log.Println("Called Multiply RPC with", num1, "&", num2)
	return &multipb.MultiplyResponse{
		Result: num1 * num2,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Cannot listen to tcp: %v", err)
	}

	grpcServer := grpc.NewServer()

	multipb.RegisterMultiplyServiceServer(grpcServer, &server{})

	log.Println("Listening @PORT 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to server @port:50051 %v", err)
	}
}
