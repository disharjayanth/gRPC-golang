package main

import (
	"context"
	"log"
	"time"

	"github.com/disharjayanth/gRPC-golang/tree/main/09-clientStreaming-sum/sumpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to create client connection: %v", err)
	}

	client := sumpb.NewSumServiceClient(conn)

	clientStream, err := client.Sum(context.Background())
	if err != nil {
		log.Fatalf("Failed to create clientStream: %v ", err)
	}

	nums := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, num := range nums {
		req := &sumpb.SumRequest{
			Num: num,
		}
		clientStream.Send(req)
		time.Sleep(1 * time.Second)
	}

	err = clientStream.CloseSend()
	if err != nil {
		log.Fatalf("Failed to close client stream: %v", err)
	}
}
