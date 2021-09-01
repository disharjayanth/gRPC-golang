package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/disharjayanth/gRPC-golang/tree/main/13-biDiStreaming-double/doublepb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dailing to server: %v", err)
	}

	client := doublepb.NewDoubleServiceClient(conn)

	nums := []int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	clientStream, err := client.Double(context.Background())
	if err != nil {
		log.Fatalf("Error in client stream: %v", err)
	}

	waitC := make(chan struct{})

	// send bunch of request
	go func() {
		for _, num := range nums {
			req := &doublepb.DoubleRequest{
				Num: num,
			}
			time.Sleep(1 * time.Second)
			if err := clientStream.Send(req); err != nil {
				log.Fatalf("Error while sending request to server: %v", err)
			}
		}
		if err := clientStream.CloseSend(); err != nil {
			log.Fatalf("Error while closing send stream: %v", err)
		}
	}()

	// receive bunch of response
	go func() {
		for {
			res, err := clientStream.Recv()
			if err == io.EOF {
				// end of server stream
				close(waitC)
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving server stream: %v", err)
				close(waitC)
				break
			}
			result := res.GetResult()
			log.Println("Response: ", result)
		}
	}()

	<-waitC
}
