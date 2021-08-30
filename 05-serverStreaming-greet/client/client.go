package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/disharjayanth/gRPC-golang/tree/main/05-serverStreaming-greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatalf("cannot close client connection: %v", err)
		}
	}()
	if err != nil {
		log.Fatalf("Cannot create connection: %v", err)
	}

	client := greetpb.NewGreetServiceClient(conn)

	doUnary(client)

	doServerStream(client)
}

func doUnary(client greetpb.GreetServiceClient) {
	fmt.Println("Unary RPC call....")
	req := greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Joe",
			LastName:  "Rogan",
		},
	}

	res, err := client.Greet(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error response from server: %v", err)
	}

	fmt.Printf("Greet Response from server: %v\n", res.GetResult())
}

func doServerStream(client greetpb.GreetServiceClient) {
	fmt.Println("Server stream RPC call....")

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Joe",
			LastName:  "Alex",
		},
	}

	resStream, err := client.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error from server: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// end of stream
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		fmt.Println(msg.GetResult())
	}
}
