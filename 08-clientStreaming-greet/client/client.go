package main

import (
	"context"
	"log"

	"github.com/disharjayanth/gRPC-golang/tree/main/08-clientStreaming-greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dail server at localhost:50051 %v", err)
	}

	client := greetpb.NewGreetServiceClient(conn)

	clientStream, err := client.Greet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling server greet RPC: %v", err)
	}

	someName := map[string]string{
		"John":   "Smith",
		"Joe":    "Rogan",
		"Eddie":  "Gueero",
		"Warren": "Buffet",
		"Elon":   "Musk",
	}

	for fs, ls := range someName {
		req := &greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: fs,
				LastName:  ls,
			},
		}
		err := clientStream.Send(req)
		if err != nil {
			log.Fatalf("Error while sending client stream to server: %v", err)
		}
	}

	res, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving server response: %v", err)
	}

	log.Printf("Response from server: %v", res)
}
