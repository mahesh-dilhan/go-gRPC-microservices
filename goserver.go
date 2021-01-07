package main

import ("net"
        "log"
        "google.golang.org/grpc"
)

func main() {

  lis, err := net.Listen("tcp", ":9000")
  if err != nil {
      panic(err)
  }

  grpcserver := grpc.NewServer()
  
 if err := grpcserver.Serve(lis); err !=nil {
    log.Fatalf("server caused an issue , %s", err)

 }


}
