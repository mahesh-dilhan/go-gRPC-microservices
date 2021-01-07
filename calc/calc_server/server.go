package main

import (
	"../calcpb"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	calcpb.RegisterCalculatorServiceServer(srv, &server{})
	log.Println("Listen on :50051")
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func (*server) CalcSum(ctx context.Context, req *calcpb.CalcSumRequest) (*calcpb.CalcSumResponse, error) {
	args := req.GetArgs()
	var sum, a int32
	for _, a = range args.GetArg() {
		sum += a
	}
	res := &calcpb.CalcSumResponse{
		Sum: sum,
	}
	return res, nil
}
