package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/disharjayanth/gRPC-golang/tree/main/02-greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s *server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet function was invoked....:", req)
	first_name := req.GetGreeting().GetFirstName()
	result := "Hello " + first_name
	return &greetpb.GreetResponse{
		Result: result,
	}, nil
}

func main() {
	fmt.Println("Listening @localhost:50051")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen at tcp: %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
