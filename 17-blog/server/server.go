package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"time"

	"github.com/disharjayanth/gRPC-golang/tree/main/17-blog/blogpb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection

type server struct {
	blogpb.UnimplementedBlogServiceServer
}

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id",omitempty`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func main() {
	// if go code crashes, it prints file name and also line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// mongo client to mongo server
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://user:user@cluster0.0yfhp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatalf("Cannot create new client connection to mongo server: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Println("Connecting to mong client")
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Cannot connect client to server: %v", err)
	}
	defer client.Disconnect(ctx)

	collection = client.Database("blogDB").Collection("blog")
	fmt.Println("Collection:", collection)

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Fatalf("Error fetching databases names: %v", err)
	}

	fmt.Println("Databases: ", databases)

	// tcp listener for grpc server
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
	log.Println("Stopping tcp listener")
	lis.Close()
	log.Println("Server completly stopped.")
}
