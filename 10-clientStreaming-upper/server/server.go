package main

import (
	"io"
	"log"
	"net"
	"strings"

	"github.com/disharjayanth/gRPC-golang/tree/main/10-clientStreaming-upper/upperpb"
	"google.golang.org/grpc"
)

type server struct {
	upperpb.UnimplementedUpperServiceServer
}

func (s *server) Upper(stream upperpb.UpperService_UpperServer) error {
	conUpperString := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// client stream ends
			res := &upperpb.UpperResponse{
				UpperStr: conUpperString,
			}
			return stream.SendAndClose(res)
		}
		if err != nil {
			log.Fatalf("Failed to receive client stream: %v", err)
		}
		conUpperString = conUpperString + " " + strings.ToUpper(req.GetStr())
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Failed to create tcp listener: %v", err)
	}

	grpcServer := grpc.NewServer()

	upperpb.RegisterUpperServiceServer(grpcServer, &server{})

	log.Println("Server serving @port:50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve @port:50051 %v", err)
	}
}
