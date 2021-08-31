package main

import (
	"log"
	"net"
	"time"

	"github.com/disharjayanth/gRPC-golang/tree/main/07-serverStreaming-prime/primepb"
	"google.golang.org/grpc"
)

type server struct {
	primepb.UnimplementedPrimeServiceServer
}

func (s *server) Prime(req *primepb.PrimeRequest, stream primepb.PrimeService_PrimeServer) error {
	log.Println("Prime RPC called: ", req)
	num := req.GetNum()
	k := 2
	for num > 1 {
		if num%int32(k) == 0 {
			res := &primepb.PrimeResponse{
				Result: int32(k),
			}
			stream.Send(res)
			num = num / int32(k)
			time.Sleep(1 * time.Second)
		} else {
			k = k + 1
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Cannot listen to tcp connection at port 50051: %v", err)
	}

	grpcSever := grpc.NewServer()

	primepb.RegisterPrimeServiceServer(grpcSever, &server{})

	log.Println("Servering server at port@50051")
	if err := grpcSever.Serve(lis); err != nil {
		log.Fatalf("Cannot serve server %v", err)
	}
}
