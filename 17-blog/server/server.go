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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var collection *mongo.Collection

type blogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	AuthorID string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

type server struct {
	blogpb.UnimplementedBlogServiceServer
}

func (s *server) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	blog := req.GetBlog()

	blogData := blogItem{
		AuthorID: blog.GetAuthorId(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	res, err := collection.InsertOne(context.Background(), blogData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error while inserting blog: %v", err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprint("Cannot convert to OID"))
	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       oid.Hex(),
			AuthorId: blog.GetAuthorId(),
			Title:    blog.GetTitle(),
			Content:  blog.GetContent(),
		},
	}, nil
}

func (s *server) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	blogId := req.GetBlogId()
	oid, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse given blog id: %v", err))
	}

	data := &blogItem{}
	filter := bson.M{"_id": oid}

	err = collection.FindOne(context.Background(), filter).Decode(data)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cannot find blog with given id: %v", err))
	}

	return &blogpb.ReadBlogResponse{
		Blog: &blogpb.Blog{
			Id:       data.ID.String(),
			AuthorId: data.AuthorID,
			Title:    data.Title,
			Content:  data.Content,
		},
	}, nil
}

func (s *server) UpdateBlog(ctx context.Context, req *blogpb.UpdateBlogRequest) (*blogpb.UpdateBlogResponse, error) {
	blogId := req.GetBlog().GetId()
	blog := req.GetBlog()
	oid, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse given blog id: %v", err))
	}

	data := &blogItem{}
	filter := bson.D{{Key: "_id", Value: oid}}

	update := bson.D{
		{Key: "$set", Value: bson.D{{Key: "title", Value: blog.GetTitle()}}},
		{Key: "$set", Value: bson.D{{Key: "content", Value: blog.GetContent()}}},
		{Key: "$set", Value: bson.D{{Key: "author_id", Value: blog.GetAuthorId()}}},
	}

	if err := collection.FindOneAndUpdate(context.Background(), filter, update).Decode(&data); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error while updating: %v", err))
	}

	return &blogpb.UpdateBlogResponse{
		Blog: &blogpb.Blog{
			Id:       data.ID.Hex(),
			AuthorId: data.AuthorID,
			Title:    data.Title,
			Content:  data.Content,
		},
	}, nil
}

func main() {
	// if go code crashes, it prints file name and also line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// mongo client connecting to mongo server
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://user:user@cluster0.0yfhp.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatalf("Cannot create new client connection to mongo server: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("Connecting to mongo client")
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Cannot connect client to server: %v", err)
	}
	defer client.Disconnect(ctx)

	collection = client.Database("blogDB").Collection("blog")
	fmt.Println("Collection:", collection)

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
	log.Println("Server completely stopped.")
}
