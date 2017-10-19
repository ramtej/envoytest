package main

import (
	"log"

	pb "github.com/ewanvalentine/envoytest/service/models"
	micro "github.com/micro/go-micro"
	"golang.org/x/net/context"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Authenticate(ctx context.Context, in *pb.AuthRequest, resp *pb.AuthResponse) error {
	log.Println(in.Email)
	resp.Done = true
	return nil
}

func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest, resp *pb.GetUserResponse) error {
	log.Println(in.Id)
	resp.Id = in.Id
	resp.Name = "Ewan Valentine"
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.service"),
	)

	service.Init()
	pb.RegisterAuthServiceHandler(service.Server(), &server{})

	// Run the server
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
