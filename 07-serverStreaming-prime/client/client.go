package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/disharjayanth/gRPC-golang/tree/main/07-serverStreaming-prime/primepb"
	"google.golang.org/grpc"
)

func main() {
	var num int32
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Println("Failed to dail localhost:50051", err)
	}

	client := primepb.NewPrimeServiceClient(conn)

	fmt.Println("Enter number for prime decomposition:")
	fmt.Scanf("%d", &num)
	req := &primepb.PrimeRequest{
		Num: num,
	}

	resStream, err := client.Prime(context.Background(), req)
	if err != nil {
		log.Fatalf("Error from stream server: %v", err)
	}
	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Error from server: %v", err)
		}

		result := res.GetResult()
		fmt.Println(result)
	}
}
