package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/disharjayanth/gRPC-golang/tree/main/16-deadlines/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	greetpb.UnimplementedGreetDeadlineServiceServer
}

func (s *server) GreetDeadline(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	log.Println("GreetDeadline rpc called: ", req)
	time.Sleep(3 * time.Second)

	if ctx.Err() == context.Canceled {
		log.Println("Client cancelled request")
		return nil, status.Error(codes.DeadlineExceeded, "Client cancelled request")
	}

	firstName := req.GetGreeting().GetFirstName()

	res := &greetpb.GreetResponse{
		Result: "Hello " + firstName,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create tcp listener: %v", err)
	}

	grpcServer := grpc.NewServer()

	greetpb.RegisterGreetDeadlineServiceServer(grpcServer, &server{})

	log.Println("Server listening @port:50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error listening at @port:50051 %v", err)
	}
}
