package main

import (
	"log"
	"net"
	"time"

	"github.com/disharjayanth/gRPC-golang/tree/main/06-serverStreaming-multiples/multiplespb"
	"google.golang.org/grpc"
)

type server struct {
	multiplespb.UnimplementedMultiplesServiceServer
}

func (s *server) Multiples(req *multiplespb.MultiplesRequest, stream multiplespb.MultiplesService_MultiplesServer) error {
	log.Println("Multiples RPC called : ", req.GetNum())
	num := req.GetNum()
	var i int32 = 1
	for {
		multiNum := num * i
		if multiNum == 100 {
			break
		}

		res := &multiplespb.MultiplesResponse{
			Nums: multiNum,
		}

		stream.Send(res)

		i++

		time.Sleep(1 * time.Second)
	}

	return nil
}

func main() {
	conn, err := net.Listen("tcp", "0.0.0.0:50051")
	errorLog("tcp connection error", err)

	grpcServer := grpc.NewServer()

	multiplespb.RegisterMultiplesServiceServer(grpcServer, &server{})

	log.Println("Listening server @port localhost:50051")
	err = grpcServer.Serve(conn)
	errorLog("serving server error", err)
}

func errorLog(errormsg string, err error) {
	if err != nil {
		log.Printf("%s - %v\n", errormsg, err)
	}
}
