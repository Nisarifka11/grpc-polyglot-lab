
package main

import (
	"context"
	"log"
	"net"

	pb "grpc-polyglot-lab/server"

	"google.golang.org/grpc"
)

type userServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *userServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	log.Printf("Menerima permintaan untuk user ID: %d", req.GetId())

	return &pb.UserResponse{
		Id:    req.GetId(),
		Name:  "Nisatur Rifka",
		Email: "nisaturrifka@gmail.com",
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Gagal membuka port: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &userServer{})

	log.Println("Server gRPC berjalan di port 50051")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Server gagal: %v", err)
	}
}
