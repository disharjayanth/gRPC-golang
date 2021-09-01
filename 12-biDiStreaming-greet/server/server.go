package main

import (
	"io"
	"log"
	"net"

	"github.com/disharjayanth/gRPC-golang/tree/main/12-biDiStreaming-greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetEveryoneServiceServer
}

func (s *server) GreetEveryone(stream greetpb.GreetEveryoneService_GreetEveryoneServer) error {
	log.Println("GreetEveryone RPC was called")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// client is done sending requests
			return nil
		}
		if err != nil {
			log.Fatalf("Error while receiving client stream: %v", err)
			return err
		}
		firstName := req.GetGreeting().GetFirstName()
		result := "Hello " + firstName + "!"
		res := &greetpb.GreetEveryoneResponse{
			Result: result,
		}

		err = stream.Send(res)
		if err != nil {
			log.Fatalf("Error while sending stream response to client: %v", err)
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Error creating tcp listener: %v", err)
	}

	grpcServer := grpc.NewServer()

	greetpb.RegisterGreetEveryoneServiceServer(grpcServer, &server{})

	log.Println("Server running at @port:50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error servering server @port:50051 %v", err)
	}
}
