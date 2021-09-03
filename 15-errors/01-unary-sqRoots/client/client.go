package main

import (
	"context"
	"fmt"
	"log"

	"github.com/disharjayanth/gRPC-golang/tree/main/15-errors/01-unary-sqRoots/sqrootpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error while dailing to client: %v", err)
	}

	client := sqrootpb.NewSquareRootServiceClient(conn)

	var num int32
	fmt.Println("Enter the num for squareroot:")
	fmt.Scanf("%d", &num)
	req := &sqrootpb.SquareRootRequest{
		Num: num,
	}

	res, err := client.SquareRoot(context.Background(), req)
	if err != nil {
		// status.FromError is grpc friendly error
		respErr, ok := status.FromError(err)
		if ok {
			// error from user input
			log.Printf("Error with server response: %v", respErr.Message())
			log.Fatalf(respErr.Code().String())
		} else {
			// big error from SquareRoot function
			log.Fatalf("Error %v", err)
		}
	}

	log.Println("Response:", res.GetResult())
}
