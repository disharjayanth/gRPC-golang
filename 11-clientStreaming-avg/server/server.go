package main

import (
	"io"
	"log"
	"net"

	"github.com/disharjayanth/gRPC-golang/tree/main/11-clientStreaming-avg/avgpb"
	"google.golang.org/grpc"
)

type server struct {
	avgpb.UnimplementedAvgServiceServer
}

func (s *server) Avg(stream avgpb.AvgService_AvgServer) error {
	var avg float64
	var sum float64
	var i int64 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// client stream ends
			avg = sum / float64(i)
			res := &avgpb.AvgResponse{
				Result: avg,
			}
			return stream.SendAndClose(res)
		}
		if err != nil {
			log.Fatalf("Error receiving client stream: %v", err)
		}
		sum += req.GetNum()
		i++
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Error creating tcp listener: %v", err)
	}

	grpcServer := grpc.NewServer()

	avgpb.RegisterAvgServiceServer(grpcServer, &server{})

	log.Println("Server running @port:50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error serving server @port:50051 %v", err)
	}
}
