package main

import (
	"../greetpb"
	"context"
	"google.golang.org/grpc"
	"log"
        "time"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer conn.Close()
        name := "Mahesh"
        clnt := greetpb.NewGreetingServiceClient(conn)
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
	r, err := clnt.SayHello(ctx, &greetpb.GreetRequest{Name: name})      
        if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())       
       callGreetingService(clnt)
}

func callGreetingService(clnt greetpb.GreetingServiceClient) {
      name := "Mahesh"
      log.Printf("Greet :", name)
      req := &greetpb.GreetRequest{Name: name}

     res, err := clnt.SayHello(context.Background(), req)
        if err != nil {
                log.Fatalf("Failed to call Greet Service: %v", err)
        }
     log.Println("Message :", res.GetMessage())
}
