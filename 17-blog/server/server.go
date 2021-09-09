package main

import (
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/disharjayanth/gRPC-golang/tree/main/17-blog/blogpb"
	"google.golang.org/grpc"
)

type server struct {
	blogpb.UnimplementedBlogServiceServer
}

func main() {
	// if go code crashes, it prints file name and also line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("Error in tcp listener: %v", err)
	}

	grpcServer := grpc.NewServer()

	blogpb.RegisterBlogServiceServer(grpcServer, &server{})

	go func() {
		log.Println("Server serving at localhost:50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Error servering @port:50051 --> %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	log.Println("Stopping server")
	grpcServer.Stop()
	log.Println("Stopping listener")
	lis.Close()
	log.Println("Server completly stopped.")
}
