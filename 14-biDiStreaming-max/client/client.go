package main

import (
	"context"
	"io"
	"log"

	"github.com/disharjayanth/gRPC-golang/tree/main/14-biDiStreaming-max/maxpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while dailing client: %v", err)
	}

	clientStream := maxpb.NewMaxServiceClient(conn)

	client, err := clientStream.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while calling RPC max: %v", err)
	}

	waitC := make(chan struct{})
	// send bunch of request
	go func() {
		sliceOfInts := []int32{1, 5, 3, 6, 2, 20}
		for _, v := range sliceOfInts {
			req := &maxpb.MaxRequest{
				Num: v,
			}
			if err := client.Send(req); err != nil {
				log.Fatalf("Error while sending request to response: %v", err)
			}
		}

		if err := client.CloseSend(); err != nil {
			log.Fatalf("Error while closing client stream: %v", err)
		}
	}()

	// receive bunch of response
	go func() {
		for {
			res, err := client.Recv()
			if err == io.EOF {
				close(waitC)
				return
			}
			if err != nil {
				log.Fatalf("Error while receiving response from server: %v", err)
				close(waitC)
				return
			}
			log.Println("Response:", res.GetMax())
		}
	}()

	<-waitC
}
