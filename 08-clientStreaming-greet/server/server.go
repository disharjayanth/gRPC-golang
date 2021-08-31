package main

import (
	"io"
	"log"
	"net"

	"github.com/disharjayanth/gRPC-golang/tree/main/08-clientStreaming-greet/greetpb"
	"google.golang.org/grpc"
)

type server struct {
	greetpb.UnimplementedGreetServiceServer
}

func (s *server) Greet(stream greetpb.GreetService_GreetServer) error {
	log.Println("Greet RPC was called ")
	result := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// end of client request
			res := &greetpb.LongGreetResponse{
				Result: result,
			}
			return stream.SendAndClose(res)
		}
		if err != nil {
			log.Fatalf("Error while receving client stream: %v", err)
		}
		result += "Hello " + req.GetGreeting().GetFirstName() + "! "
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Cannot create tcp listener: %v", err)
	}

	grpcServer := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(grpcServer, &server{})

	log.Println("Servering server @PORT:50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Cannot serve server @port 50051: %v", err)
	}
}
