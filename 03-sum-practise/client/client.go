package main

import (
	"context"
	"log"

	"github.com/disharjayanth/gRPC-golang/tree/main/03-sum-practise/sumpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		log.Fatalf("Error dailing server: %v", err)
	}

	client := sumpb.NewCalculatorClient(conn)

	req := sumpb.SumRequest{
		Num_1: 12,
		Num_2: 3,
	}

	res, err := client.Sum(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error from server: %v", err)
	}

	log.Println("Result of sum:", res.GetResult())
}
