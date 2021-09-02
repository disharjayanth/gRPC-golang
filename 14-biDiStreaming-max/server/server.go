package main

import (
	"io"
	"log"
	"net"

	"github.com/disharjayanth/gRPC-golang/tree/main/14-biDiStreaming-max/maxpb"
	"google.golang.org/grpc"
)

type server struct {
	maxpb.UnimplementedMaxServiceServer
}

func (s *server) Max(stream maxpb.MaxService_MaxServer) error {
	var max int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// end of client stream
			return nil
		}
		if err != nil {
			log.Fatalf("Error while fetching client stream: %v", err)
			return err
		}
		num := req.GetNum()
		if num > max {
			max = num
			res := &maxpb.MaxResponse{
				Max: num,
			}
			if err := stream.Send(res); err != nil {
				log.Fatalf("Error while sending response to client: %v", err)
				return err
			}
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Error in tcp listener: %v", err)
	}

	grpcServer := grpc.NewServer()

	maxpb.RegisterMaxServiceServer(grpcServer, &server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error serving @port:50051 %v", err)
	}
}
