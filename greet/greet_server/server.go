package main

import (
	"../greetpb"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	greetpb.UnimplementedGreetingServiceServer
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
        greetpb.RegisterGreetingServiceServer(srv, &server{})
	log.Println("Listen on :50051")
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (*server) SayHello(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
    name := req.GetName()
    log.Printf("Received :", name)
    return &greetpb.GreetResponse{Message: "Hello " + name}, nil
}

