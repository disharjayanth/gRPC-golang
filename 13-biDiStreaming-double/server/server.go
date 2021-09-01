package main

import (
	"io"
	"log"
	"net"

	"github.com/disharjayanth/gRPC-golang/tree/main/13-biDiStreaming-double/doublepb"
	"google.golang.org/grpc"
)

type server struct {
	doublepb.UnimplementedDoubleServiceServer
}

func (s *server) Double(stream doublepb.DoubleService_DoubleServer) error {
	log.Println("Double RPC called")
	for {
		req, err := stream.Recv()
		log.Println(req)
		if err == io.EOF {
			// end of client stream
			return nil
		}
		if err != nil {
			log.Fatalf("Error with client stream request: %v", err)
			return err
		}
		num := req.GetNum()
		doubleNum := int(num) * 2
		res := &doublepb.DoubleResponse{
			Result: int32(doubleNum),
		}
		if err := stream.Send(res); err != nil {
			log.Fatalf("Error sending response to client: %v", err)
			return err
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Error creating tcp listener: %v", err)
	}

	grpcServer := grpc.NewServer()

	doublepb.RegisterDoubleServiceServer(grpcServer, &server{})

	log.Println("Server running at port@50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error servering at port@50051: %v", err)
	}
}
