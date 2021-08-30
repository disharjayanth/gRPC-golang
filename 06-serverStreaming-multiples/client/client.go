package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/disharjayanth/gRPC-golang/tree/main/06-serverStreaming-multiples/multiplespb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dailing %v", err)
	}

	client := multiplespb.NewMultiplesServiceClient(conn)

	req := multiplespb.MultiplesRequest{
		Num: 2,
	}

	resStream, err := client.Multiples(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error from server: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error receiving from server: %v", err)
		}

		fmt.Println(msg.GetNums())
	}
}
