package main

import (
	"context"
	"fmt"
	"log"

	"github.com/disharjayanth/gRPC-golang/tree/main/11-clientStreaming-avg/avgpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while dailing server: %v", err)
	}

	client := avgpb.NewAvgServiceClient(conn)

	sliceOfFloats := []float64{1, 2, 3, 4}

	clientStream, err := client.Avg(context.Background())

	for _, v := range sliceOfFloats {
		if err != nil {
			log.Fatalf("Error while calling Avg rpc: %v", err)
		}

		req := &avgpb.AvgRequest{
			Num: v,
		}

		err = clientStream.Send(req)
		if err != nil {
			log.Fatalf("Error while sending data to server: %v", err)
		}
	}

	res, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receving response from server: %v", err)
	}

	fmt.Printf("Response: %v", res.GetResult())
}
