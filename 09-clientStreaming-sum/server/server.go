package main

import (
	"io"
	"log"
	"net"

	"github.com/disharjayanth/gRPC-golang/tree/main/09-clientStreaming-sum/sumpb"
	"google.golang.org/grpc"
)

type server struct {
	sumpb.UnimplementedSumServiceServer
}

func (s *server) Sum(stream sumpb.SumService_SumServer) error {
	log.Println("Sum rpc called with req:", stream)
	sum := 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// end of client stream
			res := &sumpb.SumResponse{
				Result: int32(sum),
			}
			return stream.SendAndClose(res)
		}
		if err != nil {
			log.Fatalf("Failed to recieve client stream: %v", err)
		}

		num := req.GetNum()
		sum += int(num)
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to created tcp listener: %v", err)
	}

	grpcServer := grpc.NewServer()

	sumpb.RegisterSumServiceServer(grpcServer, &server{})

	log.Println("Listening @port:50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve at port@50051: %v", err)
	}
}
