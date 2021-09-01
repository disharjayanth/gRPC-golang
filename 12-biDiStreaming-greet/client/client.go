package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/disharjayanth/gRPC-golang/tree/main/12-biDiStreaming-greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error dailing client @port:50051 %v", err)
	}

	client := greetpb.NewGreetEveryoneServiceClient(conn)

	clientStream, err := client.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error receiving client stream: %v", err)
	}

	names := map[string]string{
		"John":   "Smith",
		"Joe":    "Rogan",
		"Martin": "Lurther king",
		"Joes":   "Alex",
	}

	waitC := make(chan struct{})

	// send bunch of request to server
	go func() {
		for fs, ls := range names {
			req := &greetpb.GreetEveryoneRequest{
				Greeting: &greetpb.Greeting{
					FirstName: fs,
					LastName:  ls,
				},
			}

			time.Sleep(1 * time.Second)
			log.Println("Sending request:", req)
			if err := clientStream.Send(req); err != nil {
				log.Fatalf("Errore sending request to server: %v", err)
			}
		}

		if err := clientStream.CloseSend(); err != nil {
			log.Fatalf("Error closing send stream: %v", err)
		}
	}()

	// receive bunch of response from server
	go func() {
		for {
			res, err := clientStream.Recv()
			if err == io.EOF {
				// server is done sending all response
				// close channel after server sends all response
				close(waitC)
				break
			}
			if err != nil {
				log.Fatalf("Error receving response from server: %v", err)
				close(waitC)
				break
			}
			fmt.Println("Response: ", res.GetResult())
		}
	}()

	// block until every response is received
	<-waitC
	fmt.Println("Exited")
}
