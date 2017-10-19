package main

import (
	"log"

	pb "github.com/ewanvalentine/envoytest/hello-service/models"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

const (
	port = ":50052"
)

// server is used to implement helloworld.GreeterServer.
type Server struct{}

// SayHello implements helloworld.GreeterServer
func (s *Server) Hello(ctx context.Context, in *pb.HelloRequest, resp *pb.HelloResponse) error {
	log.Println("test")
	greeting := "Hello, " + in.Name
	log.Println(greeting)
	resp.Greeting = greeting
	return nil
}

func (s *Server) GetGreeting(ctx context.Context, in *pb.GreetingRequest, resp *pb.GreetingResponse) error {
	greeting := "Hello, " + in.Name
	resp.Greeting = greeting
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.hello"),
	)

	service.Init()
	pb.RegisterHelloServiceHandler(service.Server(), &Server{})

	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
