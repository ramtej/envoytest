package main

import (
	"log"
	"net"

	pb "github.com/TabbDrinkLTD/envtest/service/models"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) Authenticate(ctx context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error) {
	log.Println(in.Email)
	return &pb.AuthResponse{Done: true}, nil
}

func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Println(in.Id)
	return &pb.GetUserResponse{Id: in.Id, Name: "Ewan Valentine"}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &server{})
	// Register reflection service on gRPC server.
	reflection.Register(s)

	log.Println("Listening on port:", port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
