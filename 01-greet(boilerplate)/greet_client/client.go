package main

import (
	"fmt"
	"log"

	"github.com/disharjayanth/gRPC-golang/tree/main/01-greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("Could not dail server: %v", err)
	}

	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("Created client: %f", c)
}
