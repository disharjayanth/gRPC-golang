package main

import (
	"context"
	"fmt"
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

	// req := &blogpb.CreateBlogRequest{
	// 	Blog: &blogpb.Blog{
	// 		AuthorId: "2",
	// 		Title:    "GoodBye",
	// 		Content:  "bye bye see you!!!",
	// 	},
	// }

	// createBlogRes, err := client.CreateBlog(context.Background(), req)
	// if err != nil {
	// 	log.Fatalf("Error creating blog: %v", err)
	// }

	// log.Printf("Response: %v", createBlogRes)

	// req := &blogpb.ReadBlogRequest{
	// 	BlogId: "613cefd6caac717672d7304a",
	// }

	// res, err := client.ReadBlog(context.Background(), req)
	// if err != nil {
	// 	log.Printf("Error reading blog from server: %v", err)
	// }

	// if res != nil {
	// 	log.Println("Read blog with id 613cefd6caac717672d7304b", res.Blog)
	// }

	// // update blog
	updatingBlog := &blogpb.UpdateBlogRequest{
		Blog: &blogpb.Blog{
			Id:       "613cf016caac717672d7304c",
			AuthorId: "2",
			Title:    "Ta Ta bye bye!!!!",
			Content:  "See you later",
		},
	}

	updatedBlog, err := client.UpdateBlog(context.Background(), updatingBlog)
	if err != nil {
		fmt.Printf("Error while updating: %v", err)
	}

	fmt.Printf("Updated Blog: %v", *updatedBlog.GetBlog())
}
