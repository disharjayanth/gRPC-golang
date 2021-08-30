package main

import (
	"context"
	"fmt"
	"log"

	"github.com/disharjayanth/gRPC-golang/tree/main/04-multi-practise/multipb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to created connection at right target: %v", err)
	}

	client := multipb.NewMultiplyServiceClient(conn)

	req := multipb.MultiplyRequest{
		Num_1: 10,
		Num_2: 2,
	}
	res, err := client.Multiply(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error from server: %v", err)
	}

	log.Printf("Response: %v", res.GetResult())
}
