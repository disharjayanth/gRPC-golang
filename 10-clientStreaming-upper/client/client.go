package main

import (
	"context"
	"fmt"
	"log"

	"github.com/disharjayanth/gRPC-golang/tree/main/10-clientStreaming-upper/upperpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while dailing server: %v", err)
	}

	client := upperpb.NewUpperServiceClient(conn)

	sliceString := []string{"hi", "how", "are", "you", "?", "I", "am", "doing", "fine"}

	clientStream, err := client.Upper(context.Background())
	if err != nil {
		log.Fatalf("Error while calling Upper rpc service: %v", err)
	}

	for _, str := range sliceString {
		req := &upperpb.UpperRequest{
			Str: str,
		}
		clientStream.Send(req)
	}

	res, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from server: %v", err)
	}
	fmt.Println("Response from server:", res.GetUpperStr())
}
