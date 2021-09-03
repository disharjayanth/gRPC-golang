package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"github.com/disharjayanth/gRPC-golang/tree/main/15-errors/01-unary-sqRoots/sqrootpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	sqrootpb.UnimplementedSquareRootServiceServer
}

func (s *server) SquareRoot(ctx context.Context, req *sqrootpb.SquareRootRequest) (*sqrootpb.SquareRootResponse, error) {
	log.Println("SquareRoot rpc called with req:", req)
	num := req.GetNum()
	if num < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Received a negative number: %v", num))
	}
	sqRoot := math.Sqrt(float64(num))
	res := &sqrootpb.SquareRootResponse{
		Result: sqRoot,
	}
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Error in tcp listener: %v", err)
	}

	grpcServer := grpc.NewServer()

	sqrootpb.RegisterSquareRootServiceServer(grpcServer, &server{})

	log.Println("Server running @port:50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error listening at port:50051 %v", err)
	}
}
