package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/disharjayanth/gRPC-golang/tree/main/16-deadlines/greetpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Error dailing client @port:50051", err)
	}

	client := greetpb.NewGreetDeadlineServiceClient(conn)

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Joe",
			LastName:  "Rogan",
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res, err := client.GreetDeadline(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			// grpc Server error
			if statusErr.Code() == codes.DeadlineExceeded {
				log.Fatalln("Timeout was hit! Deadline was exceeded")
			} else {
				log.Fatalln("Unexpected error:", err)
			}
		} else {
			log.Fatalf("Error from server with GreetDeadline rpc call: %v", err)
		}
	}

	fmt.Println("Response: ", res.GetResult())
}
