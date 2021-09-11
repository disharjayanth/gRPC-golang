package main

import (
	"context"
	"log"

	"github.com/disharjayanth/gRPC-golang/tree/main/17-blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dail server: %v", err)
	}

	client := blogpb.NewBlogServiceClient(conn)

	req := &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			AuthorId: "2",
			Title:    "GoodBye",
			Content:  "bye bye see you!!!",
		},
	}

	createBlogRes, err := client.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Error creating blog: %v", err)
	}

	log.Printf("Response: %v", createBlogRes)
}
