package main

import (
	"log"
	"net"

	pb "github.com/ewanvalentine/envoytest/hello-service/models"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50052"
)

// server is used to implement helloworld.GreeterServer.
type Server struct{}

// SayHello implements helloworld.GreeterServer
func (s *Server) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Println("test")
	greeting := "Hello, " + in.Name
	log.Println(greeting)
	return &pb.HelloResponse{Greeting: greeting}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println(&Server{})

	s := grpc.NewServer()
	pb.RegisterHelloServiceServer(s, &Server{})
	reflection.Register(s)

	log.Println("running on port:", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
