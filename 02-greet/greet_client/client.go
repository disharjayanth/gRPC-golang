package main

import (
	"context"
	"fmt"
	"log"

	"github.com/disharjayanth/gRPC-golang/tree/main/02-greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer cc.Close()
	if err != nil {
		log.Fatalf("Couldn't dail: %v", err)
	}

	c := greetpb.NewGreetServiceClient(cc)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	req := greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Joe",
			LastName:  "Rogan",
		},
	}

	res, err := c.Greet(context.Background(), &req)
	if err != nil {
		log.Fatalln("Error from server RPC:", err)
	}

	fmt.Println("Response from server RPC:", res.GetResult())
}
