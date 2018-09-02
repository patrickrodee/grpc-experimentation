package main

import (
  "log"
  "net"
  
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
  pb "github.com/patrickrodee/grpc-experimentation/protos/helloworld"
)

const (
  port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
  log.Printf("Message received: %s\n", in.Name)
  return &pb.HelloResponse{Reply: "Hello " + in.Name}, nil
}

func main() {
  lis, err := net.Listen("tcp", port)
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  log.Print("Server started")
  s := grpc.NewServer()
  pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on gRPC server.
  reflection.Register(s)
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}